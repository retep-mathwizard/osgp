package sec

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/howeyc/gopass"
	c "github.com/skilstak/go/colors"
	s "github.com/skilstak/go/settings"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

const (
	AuthSeconds  = 300
	AuthEID      = 1000    // standard `sks` account ID
	AuthEName    = "admin" // standard `sks` account ID
	AuthPassReq  = false   // require passphrase for all Set
	privFileName = "/home/admin/.key.priv"
	pubFileName  = "/home/admin/.key.pub"
)

var origUID int

// cache
var cached_passphrase []byte
var cached_priv string
var cached_pub string

// Change effective run permissions to settings.AdminUID
func AdminOn() error {
	origUID = syscall.Geteuid()
	// note that Setuid equiv seems buggy, Setreuid is solid currently
	if err := syscall.Setreuid(s.AdminUID, s.AdminUID); err != nil {
		return err
	}
	return nil
}

// Change effective run permissions back to read uid
func AdminOff() error {
	if err := syscall.Setreuid(origUID, origUID); err != nil {
		return err
	}
	return nil
}

// Execute function as admin (needs `chown admin foo; chmod u+s foo`)
func AsAdmin(do func() error) error {
	if err := AdminOn(); err != nil {
		return err
	}
	if err := do(); err != nil {
		return err
	}
	if err := AdminOff(); err != nil {
		return err
	}
	return nil
}

func IsAuthEID() bool {
	return os.Geteuid() == AuthEID
}

func PrivatePrompt(s string) []byte {
	fmt.Printf(s)
	p := gopass.GetPasswd()
	return p
}

// AuthenticateAdmin first checks a previous authentication has succeeded.
// If the passphrase typed in fails the program exits. This encourages
// early authentication rather than waiting later if authentication is
// required at all.
func AuthenticateAdmin(p []byte) bool {
	if len(p) <= 0 {
		p = cached_passphrase
		if len(p) >= 0 {
			return true
		}
	}
	etoken, _ := Encrypt("blah")
	dtoken, _ := Decrypt(etoken, p)
	if dtoken == "blah" {
		return true
	}
	return false
}

func Encrypt(s string) (string, error) {
	bstream := bytes.NewBufferString(PubKey())
	ring, err := openpgp.ReadArmoredKeyRing(bstream)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, ring, nil, nil, nil)
	if err != nil {
		return "", err
	}
	w.Write([]byte(s))
	w.Close()
	bytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// Make sure to pass in a passphrase to prevent interactive prompt.
func Decrypt(e string, p []byte) (string, error) {
	if len(p) <= 0 {
		p = PrivatePrompt(c.R + "Oh really? " + c.X)
		cached_passphrase = p
	}
	b := bytes.NewBufferString(PrivKey())
	ring, err := openpgp.ReadArmoredKeyRing(b)
	if err != nil {
		return "", err
	}
	first := ring[0]
	if first.PrivateKey != nil && first.PrivateKey.Encrypted {
		err := first.PrivateKey.Decrypt(p)
		if err != nil {
			return "", err
		}
	}
	for _, subkey := range first.Subkeys {
		if subkey.PrivateKey != nil && subkey.PrivateKey.Encrypted {
			subkey.PrivateKey.Decrypt(p)
		}
	}
	dec, err := base64.StdEncoding.DecodeString(e)
	if err != nil {
		return "", err
	}
	md, err := openpgp.ReadMessage(bytes.NewBuffer(dec), ring, nil, nil)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PrivKey() (k string) {
	if len(cached_priv) > 0 {
		k = cached_priv
		return
	}
	return ReadFile(privFileName, true)
}

func PubKey() (k string) {
	if len(cached_pub) > 0 {
		k = cached_pub
		return
	}
	return ReadFile(pubFileName, false)
}

func ReadFile(fname string, check bool) string {
	path := PathInHome(fname)
	if check {
		PanicIfUnsafe(path)
	}
	b, _ := ioutil.ReadFile(path)
	return string(b)
}

func PathInHome(fname string) string {
	u, err := user.Lookup(AuthEName)
	if err != nil {
		panic(err)
	}
	dir := u.HomeDir
	return filepath.Join(dir, fname)
}

// Write a file securly (0600) with optional encryption
func WriteFile(f string, s string, e bool) error {
	path := PathInHome(f)
	var stuff = s
	var err error
	if e {
		stuff, err = Encrypt(s)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(path, []byte(stuff), 0600)
}

func PanicIfUnsafe(path string) {
	finfo, err := os.Stat(path)
	if err != nil {
		return
	}
	fmode := finfo.Mode().Perm()
	if fmode&0004 != 0 {
		panic("Private key has read permission for world (o).")
	}
	if fmode&0002 != 0 {
		panic("Private key has write permission for world (o).")
	}
	if fmode&0001 != 0 {
		panic("Private key has execute permission for world (o).")
	}
	dir := filepath.Dir(path)
	dinfo, _ := os.Stat(dir)
	dmode := dinfo.Mode().Perm()
	if dmode&0002 != 0 {
		panic("Private key directory is writable by world (o).")
	}
}

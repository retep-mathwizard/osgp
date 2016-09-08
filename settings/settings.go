/*
Package settings contains the default settings for any go apps at
SkilStak.  By putting them in one place we reduce the potential
work needed later if we decide to change our mind about where
dependent elements of SkilStak infrastructure should live and play.
This eventually will include hard-coded paths to data file and other
dependent commands as well as URLs and univeral API settings.

Constants

Use of constants for expected settings ensures consistency across
systems to catch those that may have been setup in ways that are
not compatible with general SkilStak infrastructure requirements,
for example, the Admin*, You*, and Stud* constants.

Functions

A few universally applicable functions have been added to the
settings package as well as a convenience because they deal with
settings. Please refrain from adding any function unrelated to
settings to this package.

Architecture

The settings represented are further explained in the
http://github.com/skilstak/share document.

*/
package settings

import (
	"io/ioutil"
	"os"
)

// Expected name of the account to use for admin privledges.
const AdminName string = "admin"

// Expected UID number for the admin account.
const AdminUID int = 1000

func IsAdmin() bool {
	return os.Geteuid() == AdminUID
}

// Expected GID number for the admin account.
const AdminGID int = 1000

// Expected name of the special "student" semi-public account.
const StudName string = "student"

// Expected UID of the special "student" semi-public account.
const StudUID int = 1001

// Expected GID of the special "student" semi-public account.
const StudGID int = 1001

// Expected name of the special "you" public account.
const YouName string = "you"

// Expected UID of the special "you" public account.
const YouUID int = 1002

// Expected GID of the special "you" public account.
const YouGID int = 1002

/*
Share is where all things SkilStak are installed on standard images.
Sym links are created in the system tying to this location. This
allows everything needed for a multi-user system to be located in
one place and saved as changes are made to a single GitHub repo.
(See http://github.com/skilstak/share for more.)
*/
const ShareDir string = "/usr/share/skilstak"

/*
DataDir contains all data used by all SkilStak apps on the system.
This usually includes a subdirectory for every skilstak user where
data is stored. Depending on the system and application this data
may need different security treatment. Generally, however, this
location should never contain private information that anyone on
the system should not be able to read. Usually this directory and
everything under it is owned by 'admin' (1001). Tools that modify
data here should be run as admin (or Setreuid/Setgeuid admin).
*/
const DataDir string = "/var/lib/skilstak"

func GetUserNames() []string {
	names, _ := ioutil.ReadDir(DataDir)
	var n []string
	for _, f := range names {
		n = append(n, f.Name())
	}
	return n
}

/*
UserDir returns a full path string pointing to the DataDir containing
the user's directory. UserDir should be used over constructing paths
with DataDir when users are involved. UserDir should not be confused
with a given user's home directory.
*/
func UserDir(user string) string {
	return DataDir + "/" + user
}

/*
InitUser is intended to setup a user in the DataDir as all apps will expect.
This will not initialize an existing user if detected and can
therefore be called safely by functions that expect a user to exist. To
reset a user to initial state call ResetUser instead, which calls DelUser
followed by InitUser.
*/
func InitUser(user string) error {
	// TODO check the effective run ownership and group to match DataDir
	// TODO create the home directory
	err := os.Mkdir(UserDir(user), 0755)
	if err != nil {
		return err
	}
	return nil
	//return fmt.Errorf("settings: unable to init user %s: %s", user, err)
}

// Calls DelUser then InitUser.
func ResetUser(user string) (err error) {
	err = DelUser(user)
	if err != nil {
		return err
	}
	err = InitUser(user)
	if err != nil {
		return err
	}
	return nil
}

// Deletes a user from the DataDir and any cache (if it exists)
func DelUser(user string) error {
	return os.RemoveAll(UserDir(user))
}

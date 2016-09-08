/*
The bux package provides the API for handling SkilStak SkilBux.
Contributors to this library should not add anything that would
create a dependency on any specific view of the any of the
data going in or coming out (no fmt.Println's for example). This
allows localized commands as well as microservice HTTP APIs to
use this same library.

Public functions should include any data validation required,
private functions can assume this has already been done and trust
the callers to provide correct parameters for their arguments.

Note: these functions are not transactional themselves so callers
of this library will need to work that out if and when high volume
concurrency is an issue (with semaphores, alternative transactional
database implementation, etc.).
*/
package bux

import (
	"fmt"
	c "github.com/skilstak/go/counter"
	h "github.com/skilstak/go/human"
	"github.com/skilstak/go/sec"
	s "github.com/skilstak/go/settings"
	"os"
	u "os/user"
	"strconv"
	"time"
)

const onlyAdmin bool = true

const NOTUSER string = "Doesn't look like a user."
const NOTNUM string = "Doesn't look like a number."
const NEEDSCOMMENT string = "Comment required."
const NEEDSPOS string = "Must be a positive number greater than zero."
const NOTENOUGH string = "Sorry you don't have enough."
const NOTHUMAN string = "Sorry you don't look human to me."
const SAME string = "Lazily refusing to transfer bux from you to you."

const FREESECONDS float64 = 86400

func buxFile(user string) string {
	return s.UserDir(user) + "/bux"
}

func buxLog(user string) string {
	return s.UserDir(user) + "/buxlog"
}

/*
Write a time-stamped entry to user's bux log.  This function is
public because sometimes a log entry unrelated to adding or subtracting
SkilBux is needed.
*/
func Log(user string, amount interface{}, comment string) error {
	switch v := amount.(type) {
	case int:
		amount = strconv.Itoa(v)
	}
	if !h.IsUser(user) {
		return fmt.Errorf(NOTUSER)
	}
	return log(user, amount.(string), comment)
}

func log(user string, amount string, comment string) error {
	if onlyAdmin {
		if err := sec.AdminOn(); err != nil {
			return err
		}
	}
	tstamp := time.Now().Format(time.RFC1123)
	f, err := os.OpenFile(buxLog(user), os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	fmt.Fprintf(f, "%s %s %s %s\n", tstamp, user, amount, comment)
	defer f.Close()
	if onlyAdmin {
		if err := sec.AdminOff(); err != nil {
			return err
		}
	}
	return nil
}

// Add or remove bux for the given user with required comment.
func Adjust(amount interface{}, user string, comment string) error {
	if onlyAdmin {
		if err := sec.AdminOn(); err != nil {
			return err
		}
	}
	switch v := amount.(type) {
	case int:
		amount = strconv.Itoa(v)
	}
	if err := c.AddInt(buxFile(user), amount.(string)); err != nil {
		return err
	}
	if err := log(user, amount.(string), comment); err != nil {
		return err
	}
	if onlyAdmin {
		if err := sec.AdminOff(); err != nil {
			return err
		}
	}
	return nil
}

// Returns Bux for given user or 0 if error. Always check for error.
func Get(u string) (int, error) {
	if !h.IsUser(u) {
		return 0, fmt.Errorf(NOTUSER)
	}
	return read(u)
}

// Return the total of all user SkilBux in the system (minus admin account)
func GetAll() (int, error) {
	students := s.GetUserNames()
	total := 0
	for _, student := range students {
		if student != "admin" {
			n, err := read(student)
			if err != nil {
				return 0, err
			}
			total += n
		}
	}

	return total, nil
}

func read(u string) (int, error) {
	path := buxFile(u)
	return c.ReadInt(path)
}

/*
Debit the current user and credit to target user.  Only positive
numbers allowed.  Not designed for high volume usage (no transactional
locking). Transfers to the same person from blocked.
*/
func Transfer(to string, amount interface{}, comment string) error {
	var err error
	if !h.IsUser(to) {
		return fmt.Errorf(NOTUSER)
	}
	switch v := amount.(type) {
	case string:
		if amount, err = strconv.Atoi(v); err != nil {
			return err
		}
	}
	current, err := u.Current()
	if err != nil {
		return err
	}
	from := current.Username
	if to == from {
		return fmt.Errorf(SAME)
	}
	has, err := read(from)
	if err != nil {
		return err
	}
	if amount.(int) <= 0 {
		return fmt.Errorf(NEEDSPOS)
	}
	if amount.(int) > has {
		return fmt.Errorf(NOTENOUGH)
	}
	// TODO make Add + Sub atomic transaction (but will need db)
	if err = Adjust(amount.(int), to, "from "+from+": "+comment); err != nil {
		return err
	}
	if err = Adjust(-amount.(int), from, "to "+to+": "+comment); err != nil {
		return err
	}
	return nil
}

// Return the number of seconds since the last Bux log entry
func Last(u string) float64 {
	stats, err := os.Stat(buxLog(u))
	if err != nil {
		return -1
	}
	return time.Since(stats.ModTime().Round(time.Second)).Seconds()
}

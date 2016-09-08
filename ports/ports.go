package ports

import (
	"os/user"
)

// Www returns an integer to be used as the www port for a user (8+uid)
func Www() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return "8" + u.Uid, nil
}

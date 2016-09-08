/*
The bux command provides a CLI to the SkilStak SkilBux system
using the same library called by the HTTP API. Contributors
to this command should not add 'business logic' to this
command-line interface adding it instead to the bux library
itself.
*/
package main

import (
	"fmt"
	"github.com/skilstak/go/bux"
	c "github.com/skilstak/go/colors"
	"github.com/skilstak/go/human"
	s "github.com/skilstak/go/settings"
	"os"
	"os/user"
	"strings"
)

const usage string = `usage: bux 
       bux USER
       bux USER AMOUNT WHY
`

var current string

func fail(m string) { fmt.Fprintln(os.Stderr, c.R+m+c.X+"\n"+usage); os.Exit(1) }

func transfer() {
	to := os.Args[1]
	astring := os.Args[2]
	comment := strings.Join(os.Args[3:], " ")
	err := bux.Transfer(to, astring, comment)
	if err != nil {
		fail(err.Error())
	}
	cnew, err := bux.Get(current)
	if err != nil {
		fail(err.Error())
	}
	tnew, err := bux.Get(to)
	if err != nil {
		fail(err.Error())
	}
	fmt.Printf("%s%s%s bux from %s%s %s(%v) %sto %s%s %s(%v)%s\n",
		c.B3, astring, c.Y, c.B2, current, c.R, cnew, c.Y, c.B2, to, c.G, tnew, c.X)
}

func all() {
	a, err := bux.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func list() {
	users := s.GetUserNames()
	for _, user := range users {
		n, _ := bux.Get(user)
		fmt.Printf("%20s %v\n", user, n)
	}
}

func show(u string) {
	if u == "all" {
		all()
	} else if u == "list" {
		list()
	} else {
		n, err := bux.Get(u)
		if err != nil {
			fmt.Println(err)
		} else {
			who := u + " has"
			if u == current {
				who = "You have"
			}
			plural := "SkilBux"
			if n == 1 {
				plural = "SkilBuk"
			}
			fmt.Printf("%s%s %s%v%s %s.%s\n", c.Y, who, c.M, n, c.Y, plural, c.X)
		}
	}
}

func freebuk() {
	if bux.Last(current) > bux.FREESECONDS {
		fmt.Println(c.Y + "Hey looks like you are elibible for a free SkilBuk!" + c.X)
		if !human.Confirm() {
			fail(bux.NOTHUMAN)
		}
		err := bux.Adjust(1, current, fmt.Sprintf("free every %v", bux.FREESECONDS))
		if err != nil {
			fmt.Println(err)
		}

	}
}

func main() {
	u, _ := user.Current()
	current = u.Username
	if current != "admin" {
		freebuk()
	}
	argc := len(os.Args)
	switch {
	case argc == 1:
		show(current)
	case argc == 2:
		show(os.Args[1])
	case argc == 3:
		fail(bux.NEEDSCOMMENT)
	case argc > 3:
		transfer()
	}
}

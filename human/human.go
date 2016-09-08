/* The human package helps make sure the person involved is alive.
Sometimes you need to do more than just test for an interactive shell.
*/
package human

import (
	"fmt"
	"github.com/skilstak/go/choice"
	c "github.com/skilstak/go/colors"
	"github.com/skilstak/go/input"
	"os/user"
	"regexp"
)

var Prompt string = c.B2 + "First let's see if you're human." + c.X

// Question and answer as regular expression string
type Challenge struct {
	Q string
	A string
}

type ChallengeSet struct {
	Challenges []Challenge
}

// Is this a user on the current system
func IsUser(u string) bool {
	// TODO add rules eliminating pseudo and other users
	_, err := user.Lookup(u)
	return err == nil
}

// Utility function to genericize when needed (i.e. choice.Interfaces())
func (set ChallengeSet) Interfaces() []interface{} {
	chinterfaces := make([]interface{}, len(set.Challenges))
	for i, d := range set.Challenges {
		chinterfaces[i] = d
	}
	return chinterfaces
}

// The specific ChallengeSet that will be used (with defaults)
// Set this to the ChallengeSet of questions and answers you prefer.
var Challenges = ChallengeSet{
	Challenges: []Challenge{
		{
			Q: "What is the air speed velocity of an unladen swallow?",
			A: "(?i)27|don't know|african|european",
		},
		{
			Q: "What type of animal is the Go mascot?",
			A: "(?i)gopher|mammal",
		},
		{
			Q: "In which town was SkilStak first started?",
			A: "(?i)cornelius",
		},
		{
			Q: "What is the best operating system for hosting?",
			A: "(?i)linux|unix|ubuntu|debian|red hat",
		},
		{
			Q: "What is the best editor every created by humans?",
			A: "(?i)vi",
		},
		{
			Q: "What is Mr. Rob's favorite programming language?",
			A: "(?i)go",
		},
		{
			Q: "Is Mr. Rob opinionated?",
			A: "(?i)yes|are you kidding|of course|duh|yep",
		},
		{
			Q: "What is the number 10 in binary?",
			A: "(?i)^2$|two",
		},
	},
}

// Prompts to answer a random question to prove interactive (human) use.
// (Requires command prompt.)
func Confirm() bool {
	fmt.Println(c.Y + Prompt + c.X)
	challenge := choice.Interfaces(Challenges.Interfaces()).(Challenge)
	answer := input.Ask(c.Y + challenge.Q + " " + c.B3)
	match, _ := regexp.MatchString(challenge.A, answer)
	if match {
		fmt.Println(c.M + "YES!" + c.X)
	} else {
		fmt.Println(c.R + "NO!" + c.X)
	}
	return match
}

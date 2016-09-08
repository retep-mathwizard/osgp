/* Package choice helps choose from a list of different types.
 */

package choice

import (
	"math/rand"
	"time"
)

func Strings(list []string) string {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

// Like the Python equivalent
func Choice(list []string) string {
	return Strings(list)
}

func Interfaces(list []interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

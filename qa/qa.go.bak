// The QA module handles questions and answers
// contained in a JSON file list of QA pairs
// where the question ("Q") is simply a string
// and the answer ("A") is a regular expression
// that is matched will produce a correct (true)
// match.
package qa

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

type QA struct {
	Q string `json:"q"`
	A string `json:"a"`
}

var List []QA

func Load(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	json.Unmarshal(file, &List)
	return nil
}

func Check(q QA, a string) bool {
	isMatch, _ := regexp.MatchString(q.A, a)
	return isMatch
}

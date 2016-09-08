package qa_test

import (
	"fmt"
	"github.com/skilstak/go/qa"
	"github.com/skilstak/go/settings"
	"path/filepath"
)

/*
Assuming you have at least the following at /var/lib/skilstak/.qa.json:

  [
    {
      "q": "What is the air speed velocity of an unladen swallow?",
      "a": "(?i)24|don't know|african|european"
    },
    ...
  ]

Note that the regular expression is Go style but keep them
compatible with other regx engines for portability across
languages.
*/
func Example() {
	path := filepath.Join(settings.DataDir, ".qa.json")
	fmt.Println(path)
	qa.Load(path)
	one := qa.List[0]
	fmt.Println(one.Q)
	fmt.Println(one.A)
	fmt.Println(qa.Check(one, "blah"))
	fmt.Println(qa.Check(one, "24"))
	fmt.Println(qa.Check(one, "11"))
	fmt.Println(qa.Check(one, "   I don't know that"))

	//Output:
	///var/lib/skilstak/.qa.json
	//What is the air speed velocity of an unladen swallow?
	//(?i)24|11|don't know|african|european
	//false
	//true
	//true
	//true
}

package ports_test

import (
	"fmt"
	"github.com/skilstak/go/ports"
)

func ExampleWww() {
	port, _ := ports.Www()
	fmt.Println(port)
	//Output: 8501
}

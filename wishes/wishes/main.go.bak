package main

import (
	"fmt"
	"os"
	w "github.com/skilstak/go/wishes"
	"github.com/skilstak/go/colors"
)

func main() {
	fmt.Print(colors.CL)
	if len(os.Args) == 2 {
		w.WishList(os.Args[1])
	} else if _, err := os.Stat(os.Getenv("HOME") + "/.wishes"); os.IsNotExist(err) {
		fmt.Print(colors.CL)
		w.WishMake(w.WishTake())
	} else {
		w.WishPrompt()
	}

}

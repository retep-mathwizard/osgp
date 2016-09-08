package wishes

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"github.com/skilstak/go/input"
	c "github.com/skilstak/go/colors"
)

const prefix string = "/pretend"

func WishMake(a []string) {
	f, err := os.Create(os.Getenv("HOME") + "/.wishes")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for i, _ := range a {
		io.WriteString(f, a[i] + "\n")
	}
	
}

func WishTake() []string {
	fmt.Print(c.M + "I am the Genie of the Shelf! " + c.X)
	fmt.Println(c.Y + "Tell me three items (under $30 each)" + c.X)
	fmt.Print(c.Y + "that you would like to see on the " + c.X)
	fmt.Println(c.R +"Red Shelf " + c.Y + "and perhaps your wishes" + c.X)
	fmt.Println(c.Y + "will be granted:" + c.X)
	a := make([]string, 3)
	for i, _ := range a {
		wish := input.Ask(c.R + strconv.Itoa(i+1) + ". " + c.B3)
		a[i] = wish
	}
	fmt.Println(c.Y + "Your wishes have been saved." + c.X)
	fmt.Println(c.Y + "Type the word " + c.Cyan + "wishes" + c.Y + " to see or change your wishes.")
	return a
}

func WishPrompt() {
	f, err := os.Open(os.Getenv("HOME") + "/.wishes")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	fmt.Println(c.Y + "You currently are wishing for these three things:")
	wishcount := 0
	for scanner.Scan() {
		wishcount += 1
		if wishcount >= 4 {
			fmt.Println(c.R + "HEY! Don't add more than three wishes to your wishfile! This wish won't be counted!")
		}
		num := strconv.Itoa(wishcount)
		fmt.Println(c.Red + num + ". " + scanner.Text())
	}
	fmt.Println(c.Y + "Would you like to change your three wishes?" + c.X)
	answer := input.Ask("> " + c.B3)
	if answer == "yes" || answer == "y" || answer == "Yes" {
		WishMake(WishTake())
	} else {
		fmt.Println(c.Y + "See you next time, then.")
	}

}

func WishList(n string) {
	if _, err := os.Stat("/home/" + n + "/.wishes"); os.IsNotExist(err) {
		fmt.Println(c.M + os.Args[1] + c.Y + " doesn't have any wishes yet.")
	} else {
		f, err := os.Open("/home/" + n + "/.wishes")                                                                                        
		if err != nil {                                                                                                                          
			fmt.Println(err)                                                                                                                 
		}                                                                                                                                        
		defer f.Close()                                                                                                                          
	        scanner := bufio.NewScanner(f)                                                                                                                                                                         
	        wishcount := 0                                                                                                                           
		for scanner.Scan() {                                                                                                                     
			wishcount += 1                                                                                                                   
			if wishcount >= 4 {                                                                                                              
				fmt.Println(c.R + "Looks like someone's trying to cheat the system! This wish coesn't count.")                 
			}                                                                                                                                
			num := strconv.Itoa(wishcount)                                                                                                   
			fmt.Println(c.Red + num + ". " + scanner.Text())                                                                                 
		}
	
	}

}
	



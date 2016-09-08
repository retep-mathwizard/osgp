package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt(p string) (string, error) {
	fmt.Print(p)
	reader := bufio.NewReader(os.Stdin)
	t, err := reader.ReadString('\n')
	if err != nil {
		return t, err
	}
	t = strings.TrimSpace(t)
	return t, nil
}

func Ask(q string) string {
	t, err := Prompt(q)
	if err != nil {
		return ""
	}
	return t
}

func HasStdin() (b bool) {
	stat, _ := os.Stdin.Stat()
	b = (stat.Mode() & os.ModeCharDevice) == 0
	return
}

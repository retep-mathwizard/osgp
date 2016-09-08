package sec

import (
	"fmt"
	"syscall"
	"testing"
)

func TestAsAdmin(t *testing.T) {
	before := syscall.Geteuid()
	println(before)
	hello := func() error {
		fmt.Println("Hello World")
		return nil
	}
	err := AsAdmin(hello)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	after := syscall.Geteuid()
	println(after)
}

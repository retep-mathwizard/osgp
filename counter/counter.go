package counter

import (
	"bufio"
	"os"
	"strconv"
)

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func ReadInt(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	s.Scan()
	number := s.Text()
	n, err := strconv.Atoi(number)
	if err != nil {
		return n, err
	}
	defer f.Close()
	return n, nil
}

func WriteInt(file string, a int) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	f.WriteString(strconv.Itoa(a))
	defer f.Close()
	return nil
}

func AddInt(file string, a interface{}) error {
	var err error
	var i int
	switch v := a.(type) {
	case string:
		if a, err = strconv.Atoi(v); err != nil {
			return err
		}
	}
	if i, err = ReadInt(file); err != nil {
		return err
	}
	return WriteInt(file, i+a.(int))
}

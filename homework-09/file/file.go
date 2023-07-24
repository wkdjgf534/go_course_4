package main

import (
	"fmt"
	"io"
	"os"
)

// WriteToFile - write only strings into a file
func WriteToFile(w io.Writer, data ...any) error {
	for _, d := range data {
		switch v := d.(type) {
		case string:
			_, err := w.Write([]byte(v))
			if err != nil {
				return err
			}
		default:
			continue
		}
	}
	return nil
}

func main() {
	f, err := os.Create("./test.txt")
	if err != nil {
		fmt.Printf("you have got an error: %s", err)
	}
	defer f.Close()

	str1 := "Hello"
	str2 := "World"
	num1 := 100
	num2 := 101
	bool1 := true
	bool2 := false

	err = WriteToFile(f, str1, num1, str2, num2, bool1, bool2)
	if err != nil {
		fmt.Printf("you have got an error: %s", err)
	}
}

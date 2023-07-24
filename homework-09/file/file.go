package main

import (
	"fmt"
	"io"
	"os"
)

// WriteToFile - функция записи в файл
func WriteToFile(w io.Writer, data ...interface{}) {
	for _, d := range data {
		switch v := d.(type) {
		case string:
			str := v
			w.Write([]byte(str))
		default:
			continue
		}
	}
}

func main() {
	f, err := os.Create("./test.txt")
	if err != nil {
		fmt.Errorf("error", err)
	}
	defer f.Close()

	str1 := "Hello"
	str2 := "World"
	num1 := 100
	num2 := 101

	fmt.Println("The eldest person", WriteToFile(f, str1, num1, str2, num2))
}

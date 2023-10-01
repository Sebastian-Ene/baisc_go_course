package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)
	argFileName := getFileName()
	fmt.Println(argFileName)
	file, err := os.Open(argFileName)
	if err != nil {
		fmt.Println("Can't open file. Error: " + err.Error())
		return
	}
	io.Copy(os.Stdout, file)

}

func getFileName() string {
	if len(os.Args) < 2 {
		fmt.Println("No argument was passed in!")
		return ""
	}
	return os.Args[1]
}

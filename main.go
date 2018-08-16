package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "filehandles.test"

	data := []byte("hello\ngo\n")
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
	defer os.Remove(filename)

	files := make([]*os.File, 0)
	errorCount := 0
	for done := false; !done; {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file", filename, err, errorCount)
			errorCount += 1

			if errorCount > 5 {
				fmt.Println("Holding", len(files), "files open")
				fmt.Println("Press <Enter> to quit")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				done = true
			}
		} else {
			files = append(files, f)
			defer f.Close()
			fmt.Print(".")
		}
	}
}
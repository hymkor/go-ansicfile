package main

import (
	"bufio"
	"fmt"
	"os"

	".."
)

func main() {
	fp, err := ansicfile.Open("test", "w")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	for i := 0; i < 5; i++ {
		fmt.Fprintf(fp, "test: %d\n", i)
	}
	fp.Close()

	fp, err = ansicfile.Open("test", "r")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fp.Close()
}

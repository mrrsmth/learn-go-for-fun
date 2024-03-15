package txtread

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Cr() {
	f, err := os.Create("fun/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("old falcon\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fr, err := os.Open("fun/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = fr.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(fr)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text())  // token in unicode-char
		fmt.Println(scanner.Bytes()) // token in bytes
	}

	content, err := os.ReadFile("fun/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	fmt.Println("done")
	fmt.Println(lines)
}

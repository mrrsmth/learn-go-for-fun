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

	fr, err3 := os.Open("fun/data.txt")
	if err != nil {
		log.Fatal(err3)
	}
	defer func() {
		if err = fr.Close(); err != nil {
			log.Fatal(err3)
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

	// str
	strOne := "one"
	strTwo := "one"
	// output := strings.Compare(strOne, strTwo)
	// if output == -1 {
	// 	fmt.Println("not found")
	// } else {
	// 	fmt.Println("found")
	// }
	fmt.Println(foundString(strOne, strTwo))
	check := foundString(strOne, strTwo)
	if check {
		fmt.Println("found again")
	} else {
		fmt.Println("not found again")
	}

}

func foundString(strOne, strTwo string) bool {
	output := strings.Compare(strOne, strTwo)
	if output == -1 {
		return false
	} else {
		return true
	}
}

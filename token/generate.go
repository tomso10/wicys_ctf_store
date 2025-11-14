package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	amount       int    = 100
	token_length int    = 3
	_type        string = "purple"
)

func main() {
	// read file
	file, err := os.Open("wordlist.txt")
	if err != nil {
		panic(err)
	}

	var wordlist []string

	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())

		if len(word) >= 2 && len(word) <= 6 {
			wordlist = append(wordlist, word)
		}
	}

	// close file
	err = file.Close()
	if err != nil {
		panic(err)
	}

	length := len(wordlist)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < amount; i++ {
		parts := make([]string, token_length)

		for j := 0; j < token_length; j++ {
			parts[j] = wordlist[r1.Intn(length)]
		}

		token := strings.Join(parts, "-")

		fmt.Printf("  - value: %v\n    type: %v\n", token, _type)

	}

}

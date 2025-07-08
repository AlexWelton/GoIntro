package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	name := getName()
	fmt.Println(getResponse(name))
}

func palindrome(name string) bool {
	cleanedName := strings.Join(strings.Fields(strings.ToLower(name)), "")
	for i := 0; i < len(cleanedName)/2; i++ {
		if cleanedName[i] != cleanedName[len(cleanedName)-i-1] {
			return false
		}
	}
	return true
}

func getResponse(name string) string {
	if name == "" {
		return "Okay, no greeting for you"
	} else {
		return greeting(name)
	}
}

func greeting(name string) string {
	designers := [3]string{"Robert Greisemer", "Rob Pike", "Ken Thompson"}

	greeting := "Hello, %s."

	//Designers
	for _, designer := range designers {
		if name == designer {
			greeting += " Thanks for creating me!"
		}
	}

	//Palindrome
	if palindrome(name) {
		greeting += " Cool, a palindromic name!"
	}

	//Overlength
	if len(name) > 20 {
		name = name[:20]
		greeting += ".. Wow that name's too long for me!"
	} else {
		//Extract first name
		subNames := strings.Fields(name)
		name = subNames[0]
	}

	return fmt.Sprintf(greeting, name)
}

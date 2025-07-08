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

func isDesigner(name string) bool {
	designers := [3]string{"Robert Greisemer", "Rob Pike", "Ken Thompson"}
	for _, designer := range designers {
		if name == designer {
			return true
		}
	}

	return false
}

func shortenName(name string) string {
	return name[:20]
}

func getFirstName(name string) string {
	subNames := strings.Fields(name)
	return subNames[0]
}

func greeting(name string) string {
	greeting := "Hello, %s."
	var messageName string

	if len(name) > 20 {
		messageName = shortenName(name)
		greeting += ".. Wow that name's too long for me!"
	} else {
		messageName = getFirstName(name)
	}

	if isDesigner(name) {
		greeting += " Thanks for creating me!"
	}

	if palindrome(name) {
		greeting += " Cool, a palindromic name!"
	}

	return fmt.Sprintf(greeting, messageName)
}

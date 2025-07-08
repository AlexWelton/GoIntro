package main

import "testing"

var greetingTests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice."},
	{"Bob", "Hello, Bob. Cool, a palindromic name!"},
	{"Robert Greisemer", "Hello, Robert. Thanks for creating me!"},
	{"Rob Pike", "Hello, Rob. Thanks for creating me!"},
	{"Prince Rupert Grayson Robertson the Third", "Hello, Prince Rupert Grayso... Wow that name's too long for me!"},
	{"A man a plan a canal panama", "Hello, A man a plan a canal... Wow that name's too long for me! Cool, a palindromic name!"},
	{"Haskell Curry", "Hello, Haskell."},
	{"Abi Iba", "Hello, Abi. Cool, a palindromic name!"},
}

func TestGreeting(t *testing.T) {
	for _, test := range greetingTests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}

var responseTests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice."},
	{"", "Okay, no greeting for you"},
	{" Tommy", "Hello, Tommy."},
}

func TestGetResponse(t *testing.T) {
	for _, test := range responseTests {
		result := getResponse(test.name)
		if result != test.expected {
			t.Errorf("incorrect response, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}

var firstNameTests = []struct {
	name     string
	expected string
}{
	{"John Doe", "John"},
	{"Johndoe", "Johndoe"},
}

func TestGetFirstName(t *testing.T) {
	for _, test := range firstNameTests {
		result := getFirstName(test.name)
		if result != test.expected {
			t.Errorf("incorrect response, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}

var shortenNameTests = []struct {
	name     string
	expected string
}{
	{"Prince Rupert Grayson Robertson the Third", "Prince Rupert Grayso"},
	{"A man a plan a canal panama", "A man a plan a canal"},
}

func TestShortenName(t *testing.T) {
	for _, test := range shortenNameTests {
		result := shortenName(test.name)
		if result != test.expected {
			t.Errorf("incorrect response, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}

var isDesignerTests = []struct {
	name     string
	expected bool
}{
	{"Bob", false},
	{"Robert Greisemer", true},
	{"Rob Pike", true},
}

func TestIsDesigner(t *testing.T) {
	for _, test := range isDesignerTests {
		result := isDesigner(test.name)
		if result != test.expected {
			t.Errorf("incorrect response, for name %s: got %t, expected %t", test.name, result, test.expected)
		}
	}
}

var palindromeTests = []struct {
	name     string
	expected bool
}{
	{"Bob", true},
	{"A man a plan a canal panama", true},
	{"Abi Iba", true},
	{"Pally Drome", false},
}

func TestPalindrome(t *testing.T) {
	for _, test := range palindromeTests {
		result := palindrome(test.name)
		if result != test.expected {
			t.Errorf("incorrect response, for name %s: got %t, expected %t", test.name, result, test.expected)
		}
	}
}

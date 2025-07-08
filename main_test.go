package main

import "testing"

var greetingTests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice."},
	{"Bob", "Hello, Bob. Cool, a palindromic name!"},
	{"Robert Greisemer", "Hello, Robert. Thanks for creating me!"},
	{"Prince Rupert Grayson Robertson the Third", "Hello, Prince Rupert Grayso... Wow that name's too long for me!"},
	{"A man a plan "},
	{"Haskell Curry", "Hello, Haskell."},
	{"Abi Iba", "Hello, Abi. Cool, a palindromic name!"},
}

var responseTests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice."},
	{"", "Okay, no greeting for you"},
	{" Tommy", "Hello, Tommy."},
}

func TestGreeting(t *testing.T) {
	for _, test := range greetingTests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}

func TestGetResponse(t *testing.T) {
	for _, test := range responseTests {
		result := getResponse(test.name)
		if result != test.expected {
			t.Errorf("incorrect response, for name %s: got %s, expected %s", test.name, result, test.expected)
		}
	}
}

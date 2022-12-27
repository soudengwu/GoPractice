package hello

import (
	"fmt"
)

// Go Placeholder strings. https://pkg.go.dev/fmt#hdr-Printing

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanishLanguageText = "Spanish"
const englishLanguageText = "English"
const frenchLanguageText = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case frenchLanguageText:
		prefix = frenchHelloPrefix
	case spanishLanguageText:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Var() {
	var b string
	var a string = "Hello a"
	fmt.Println(a)
	b = "Hello B"
	fmt.Println(b)

	c := "Hello C"
	fmt.Println(c)
}

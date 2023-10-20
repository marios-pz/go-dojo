package main

import "fmt"

var LangTable map[string]string = map[string]string{
	"EN": "Hello dojo",
	"JP": "私の道場へようこそ",
}

func Hello(msg string, language string) string {
	if value, ok := LangTable[language]; ok {
		if msg == "" {
			return value
		}
	}

	return msg
}

func main() {
	fmt.Println(Hello("", "EN"))
}

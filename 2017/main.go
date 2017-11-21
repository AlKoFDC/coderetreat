package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

const (
	blank      = " "
	pipe       = "|"
	underscore = "_"
)

func ToLCD(number int) string {
	switch number {
	case 2:
		return `
 _
 _|
|_
`

	}
	return ""
}

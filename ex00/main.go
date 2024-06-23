package main

import {
	"ex00/vendor/ft"
	"ex00/vendor/piscine"
}

func main() {
	alphabet := piscine.GenerateAlphabet()
    for _, char := range alphabet {
        ft.PrintRune(char)
    }
}

package main

import "fmt"
import "unicode/utf8"

func main() {
  fmt.Print("What is the input string? ")
  var input string
  fmt.Scanln(&input)
  var characters int = utf8.RuneCountInString(input)
  fmt.Println(input, "has", characters, "characters.")
}

package main

import "fmt"

func main() {
  fmt.Println("What is the quote? ")
  var quote string
  fmt.Scanln(&quote)
  fmt.Println("Who said it? ")
  var author string
  fmt.Scanln(&author)
  fmt.Println(author + " says, " + "\"" + quote + "\"")
}

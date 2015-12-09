package main

import "fmt"

func main() {
  fmt.Print("What is your name? ")
  var name string
  fmt.Scanln(&name)
  var greeting string = "Hello, " + name + ", nice to meet you!"
  fmt.Println(greeting)
}

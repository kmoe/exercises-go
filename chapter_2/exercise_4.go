package main

import "fmt"

func main() {
  fmt.Print("Enter a noun: ")
  var noun string
  fmt.Scanln(&noun)
  fmt.Print("Enter a verb: ")
  var verb string
  fmt.Scanln(&verb)
  fmt.Print("Enter an adjective: ")
  var adjective string
  fmt.Scanln(&adjective)
  fmt.Print("Enter an adverb: ")
  var adverb string
  fmt.Scanln(&adverb)

  fmt.Printf("Do you %v your %v %v %v? That's hilarious!", verb, adjective, noun, adverb)
}

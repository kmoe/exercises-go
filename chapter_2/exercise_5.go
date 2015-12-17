package main

import "fmt"
import "strconv"

func main() {
  fmt.Print("What is the first number? ")
  var first string
  fmt.Scan(&first)
  fmt.Print("What is the second number? ")
  var second string
  fmt.Scan(&second)

  var firstNumber, _ = strconv.ParseFloat(first, 64)
  var secondNumber, _ = strconv.ParseFloat(second, 64)

  // if firstNumber, err := strconv.ParseFloat(first, 64); err == nil {

  // }

  fmt.Println(first, "+", second, "=", firstNumber + secondNumber,
              "\n" + first, "-", second, "=", firstNumber - secondNumber,
              "\n" + first, "*", second, "=", firstNumber * secondNumber,
              "\n" + first, "/", second, "=", firstNumber / secondNumber)

}

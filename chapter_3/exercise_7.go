package main

import "fmt"
import "strconv"

func main() {
  const conversionFactor float64 = 0.09290304

  fmt.Print("What is the length of the room in feet? ")
  var lengthString string
  fmt.Scan(&lengthString)
  fmt.Print("What is the width of the room in feet? ")
  var widthString string
  fmt.Scan(&widthString)
  fmt.Println("You entered dimensions of", lengthString, "feet and", widthString, "feet.")

  fmt.Println("The area is")

  var length, _ = strconv.ParseFloat(lengthString, 64)
  var width, _ = strconv.ParseFloat(widthString, 64)

  var areaInSquareFeet float64 = length * width
  var areaInSquareMetres = areaInSquareFeet * conversionFactor

  fmt.Println(areaInSquareFeet, "square feet")
  fmt.Println(areaInSquareMetres, "square metres")
}

package main

import "fmt"
import "strconv"

func main() {
  const conversionFactor float64 = 0.09290304

  fmt.Print("What is the length of the room in feet? ")
  var length string
  fmt.Scan(&length)
  fmt.Print("What is the width of the room in feet? ")
  var width string
  fmt.Scan(&width)
  fmt.Println("You entered dimensions of", length, "feet and", width, "feet.")

  fmt.Println("The area is")

  var lengthNumber, _ = strconv.ParseFloat(length, 64)
  var widthNumber, _ = strconv.ParseFloat(width, 64)

  var areaInSquareFeet float64 = lengthNumber * widthNumber
  var areaInSquareMetres = areaInSquareFeet * conversionFactor

  fmt.Println(areaInSquareFeet, "square feet")
  fmt.Println(areaInSquareMetres, "square metres")
}

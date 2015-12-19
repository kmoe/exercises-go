package main

import "fmt"
import "strconv"
import "math"

func main() {
  const squareFeetPerGallon float64 = 350

  fmt.Print("What is the length of the room in feet? ")
  var lengthString string
  fmt.Scan(&lengthString)
  fmt.Print("What is the width of the room in feet? ")
  var widthString string
  fmt.Scan(&widthString)

  var length, _ = strconv.ParseFloat(lengthString, 64)
  var width, _ = strconv.ParseFloat(widthString, 64)

  var area float64 = length * width

  var gallons float64 = area / squareFeetPerGallon

  var gallonsNeeded float64 = math.Ceil(gallons)

  fmt.Println("You will need to purchase", gallonsNeeded, "gallons of paint to cover", area, "square feet.")
}

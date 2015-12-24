package main

import "fmt"
import "strconv"
import "math"

func main() {
  const squareFeetPerGallon = 350

  fmt.Print("What is the length of the room in feet? ")
  var lengthString string
  fmt.Scan(&lengthString)
  fmt.Print("What is the width of the room in feet? ")
  var widthString string
  fmt.Scan(&widthString)

  var length, _ = strconv.ParseFloat(lengthString, 64)
  var width, _ = strconv.ParseFloat(widthString, 64)

  var area = length * width

  var gallons = area / squareFeetPerGallon

  var gallonsNeeded = math.Ceil(gallons)

  fmt.Println("You will need to purchase", gallonsNeeded, "gallons of paint to cover", area, "square feet.")
}

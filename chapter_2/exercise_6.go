package main

import "fmt"
import "time"
import "strconv"

func main() {
  fmt.Print("What is your current age? ")
  var currentAgeStr string
  fmt.Scan(&currentAgeStr)
  fmt.Print("At what age would you like to retire? ")
  var retirementAgeStr string
  fmt.Scan(&retirementAgeStr)

  var currentAge, _ = strconv.ParseInt(currentAgeStr, 10, 64)
  var retirementAge, _ = strconv.ParseInt(retirementAgeStr, 10, 64)

  var yearsUntilRetirement = retirementAge - currentAge

  var currentYear = int64(time.Now().Year())
  var retirementYear = currentYear + yearsUntilRetirement

  fmt.Printf("You have %v years left until you can retire.\n", yearsUntilRetirement)
  fmt.Printf("It's %v, so you can retire in %v.", currentYear, retirementYear)
}

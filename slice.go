package main

import "fmt"

//Append Example

// func main() {
//    slice1 := []int{1,2,3}
//    slice2 := append(slice1, 4, 5)
//    fmt.Println(slice1, slice2)

//Copy Example

//func main() {
//  slice1 := []int{1,2,3}
//  slice2 := make([]int, 2)
//  copy(slice2, slice1)
//  fmt.Println(slice1, slice2)

//Map example
func main() {
//  elements := make(map[string]string)
//  elements["H"] = "Hydrogen"
//  elements["He"] = "Helium"
//  elements["Li"] = "Lithium"
//  elements["Be"] = "Beryllium"
//  elements["B"] = "Boron"
//  elements["C"] = "Carbon"
//  elements["N"] = "Nitrogen"
//  elements["O"] = "Oxygen"
//  elements["F"] = "Fluorine"
//  elements["Ne"] = "Neon"

elements := map[string]string{
  "H":  "Hydrogen",
  "He": "Helium",
  "Li": "Lithium",
  "Be": "Beryllium",
  "B":  "Boron",
  "C":  "Carbon",
  "N":  "Nitrogen",
  "O":  "Oxygen",
  "F":  "Fluorine",
  "Ne": "Neon",
}

  fmt.Println(elements["Li"])

}

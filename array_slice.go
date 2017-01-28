package main

import "fmt"

// func main() {
//  var x [5]int
//  x[4] = 100
//  fmt.Println(x)

func main() {
//  var x [5]float64
//  x[0] = 98
//  x[1] = 93
//  x[2] = 77
//  x[3] = 82
//  x[4] = 83

x := [6]float64{
  98,
  93,
  77,
  82,
  83,
//  100,
}

//  var total float64 = 0
//  for i := 0; i < 5; i++ {
//    total += x[i]
//  }
//  fmt.Println(total / 5)

// var total float64 = 0
// for i := 0; i < len(x); i++ {
//   total += x[i]
// }
// fmt.Println(total / float64(len(x)))

var total float64 = 0
for _, value := range x {
  total += value
}
fmt.Println(total / float64(len(x)))

}


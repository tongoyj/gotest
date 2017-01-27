package main

import "fmt"

func main() {
  var x string
  x = "first"
  fmt.Println(x)
  x += "second"  //same as x = x + "second"
  fmt.Println(x)

  var a string = "hello"
  var b string = "hello"
  fmt.Println(a == b)

  y :=7
  var w = "Jesus Saves"
  fmt.Println(w, y)
}

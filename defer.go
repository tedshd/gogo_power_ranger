package main

import "fmt"

func main() {
  // defer is LIFO
  fmt.Println("1")
  defer fmt.Println("2")
  fmt.Println("3")
  defer fmt.Println("4")
  return
  defer fmt.Println("5")
}

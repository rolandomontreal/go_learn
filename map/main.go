package main

import "fmt"

func main() {
  // var colors map[string]string

  // colors := make(map[string]string)
  
  // colors["red"] = "poop"
  
  // fmt.Println(colors["red"])
  
  // delete(colors, "red")
  // fmt.Println(colors["red"])

  colors := map[string]string{
    "red": "#FF0000",
    "green": "#00FF00",
  }

  for color, hex := range colors {
    fmt.Println(color, hex)
  }
}

package main

import (
	"fmt"
)

type logWriter struct{}

// func main() {
// 	resp, error := http.Get("https://google.com")

// 	if error != nil {
// 		fmt.Println("Error: ", error)
// 		os.Exit(1)
// 	}

// 	// Example 1
// bs := make([]byte, 99999)
// resp.Body.Read(bs)
// fmt.Println(string(bs))

// 	// Example 2
// 	// io.Copy(os.Stdout, resp.Body)

// 	lw := logWriter{}
// 	io.Copy(lw, resp.Body)
// }

func (l logWriter) Write(bs []byte) (int, error) {
  fmt.Println(string(bs))
  fmt.Println("Just wrote this many bytes: ", len(bs))
  return len(bs), nil
}
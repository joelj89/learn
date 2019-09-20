// Echo1 prints its command-line arguments.
package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func main() {
  start1 := time.Now()
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep+ os.Args[i]
    sep = " "
  }
  fmt.Println(s)
  fmt.Printf("Loop: %.2fs second(s)\n", time.Since(start1).Seconds())

  start2 := time.Now()
  fmt.Println(strings.Join(os.Args[1:], " "))
  fmt.Printf("Join: %.2fs second(s)\n", time.Since(start2).Seconds())
}

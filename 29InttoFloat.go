package main

import "fmt"

func main() {

      var i int = 7
      fmt.Printf("%v, %T\n", i, i)
      var j float32
      j = float32(i)
      fmt.Printf("%v, %T\n", j, j)
}

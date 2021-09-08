package main

import (
  "datastructure/exercise2/sequence"
  "fmt"
)

func main() {

  x := sequence.Numbers{20, 7, 8, 10, 2, 5, 6}

  fmt.Printf("%v\n", x.Exist([]int{7, 8}))     // true
  fmt.Printf("%v\n", x.Exist([]int{10, 2, 5})) //true
  fmt.Printf("%v\n", x.Exist([]int{10, 5}))    //false
  fmt.Printf("%v\n", x.Exist([]int{5, 6}))     //true
  fmt.Printf("%v\n", x.Exist([]int{6, 8}))     //false
  fmt.Printf("%v\n", x.Exist([]int{20}))
  fmt.Printf("%v\n", x.Exist([]int{6}))
  fmt.Printf("%v\n", x.Exist([]int{3}))
  fmt.Printf("%v\n", x.Exist([]int{8, 10, 2}))

}

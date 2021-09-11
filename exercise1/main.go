package main

import (
  "datastructure/exercise1/tree"
  "fmt"
)

func main() {

  root := tree.Node{
    Val: 6,
    Left: &tree.Node{
      Val:   7,
      Left:  tree.NewBothNil(10),
      Right: tree.NewBothNil(8),
    },
    Right: &tree.Node{
      Val:   100,
      Left:  nil,
      Right: tree.NewBothNil(14),
    },
  }

  fmt.Printf("%v\n", root.Find(6))   // true
  fmt.Printf("%v\n", root.Find(7))   // true
  fmt.Printf("%v\n", root.Find(10))  // true
  fmt.Printf("%v\n", root.Find(8))   // true
  fmt.Printf("%v\n", root.Find(100)) // true
  fmt.Printf("%v\n", root.Find(14))  // true

  fmt.Printf("%v\n", root.Find(9)) // false
  fmt.Printf("%v\n", root.Find(5)) // false

}

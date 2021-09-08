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

  fmt.Printf("%v\n", root.Find(6))
  fmt.Printf("%v\n", root.Find(7))
  fmt.Printf("%v\n", root.Find(8))
  fmt.Printf("%v\n", root.Find(9))
  fmt.Printf("%v\n", root.Find(10))
  fmt.Printf("%v\n", root.Find(14))

}

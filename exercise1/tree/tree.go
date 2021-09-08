package tree

type Node struct {
  Val   int
  Left  *Node
  Right *Node
}

func NewBothNil(x int) *Node {
  return &Node{Val: x, Left: nil, Right: nil}
}

func New(x int, left *Node, right *Node) *Node {
  return &Node{Val: x, Left: left, Right: right}
}

func (n Node) Find(x int) bool {
  if n.Val == x {
    return true
  }

  if n.Left != nil && n.Left.Find(x) {
    return true
  }

  if n.Right != nil && n.Right.Find(x) {
    return true
  }

  return false
}

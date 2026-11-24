package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Postorder returns left-right-root order iteratively, via the reversal
// trick: a preorder visiting right before left, reversed, is postorder.
func Postorder(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, n.Val)
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
	}
	for l, r := 0, len(out)-1; l < r; l, r = l+1, r-1 {
		out[l], out[r] = out[r], out[l]
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, nil, tn(2, tn(3, nil, nil), nil))
	fmt.Println(Postorder(root))
}

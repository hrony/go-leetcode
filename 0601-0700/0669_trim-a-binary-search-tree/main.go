package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 0}

	first.Left = &third
	first.Right = &second

	fmt.Println(trimBST(&first, 1, 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Val < L{
		return trimBST(root.Right,L,R)
	}
	if R < root.Val{
		return trimBST(root.Left,L,R)
	}

	root.Left = trimBST(root.Left,L,R)
	root.Right = trimBST(root.Right,L,R)
	return root
}*/
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val >= L && root.Val <= R {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}
	if L > root.Val {
		return trimBST(root.Right, L, R)
	}
	if R < root.Val {
		return trimBST(root.Left, L, R)
	}
	return nil
}

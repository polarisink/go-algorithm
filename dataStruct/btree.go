/**
 * @author: lqs
 * @date: 2022/1/6
 */
package main

import (
	"math"
)

var (
	res   = 0
	depth = 0
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func maxDepth(root *TreeNode) int {
	traverse(root)
	return res
}

//递归遍历树 
func traverse(root *TreeNode) {
	if root == nil {
		res = int(math.Max(float64(res), float64(depth)))
		return
	}
	depth++
	traverse(root.left)
	traverse(root.right)
	depth--
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := MaxDepth(root.left)
	rightMax := MaxDepth(root.right)
	return int(math.Max(float64(leftMax), float64(rightMax))) + 1
}


func main() {

}

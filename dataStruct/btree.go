/**
 * @author: lqs
 * @date: 2022/1/6
 */
package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func main() {

}

//递归遍历树
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	traverse(root.left)
	traverse(root.right)
}

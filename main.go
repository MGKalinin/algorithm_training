package main

import (
	"fmt"
	"runtime"
)

// 653. Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/?envType=problem-list-v2&envId=binary-search-tree

func main() {

	// Конструируем бинарное дерево из среза root = [5,3,6,2,4,null,7]
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}
	k := 9
	//Output: true

	fmt.Println(findTarget(root, k))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {

}

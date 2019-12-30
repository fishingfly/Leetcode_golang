package problems

import "fmt"

//Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func LevelOrder(root *TreeNode) {
	levelOrder(root)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var nodeList []*TreeNode
	var result [][]int
	var ceng []int
	nodeList = append(nodeList, root)
	count := 0
	nextNum := 0
	currentNum := 1
	for len(nodeList) != 0 {
		count++
		if nodeList[0] != nil {
			ceng = append(ceng, nodeList[0].Val)
			nodeList = append(nodeList, nodeList[0].Left)
			nodeList = append(nodeList, nodeList[0].Right)
			nextNum += 2
		}
		if  count == currentNum { //分层触发条件
			count = 0
			if len(ceng) == 0 {
				break
			}
			result = append(result, ceng)
			ceng = ceng[len(ceng) : ]
			currentNum ,nextNum = nextNum, 0
		}
		nodeList = nodeList[1:] //去掉头部
	}
	fmt.Println(result)
	return result
}
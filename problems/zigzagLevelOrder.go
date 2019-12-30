package problems

import "fmt"

func ZigzagLevelOrder(root *TreeNode) {
	zigzagLevelOrder(root)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	zhengxiang := true
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
			if !zhengxiang {
				for i, j := 0, len(ceng)-1; i < j; i, j = i+1, j-1 {
					ceng[i], ceng[j] = ceng[j], ceng[i]
				}
			}
			result = append(result, ceng)
			ceng = ceng[len(ceng) : ]
			currentNum ,nextNum = nextNum, 0
			zhengxiang = !zhengxiang
		}
		nodeList = nodeList[1:] //去掉头部
	}
	fmt.Println(result)
	return result
}

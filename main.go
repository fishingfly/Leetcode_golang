package main

import (
	"leetcode/problems"
)

//func main() {
//	fmt.Println(problems.NumDecodings("00"))
//	fmt.Println(problems.NumDecodings("10"))
//	fmt.Println(problems.NumDecodings("226"))
//	fmt.Println(problems.NumDecodings("27"))
//	fmt.Println(problems.NumDecodings("01"))
//	fmt.Println(problems.NumDecodings("230"))
//	fmt.Println(problems.NumDecodings("4757562545844617494555774581341211511296816786586787755257741178599337186486723247528324612117156948"))
//}

//func main() {
//	left2 := &problems.TreeNode{6,nil,nil}
//	right2 := &problems.TreeNode{6,nil,nil}
//	left1 := &problems.TreeNode{5,left2,nil}
//	right1 := &problems.TreeNode{5,nil,right2}
//	left := &problems.TreeNode{-57,left1,nil}
//	right := &problems.TreeNode{-57,nil,right1}
//	root := &problems.TreeNode{4,left,right}
//	fmt.Println(problems.IsSymmetric(root))
//}

//func main() {
//	nums := problems.SequentialDigits(1000,13000)
//	for i := 0; i < len(nums); i++ {
//		fmt.Println(nums[i])
//	}
//}

//func main() {
//	left2 := &problems.TreeNode{7,nil,nil}
//	right2 := &problems.TreeNode{6,nil,nil}
//	left1 := &problems.TreeNode{4,left2,nil}
//	right1 := &problems.TreeNode{5,nil,right2}
//	left := &problems.TreeNode{-57,left1,nil}
//	right := &problems.TreeNode{-57,nil,right1}
//	root := &problems.TreeNode{4,left,right}
//	problems.LevelOrder(root)
//}

func main() {
	left2 := &problems.TreeNode{7,nil,nil}
	right2 := &problems.TreeNode{6,nil,nil}
	left1 := &problems.TreeNode{4,left2,nil}
	right1 := &problems.TreeNode{5,nil,right2}
	left := &problems.TreeNode{-56,left1,nil}
	right := &problems.TreeNode{-57,nil,right1}
	root := &problems.TreeNode{4,left,right}
	problems.ZigzagLevelOrder(root)
}
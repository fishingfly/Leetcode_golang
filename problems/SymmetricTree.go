package problems

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	return isSymmetric(root)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 层次遍历
	var s []*TreeNode=nil
	s = append(s, root)
	index := 0
	for {
		if len(s) == 0 {
			break
		}
		var str []string=nil
		k := 0
		for i := 0; i < int(math.Pow(2.0, float64(index))); i++ {
			if len(s) == 0 {
				break
			}
			node := s[0]
			if node != nil {
				s = append(s, node.Left)
				s = append(s, node.Right)
				str = append(str, strconv.Itoa(int(math.Abs(float64(node.Val)))))
			} else {
				str = append(str, "&")
				s = append(s, nil)
				s = append(s, nil)
				k += 2
			}
			if k ==  int(math.Pow(2.0, float64(index + 1))) {
				s = nil
				break
			}
			s = s[1:]
		}
		fmt.Println(str)
		if !isReverseString(str) {
			return false
		}
		index ++
	}
	return true
}

func inOrderTraversal(root *TreeNode) string{
	if root == nil {
		return "&"
	}
	result := strconv.Itoa(root.Val)
	if root.Left != nil {
		result = inOrderTraversal(root.Left) + result
	} else {
		result = "&" + result
	}
	if root.Right != nil {
		result += inOrderTraversal(root.Right)
	} else {
		result += "&"
	}
	return result
}

func isReverseString(s []string) bool {
	for from, to := 0, len(s)-1; from < to; from, to = from + 1, to - 1 {
		if s[from] == s[to] {
			continue
		} else {
			return false
		}
	}
	return true
}


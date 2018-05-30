package main

import "fmt"
import "strings"

// https://leetcode.com/problems/reverse-pairs/description/

type TNode struct {
	parent *TNode
	left *TNode
	right *TNode
	cntTree int
	cntThis int
	value int
}

func (n *TNode) insert(val int) {
	n.cntTree += 1
	if (val < n.value) {
		if n.left == nil  {
			n.left = &TNode{n, nil, nil, 1, 1, val}
		} else {
			n.left.insert(val)
		}
	} else if (val > n.value) {
		if n.right == nil {
			n.right = &TNode{n, nil, nil, 1, 1, val}
		} else {
			n.right.insert(val)
		}
	} else {
		n.cntThis += 1
	}
}

func (n *TNode) print(tab int) {
	fmt.Printf("%sval=%d, cnt=%d\n", strings.Repeat("\t", tab), n.value, n.cntTree)
	if n.left == nil {
		fmt.Printf("%snil\n", strings.Repeat("\t", tab+1))
	} else {
		n.left.print(tab+1)
	}
	if n.right == nil {
		fmt.Printf("%snil\n", strings.Repeat("\t", tab+1))
	} else {
		n.right.print(tab+1)
	}
}

func (n *TNode) largerThanVal(val int) int {
	if (n.value <= val) {
		if n.right == nil {
			return 0
		} else {
			return n.right.largerThanVal(val)
		}
	} else {
		ret := n.cntThis
		if n.right != nil {
			ret += n.right.cntTree
		}
		if n.left != nil {
			ret += n.left.largerThanVal(val)
		}
		return ret
	}
}

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	root := &TNode{nil, nil, nil, 1, 1, nums[0]}
	sum := 0
	for i := 1; i < len(nums); i++ {
		sum += root.largerThanVal(2*nums[i])
		root.insert(nums[i])
	}
	// root.print(0)
	return sum
}

func main() {
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
	var arr []int
	for i := 0; i < 50000; i++ {
		arr = append(arr, i)
	}
	fmt.Println(reversePairs(arr))
}
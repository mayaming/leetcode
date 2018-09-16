package main

import (
	"fmt"
	"math"
	"math/rand"
	s "strings"
	"time"
)

// https://leetcode.com/problems/reverse-pairs/description/

const maxLevel = 16
const possibility = 0.25

type Node struct {
	pts   []*Node
	value int
	repr  string
}

func makeNode(val int, level int) *Node {
	if level < 0 {
		level = randLevel()
	}
	pts := make([]*Node, level)
	repr := fmt.Sprintf("%d", val)
	return &Node{pts, val, repr}
}

func (n *Node) printAt(level int) string {
	if level <= len(n.pts) {
		return n.repr
	} else {
		return s.Repeat("-", len(n.repr))
	}
}

func (n *Node) at(i int) *Node {
	return n.pts[i-1]
}

func (n *Node) next(i int, nn *Node) {
	n.pts[i-1] = nn
}

type NodeList struct {
	head      *Node
	listLevel int
}

func makeList() *NodeList {
	return &NodeList{makeNode(math.MinInt32, maxLevel), maxLevel}
}

func (l *NodeList) print() {
	maxLevel := l.listLevel
	for i := maxLevel; i >= 1; i-- {
		s := ""

		for curNode := l.head; curNode != nil; curNode = curNode.pts[0] {
			if len(s) > 0 {
				s += "->"
			}
			s += curNode.printAt(i)
		}
		fmt.Println(s)
	}
}

func (l *NodeList) insertNode(targetNode *Node) {
	tnLevel := len(targetNode.pts)
	tnVal := targetNode.value
	if tnLevel > l.listLevel {
		l.listLevel = tnLevel
	}

	curNode := l.head
	curLevel := tnLevel

	for ; curLevel >= 1; curLevel-- {
		for nextNode := curNode.at(curLevel); nextNode != nil; {
			if curNode.value <= tnVal && tnVal <= nextNode.value {
				break
			} else {
				curNode = nextNode
				nextNode = curNode.at(curLevel)
			}
		}
		targetNode.next(curLevel, curNode.at(curLevel))
		curNode.next(curLevel, targetNode)
	}
}

func (l *NodeList) largerThan(val int) {
	curNode := l.head
	curLevel := l.listLevel

	for ; curLevel >= 1; curLevel-- {

	}
}

func randLevel() int {
	level := 1
	for ; rand.Float64() < possibility && level < maxLevel; level++ {
	}
	if level >= maxLevel {
		return maxLevel
	} else {
		return level
	}
}

func reversePairs(nums []int) int {
	l := makeList()
	for _, n := range nums {
		l.insertNode(makeNode(n, -1))
	}
	l.print()
	return 0
}

func main() {
	arr := make([]int, 50)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i, _ := range arr {
		arr[i] = r.Intn(50)
	}
	fmt.Println(reversePairs(arr))
	// fmt.Println(reversePairs([]int{1, 3, 2, 3, 1,}))
	// fmt.Println(reversePairs([]int{1, 3, 2, 3, 1, 9, 8, 6, 10, 5, 4, 3, 2, 7}))
	/*
		fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
		var arr []int
		for i := 0; i < 50000; i++ {
			arr = append(arr, i)
		}
		fmt.Println(reversePairs(arr))
	*/
}

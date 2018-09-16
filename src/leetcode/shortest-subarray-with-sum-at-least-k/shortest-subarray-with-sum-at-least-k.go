package main

// https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/description/

type Node struct {
	value int
	size  int
	prev  *Node
	next  *Node
}

func makeNode(val int) *Node {
	return &Node{val, 1, nil, nil}
}

func (thisNode *Node) mergeWith(otherNode *Node) {
	thisNode.value += otherNode.value
	thisNode.size += otherNode.size
	thisNode.next = otherNode.next
	if thisNode.next != nil {
		thisNode.next.prev = thisNode
	}
}

type NodeList struct {
	begin  *Node
	end    *Node
	curSum int
	num    int
	curMin int
	K      int
}

func (l *NodeList) append(val int) {
	pn := makeNode(val)
	if l.end == nil {
		l.begin = pn
		l.end = pn
		l.curSum = val
		l.num += 1
	} else {
		l.end.next = pn
		pn.prev = l.end
		l.end = pn
		l.curSum += val
		l.num += 1
	}
}

func (l *NodeList) compactFromEnd() {
	cur := l.end
	for cur != nil && cur.value <= 0 {
		prev := cur.prev
		if prev == nil {
			l.begin = nil
			l.end = nil
			l.curSum = 0
			l.num = 0
			cur = nil
		} else {
			prev.mergeWith(cur)
			cur = prev
		}
	}
	l.end = cur
}

func (l *NodeList) removeFromHead() {
	if l.curSum >= l.K && (l.curMin < 0 || l.num < l.curMin) {
		l.curMin = l.num
	}

	for cur := l.begin; cur != nil && l.curSum-cur.value >= l.K; {
		l.curSum -= cur.value
		l.num -= cur.size
		if l.curMin < 0 || l.num < l.curMin {
			l.curMin = l.num
		}
		cur = cur.next
		l.begin = cur
	}
}

func shortestSubarray(A []int, K int) int {
	l := &NodeList{nil, nil, 0, 0, -1, K}
	for i := 0; i < len(A); i++ {
		l.append(A[i])
		l.compactFromEnd()
		l.removeFromHead()
	}
	return l.curMin
}

func main() {
	println(shortestSubarray([]int{1}, 1))
	println(shortestSubarray([]int{1, 2}, 4))
	println(shortestSubarray([]int{2, -1, 2}, 3))
	println(shortestSubarray([]int{27, 20, 79, 87, -36, 78, 76, 72, 50, -26}, 453))
	println(shortestSubarray([]int{17, 85, 93, -45, -21}, 150))
	println(shortestSubarray([]int{45, 95, 97, -34, -42}, 21))
	println(shortestSubarray([]int{-28, 81, -20, 28, -29}, 89))
}

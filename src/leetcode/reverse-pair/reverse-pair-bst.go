package main

import "fmt"
import "strings"

// https://leetcode.com/problems/reverse-pairs/description/

type TNode struct {
	parent *TNode
	left   *TNode
	right  *TNode
	num    int
	value  int
	red    bool
}

func (n *TNode) insert(val int) *TNode {
	n.num += 1
	if val <= n.value {
		if n.left == nil {
			n.left = &TNode{n, nil, nil, 1, val, true}
			return n.left.rbFix()
		} else {
			return n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &TNode{n, nil, nil, 1, val, true}
			return n.right.rbFix()
		} else {
			return n.right.insert(val)
		}
	}
}

func (grandParent *TNode) uncle(parent *TNode) *TNode {
	if grandParent.left == parent {
		return grandParent.right
	} else if grandParent.right == parent {
		return grandParent.left
	} else {
		fmt.Println("parent is not a son of grand parent, something should be wrong")
		return nil
	}
}

func (parent *TNode) isLeftChild(child *TNode) bool {
	return child == parent.left
}

func (parent *TNode) isRightChild(child *TNode) bool {
	return child == parent.right
}

func (parent *TNode) resetNum() {
	parent.num = 1
	if parent.left != nil {
		parent.num += parent.left.num
	}
	if parent.right != nil {
		parent.num += parent.right.num
	}
}

func (child *TNode) adjustMyParent(originalChild *TNode) {
	if child.parent != nil {
		parent := child.parent
		if parent.left == originalChild {
			parent.left = child
		}
		if parent.right == originalChild {
			parent.right = child
		}
	}
}

func setPaternity(child *TNode, parent *TNode, left bool) {
	if child != nil {
		child.parent = parent
	}

	if parent != nil {
		if left {
			parent.left = child
		} else {
			parent.right = child
		}
	}
}

func leftLineRotate(targetNode *TNode, parent *TNode, grandParent *TNode) *TNode {
	/*
				祖父（黑）
				/
			  父（红）
			 /        \
		调整节点（红）  其他

		以父节点为支点右旋，然后将父节点染黑
				 父（红）
			    /      \
		调整节点（黑）   祖父（黑）
					   /
					其他
	*/

	setPaternity(parent.right, grandParent, true)
	setPaternity(parent, grandParent.parent, grandParent.parent != nil && grandParent.parent.left == grandParent)
	setPaternity(grandParent, parent, false)
	targetNode.red = false

	grandParent.resetNum()
	parent.resetNum()

	return parent.rbFix()
}

func rightLineRotate(targetNode *TNode, parent *TNode, grandParent *TNode) *TNode {
	/*
		祖父（黑）
				\
				父（红）
				/	 \
			其他	调整节点（红）

		以父节点为支点左旋，然后将父节点染黑
				父（红）
				/    \
		祖父（黑）   调整节点（黑）
				\
				其他
	*/

	setPaternity(parent.left, grandParent, false)
	setPaternity(parent, grandParent.parent, grandParent.parent != nil && grandParent.parent.left == grandParent)
	setPaternity(grandParent, parent, true)
	targetNode.red = false

	grandParent.resetNum()
	parent.resetNum()

	return parent.rbFix()
}

func leftTurnRotate(targetNode *TNode, parent *TNode, grandParent *TNode) *TNode {
	/*
				祖父（黑）
				/
			父（红）
			      \
				   调整节点（红）
				   /
				其他

		以父节点为支点左旋，然后继续调整
				祖父（黑）
				/
			调整节点（红）
			/
		父（红）
			  \
			  其他
	*/

	setPaternity(targetNode.left, parent, false)
	setPaternity(targetNode, grandParent, true)
	setPaternity(parent, targetNode, true)

	parent.resetNum()
	targetNode.resetNum()

	return parent.rbFix()
}

func rightTurnRotate(targetNode *TNode, parent *TNode, grandParent *TNode) *TNode {
	/*
				祖父（黑）
				        \
						父（红）
			        	/
				   调整节点（红）
							   \
							   其他

		以父节点为支点右旋，然后继续调整
				祖父（黑）
				        \
						调整节点（红）
								\
								父（红）
								/
							其他
	*/

	setPaternity(targetNode.right, parent, true)
	setPaternity(targetNode, grandParent, false)
	setPaternity(parent, targetNode, false)

	parent.resetNum()
	targetNode.resetNum()

	return parent.rbFix()
}

/*
红黑树的五个性质：
1）每个结点要么是红的，要么是黑的。
2）根结点是黑的。
3）每个叶结点，即空结点（NIL）是黑的。
4）如果一个结点是红的，那么它的俩个儿子都是黑的。
5）对每个结点，从该结点到其子孙结点的所有路径上包含相同数目的黑结点。
*/

func (insertedNode *TNode) rbFix() *TNode {
	// 插入根节点
	if insertedNode.parent == nil {
		insertedNode.red = false
		return insertedNode
	} else if insertedNode.parent.red == false {
		// 红黑树的条件都符合，不需要做什么
		return nil
	} else if insertedNode.parent.red == true { //父亲节点是红色
		parent := insertedNode.parent
		grandParent := parent.parent
		uncle := grandParent.uncle(parent)
		// 父亲节点和叔叔节点都是红色，将他们都变成黑色即可，但是祖父节点需要调整为红色，并针对祖父节点进行修正
		if uncle != nil && uncle.red == true {
			parent.red = false
			uncle.red = false
			grandParent.red = true
			return grandParent.rbFix()
		} else {
			if grandParent.isLeftChild(parent) {
				if parent.isLeftChild(insertedNode) {
					return leftLineRotate(insertedNode, parent, grandParent)
				} else {
					return leftTurnRotate(insertedNode, parent, grandParent)
				}
			} else {
				if parent.isRightChild(insertedNode) {
					return rightLineRotate(insertedNode, parent, grandParent)
				} else {
					return rightTurnRotate(insertedNode, parent, grandParent)
				}
			}
		}
	} else {
		return nil
	}
}

func (n *TNode) print(tab int) {
	fmt.Printf("%sval=%d, cnt=%d\n", strings.Repeat("\t", tab), n.value, n.num)
	if n.left == nil {
		fmt.Printf("%snil\n", strings.Repeat("\t", tab+1))
	} else {
		n.left.print(tab + 1)
	}
	if n.right == nil {
		fmt.Printf("%snil\n", strings.Repeat("\t", tab+1))
	} else {
		n.right.print(tab + 1)
	}
}

func (n *TNode) largerThanVal(val int) int {
	if n.value <= val {
		if n.right == nil {
			return 0
		} else {
			return n.right.largerThanVal(val)
		}
	} else {
		ret := 1
		if n.right != nil {
			ret += n.right.num
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

	root := &TNode{nil, nil, nil, 1, nums[0], true}
	root.rbFix()
	sum := 0
	for i := 1; i < len(nums); i++ {
		sum += root.largerThanVal(2 * nums[i])
		newRoot := root.insert(nums[i])
		if newRoot != nil {
			root = newRoot
		}
	}
	root.print(0)
	return sum
}

func main() {
	// fmt.Println(reversePairs([]int{1, 3, 2}))
	// fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	// fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
	// fmt.Println(reversePairs([]int{5, 4, 3, 2, 1}))
	// var arr []int
	// for i := 0; i < 50000; i++ {
	//   arr = append(arr, i)
	// }
	// fmt.Println(reversePairs(arr))
	fmt.Println(reversePairs([]int{12, 2, 1, 17, 19, 10, 5, 23, 7, 20, 10, 17, 22, 15, 9, 18, 12, 12, 16, 16, 17, 8, 11, 19, 2, 21, 5, 19, 22, 9, 17, 24, 8, 8, 16, 5, 2, 25, 1, 0, 3, 24, 25, 0, 11, 7, 19, 0, 5, 16, 17, 4, 19, 20, 20, 0, 14, 4, 16, 15, 11, 15, 20, 11, 17, 13, 3, 18, 12, 6, 10, 25, 12, 6, 18, 6, 19, 19, 18, 13, 21, 9, 17, 1, 1, 2, 10, 15, 24, 24, 22, 7, 10, 23, 15, 9, 1, 23, 22, 15, 3, 16, 23, 25, 8, 18, 0, 5, 1, 12, 9, 0, 25, 0, 13, 11, 22, 5, 3, 13, 10, 17, 14, 24, 23, 1, 8, 1, 21, 18, 2, 16, 21, 21, 5, 3, 19, 8, 23, 6, 6, 3, 2, 4, 13, 2, 4, 14, 9, 17, 23, 18, 4, 23, 5, 13, 25, 10, 9, 14, 3, 9, 11, 5, 14, 18, 0, 10, 13, 5, 19, 17, 24, 25, 4, 8, 16, 14, 3, 24, 18, 2, 17, 22, 4, 11, 18, 9, 9, 7, 10, 4, 24, 0, 7, 0, 6, 15, 18, 13, 14, 20, 22, 17, 22, 15, 17, 9, 10, 17, 13, 0, 22, 22, 23, 2, 21, 18, 6, 10, 10, 15, 14, 4, 4, 18, 21, 15, 0, 18, 14, 0, 2, 24, 6, 10, 1, 8, 25, 20, 13, 20, 13, 20, 5, 21, 21, 9, 19, 8, 9, 9, 5, 17, 18, 18, 20, 5, 17, 18, 3, 7, 21, 6, 0, 8, 3, 3, 1, 11, 0, 21, 6, 15, 11, 10, 13, 6, 7, 21, 7, 1, 1, 14, 15, 20, 2, 8, 21, 25, 19, 12, 18, 16, 0, 4, 10, 19, 14, 23, 6, 17, 2, 15, 19, 4, 13, 8, 14, 4, 15, 21, 4, 23, 20, 3, 18, 0, 12, 14, 14, 19, 0, 21, 18, 21, 17, 13, 9, 20, 17, 25, 17, 21, 16, 22, 4, 1, 13, 20, 15, 9, 7, 18, 18, 7, 22, 8, 18, 1, 13, 0, 24, 8, 12, 16, 1, 3, 6, 23, 16, 24, 5, 0, 1, 25, 3, 16, 9, 4, 24, 1, 11, 24, 9, 16, 11, 0, 2, 20, 16, 0, 1, 6, 19, 22, 12, 3, 23, 21, 4, 20, 1, 0, 18, 24, 10, 0, 12, 21, 17, 23, 0, 13, 1, 25, 9, 19, 0, 13, 21, 23, 6, 24, 25, 16, 9, 8, 16, 2, 22, 23, 3, 7, 16, 25, 11, 18, 19, 4, 11, 1, 25, 22, 9, 11, 14, 9, 3, 16, 8, 5, 11, 12, 15, 15, 19, 15, 15, 7, 17, 24, 18, 9, 8, 20, 23, 18, 17, 7, 8, 19, 23, 9, 13, 4, 17, 23, 21, 19, 11, 22, 22, 9, 3, 19, 23, 11, 2, 23, 8, 8, 21, 15, 1, 25, 7, 6, 14, 6, 7, 11, 3, 2, 11, 14, 10, 24, 3, 8, 10, 1, 18, 4, 6, 16, 12, 18, 12, 6, 5, 25, 24, 25, 7, 12, 17, 19, 15, 8, 23, 7, 6, 11, 6, 16, 14, 15, 13, 18, 5, 9, 21, 24, 8, 17, 25, 21, 22, 19, 24, 9, 9, 25, 21, 6, 25, 24, 3, 15, 20, 19, 13, 7, 13, 3, 0, 11, 2, 3, 23, 4, 14, 13, 7, 14, 3, 2, 18, 6, 1, 24, 19, 11, 6, 22, 9, 20, 3, 15, 23, 14, 18, 11, 11, 0, 2, 14, 21, 1, 12, 8, 8, 22, 10, 25, 20, 15, 22, 15, 21, 4, 19, 23, 5, 20, 4, 10, 17, 9, 7, 8, 11, 7, 10, 2, 18, 5, 24, 4, 16, 22, 13, 0, 11, 6, 19, 8, 21, 23, 24, 14, 19, 6, 3, 1, 17, 25, 22, 9, 14, 12, 15, 2, 24, 23, 17, 3, 3, 3, 6, 11, 20, 11, 0, 12, 17, 0, 3, 12, 24, 5, 13, 11, 19, 5, 2, 5, 12, 20, 19, 23, 2, 14, 23, 19, 4, 6, 15, 12, 2, 24, 17, 18, 9, 18, 4, 12, 20, 17, 19, 21, 16, 15, 13, 0, 17, 10, 23, 22, 10, 8, 20, 6, 4, 13, 11, 0, 3, 1, 5, 19, 17, 23, 17, 10, 10, 7, 4, 1, 20, 21, 23, 21, 21, 25, 2, 1, 8, 22, 4, 10, 16, 9, 15, 12, 12, 7, 3, 10, 14, 11, 9, 0, 7, 1, 1, 18, 23, 16, 6, 4, 20, 17, 18, 20, 17, 22, 8, 19, 6, 8, 14, 23, 14, 14, 15, 3, 24, 19, 16, 18, 14, 3, 6, 10, 8, 22, 12, 6, 8, 5, 3, 20, 10, 15, 19, 17, 8, 10, 7, 22, 0, 5, 19, 18, 16, 22, 24, 6, 18, 19, 19, 21, 1, 22, 14, 0, 24, 1, 20, 21, 7, 2, 11, 13, 10, 9, 7, 13, 15, 22, 2, 17, 4, 1, 4, 22, 22, 7, 18, 3, 12, 12, 7, 6, 20, 15, 25, 8, 13, 7, 5, 1, 25, 12, 1, 25, 16, 3, 23, 25, 9, 22, 4, 11, 16, 21, 20, 15, 17, 16, 13, 14, 20, 5, 23, 9, 0, 6, 3, 21, 2, 7, 2, 22, 7, 5, 8, 17, 14, 17, 8, 18, 21, 22, 14, 8, 15, 2, 10, 24, 0, 10, 23, 11, 16, 22, 5, 5, 19, 20, 14, 2, 19, 3, 25, 5, 10, 14, 22, 3, 5, 10, 20, 22, 16, 17, 22, 15, 23, 10, 0, 21, 17, 20, 3, 15, 0, 13, 17, 2, 10, 20, 8, 24, 5, 6, 19, 9, 4, 25, 11, 19, 10, 3, 24, 0, 10, 10, 9, 21, 16, 25, 6, 20, 11, 7, 17, 20, 10, 9, 22, 19, 21, 7, 0, 4, 11, 1, 9, 18, 18, 3, 1, 25, 5, 1, 20, 13, 2, 7, 19, 10, 13, 25, 3, 23, 13, 5, 10, 15, 11, 15, 22, 9, 10, 8, 18, 0}))
}

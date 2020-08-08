package leetcode

import (
	"container/list"
	"sort"
)

/*
99. 恢复二叉搜索树
二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。

示例 1:

输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2
示例 2:

输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
进阶:

使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？
*/

func recoverTree(root *TreeNode) {
	l := list.New()
	l.PushBack(root)
	nums := new([]int)
	nodes := getAllTreeNodeValues(root, nums)
	sort.Ints(*nums)
	for idx := range nodes {
		nodes[idx].Val = (*nums)[idx]
	}
}

func getAllTreeNodeValues(root *TreeNode, nums *[]int) []*TreeNode {
	var nodes []*TreeNode
	nodes = append(nodes, root)
	*nums = append(*nums, root.Val)
	if root.Left != nil {
		nodes = append(getAllTreeNodeValues(root.Left, nums), nodes...)
	}
	if root.Right != nil {
		nodes = append(nodes, getAllTreeNodeValues(root.Right, nums)...)
	}
	return nodes
}
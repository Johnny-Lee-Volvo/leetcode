package leetcode

/*
1299. 将每个元素替换为右侧最大元素
给你一个数组 arr ，请你将每个元素用它右边最大的元素替换，如果是最后一个元素，用 -1 替换。

完成所有替换操作后，请你返回这个数组。



示例：

输入：arr = [17,18,5,4,6,1]
输出：[18,6,6,6,1,-1]


提示：

1 <= arr.length <= 10^4
1 <= arr[i] <= 10^5
*/
func replaceElements(arr []int) []int {
	size := len(arr) - 1
	ret := make([]int, size+1)
	ret[size] = -1
	for size--; size >= 0; size-- {
		if arr[size+1] > ret[size+1] {
			ret[size] = arr[size+1]
			continue
		}
		ret[size] = ret[size+1]
	}
	return ret
}

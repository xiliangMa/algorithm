package _1_sort

import "testing"

/**
 *  堆排序 heapsort
 *  parent = (i+1)/2
 *  左子节点： c1 = 2i + 1
 *  右子节点： c2 = 2i + 2
 *	n 为数组长度
 *  i 为任意指定的节点
 */
func Heap_Sort(tree []int, n, i int) []int {
	if (i >= n) {
		return tree
	}
	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i

	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}

	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}

	if max != i {
		tree[max], tree[i] = tree[i], tree[max]
		Heap_Sort(tree, n, max)
	}

	return tree
}

/**
 * 最终输出结果 为： 堆排序（递归）---后 [10 5 3 4 1 2]
 */
func Test_Heap_Sort(t *testing.T){
	tree := []int{4, 10, 3, 5, 1, 2}
	t.Log("堆排序（递归）---后", Heap_Sort(tree, 6, 0))
}

package _1_sort

import "testing"

/**
 *
 *  parent = (i+1)/2
 *  左子节点： c1 = 2i + 1
 *  右子节点： c2 = 2i + 2
 *	n 为数组长度
 *  i 为任意指定的节点
 */
func Heapify(tree []int, n, i int) []int {
	if i >= n {
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
		Heapify(tree, n, max)
	}

	return tree
}

func Build_Heap(tree []int, n int) []int {
	last_node := n - 1
	parent := (last_node - 1) / 2

	for i := parent; i >= 0; i-- {
		Heapify(tree, n, i)
	}

	return tree
}

func Heap_Sort(tree []int, n int) []int {
	heap := Build_Heap(tree, n)
	for i := n - 1; i >= 0; i-- {
		// 交换根节点和最后一个节点
		heap[i], heap[0] = heap[0], heap[i]
		Heapify(heap, i, 0)
	}
	return tree
}

/**
 * 最终输出结果 为： 堆排序（递归）---后 [10 5 3 4 1 2]
 */
func Test_Heap_Sort(t *testing.T) {
	tree := []int{2, 5, 3, 1, 10, 4}
	t.Log("堆排序（递归）---后", Heap_Sort(tree, 6))
}

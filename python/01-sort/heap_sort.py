#coding=utf-8


'''
    堆排序 heapsort
    parent = (i+1)/2
    左子节点： c1 = 2i + 1
    右子节点： c2 = 2i + 2
    n 为数组长度
    i 为任意指定的节点
'''
def Heap_Sort(tree, n, i):
    if i >= n:
        return tree

    c1 = 2 * i + 1
    c2 = 2 * i + 2
    max = i

    if c1 < n and tree[c1] > tree[max]:
        max = c1

    if c2 < n and tree[c2] > tree[max]:
        max = c2

    if max != i:
        tree[i], tree[max] = tree[max], tree[i]
        Heap_Sort(tree, n, max)

    return tree

if __name__ == '__main__':
    tree = [4, 10, 3, 5, 1, 2]
    print("堆排序(递归)---后：", Heap_Sort(tree, 6, 0))
#coding=utf-8

'''
    快速排序(1) quick sort
'''
def Test_Qsort(arrs):
    if len(arrs) >= 2:
        pivod = arrs[len(arrs) - 1]
        arrs.remove(pivod)
        left, right = [], []
        for arr in arrs:
            if arr > pivod:
                right.append(arr)
            else:
                left.append(arr)
        return Test_Qsort(left) + [pivod] + Test_Qsort(right)
    else:
        return arrs

'''
    快速排序(1) quick sort
    避免额外的内存开销 去除前面left、right 数组
'''
def Test_Qsort2(arrs, start, end):
    if start >= end:
        return
    pivod = arrs[end]
    left = start
    right = end - 1
    while left <= right:
        while left <= right and arrs[left] < pivod:
            left += 1
        while left <= right and arrs[right] > pivod:
            right -= 1
        if left >= right:
            break
        arrs[left], arrs[right] = arrs[right], arrs[left]
    arrs[left], arrs[end] = arrs[end], arrs[left]
    Test_Qsort2(arrs, start, left - 1)
    Test_Qsort2(arrs, left + 1, end)

    return arrs



if __name__ == '__main__':
    arrs = [3, 1, 6, 4, 5, 0, 8, 9, 7, 2]
    print ("快速排序(1)---后：", Test_Qsort(arrs))

    arrs = [3, 1, 6, 4, 5, 0, 8, 9, 7, 2]
    print ("快速排序(2)---后：", Test_Qsort2(arrs, 0, 9))


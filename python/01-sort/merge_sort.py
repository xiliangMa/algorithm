#coding=utf-8

'''
     归并排序 bubble sort 递归
'''

def Merge(left, right):
    i, j = 0, 0
    result = []
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result

def Merge_sort(arrs):
    if len(arrs) <= 1:
        return arrs

    minddle = len(arrs) // 2
    left = Merge_sort(arrs[:minddle])
    right = Merge_sort(arrs[minddle:])
    return Merge(left, right)


if __name__ == '__main__':
    arrs = [3, 1, 6, 4, 5, 0, 8, 9, 7, 2]
    print("归并排序(递归)---后：", Merge_sort(arrs))

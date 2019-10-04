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

if __name__ == '__main__':
    arrs = [3, 1, 6, 4, 5, 0, 8, 9, 7, 2]
    print "快速排序---后：", Test_Qsort(arrs)

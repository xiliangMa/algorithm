#coding=utf-8



'''
    冒泡排序(1) bubble sort
'''
def Bsort(arrs):
    for i in range(len(arrs)-1):
        for j in range(len(arrs)-1-i):
            if arrs[j] > arrs[j+1]:
                arrs[j], arrs[j+1] = arrs[j+1], arrs[j]
                Bsort(arrs)
    return arrs


'''
    冒泡排序(2) bubble sort 递归
'''
def Bsort1(arrs):
    for i in range(len(arrs)-1):
        if arrs[i] > arrs[i+1]:
            arrs[i+1], arrs[i] = arrs[i], arrs[i+1]
            i += 1
            Bsort(arrs)
    return arrs


if __name__ == '__main__':
    arrs = [3, 1, 6, 4, 5, 0, 8, 9, 7, 2]
    print("冒泡排序(1)---后：", Bsort(arrs))
    print("冒泡排序(2递归)---后：", Bsort1(arrs))
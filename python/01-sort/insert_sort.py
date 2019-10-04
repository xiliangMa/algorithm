#coding=utf-8


'''
     插入排序 bubble sort
'''
def Isort(arrs):
    for i in  range(len(arrs)):
        for j in range(i):
            if arrs[i] < arrs[j]:
                arrs[i], arrs[j] = arrs[j], arrs[i]
    return arrs


if __name__ == '__main__':
    arrs = [3, 1, 6, 4, 5, 0, 8, 9, 7, 2]
    print("插入排序---后：", Isort(arrs))

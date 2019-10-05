#coding=utf-8


'''
    选择排序（1） bubble sort 递归
'''
def Ssort(arrs):
    for i in range(len(arrs) - 1):
        min = arrs[i]
        if arrs[i+1] < min:
            min = arrs[i+1]
            arrs[i], arrs[i+1] = arrs[i+1], arrs[i]
            Ssort(arrs)
    return arrs

if __name__ == '__main__':
    arrs = [3, 1, 6, 4, 5, 0, 8, 9, 7, 2]
    print "选择排序(递归)---后：",Ssort(arrs)





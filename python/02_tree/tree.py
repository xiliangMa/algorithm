#coding=utf-8


'''
    树结构
'''
class Node():
    data = None
    left = None
    right = None

    def __init__(self, data, left, right):
        self.data = data
        self.left = left
        self.right = right

'''
    构造树
'''
def build_node():
    node4 =  Node(8, None, None)
    node3 =  Node(7, None, None)
    node2 =  Node(6, node4, None)
    node1 =  Node(5, node2, node3)
    return node1

'''
    前序遍历 根 > 左 > 右
'''
def pre_order(node):
    if node != None:
        print node.data
        pre_order(node.left)
        pre_order(node.right)

'''
    前序遍历 左 > 根 > 右
'''
def in_order(node):
    if node != None:
        in_order(node.left)
        print node.data
        in_order(node.right)

'''
    前序遍历 左 > 右 > 根
'''
def post_order(node):
    if node != None:
        post_order(node.left)
        post_order(node.right)
        print node.data

if __name__ == '__main__':
    root = build_node()
    print "前序遍历结果:", pre_order(root)
    print "中序遍历结果：", in_order(root)
    print "后序遍历结果：", post_order(root)





class Node:
    """链表节点"""
    def __init__(self, data, next_node=Node):
        self.__data = data
        self.__next_node = next_node
    
    @property
    def data(self):
        return self.__data

    @data.setter
    def data(self, data):
        self.data = __data

    @property
    def next_node(self):
        return self.__next_node

    @next_node.setter
    def next_node(self, next_node):
        self.__next_node = next_node
    

class LinkedList:
    """链表"""
    def __init__(self, length):
        self.__head = None
        self.__length = length

    # 在某个节点之后插入
    def insert_after_node(self, node: Node, data):
        pass

    # 在某个节点之前插入
    def insert_before_node(self, node: Node, data):
        pass

    # 在头部插入节点
    def insert_head(self, data):
       pass 

    # 在尾部插入节点
    def insert_tail(self, data):
        pass

    # 通过索引返回节点
    def find_by_index(self, index:int)->Node:
        pass

    # 删除指定节点
    def delete_node(self, node:Node)->bool:
        pass

    # 打印链表
    def print(self):
        pass
    
    # 链表翻转
    def reverse(self):
        pass

    # 检查链表中是否有环
    def has_cycle(self)->bool:
        pass

    # 删除倒数第N个节点
    def delete_bottom_node(self, n: int)->bool:
        pass

    # 获取中间节点
    def find_middle_node(self)->LinkedList:
        pass

# 合并两个有序链表
def merge_sorted_linkedlist(l1: LinkedList, l2: LinkedList)->LinkedList:
    pass



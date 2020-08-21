package linearlist

import (
	"fmt"
)

/*
判断一个链表是不是会文链表
*/

// 自己先实现一种最笨,最暴力的方法, 用两个列表分别保存链表的前后两个部分
func isPalindrome(l *LinearList) bool {
	if l.length == 0 || l.length == 1 {
		return false
	}

	cur := l.head.next
	m := l.length / 2
	// 链表的前半部分
	firstHalf := make([]interface{}, 0)
	// 链表的后半部分
	secondHalf := make([]interface{}, 0)
	for i := 0; i < l.length; i++ {
		if i < m {
			firstHalf = append(firstHalf, cur.GetValue())
		} else {
			secondHalf = append(secondHalf, cur.GetValue())
		}
		cur = cur.next
	}
	fmt.Println(firstHalf, secondHalf)
	if len(firstHalf) != len(secondHalf) {
		secondHalf = secondHalf[1:]
	}
	for i := 0; i < m; i++ {
		if firstHalf[i] != secondHalf[m-i-1] {
			return false
		}
	}
	return true
}

// 方法1:开一个栈存放链表示链表的前半段
func isPalindrome1(l *LinearList) bool {
	if l.length == 0 || l.length == 1 {
		return false
	}

	s := make([]interface{}, 0, l.length/2)
	cur := l.head
	for i := 1; i <= l.length; i++ {
		cur = cur.next
		// 如果链表的节点个数是奇数个，则直接忽略中间节点
		if l.length%2 != 0 && i == (l.length/2+1) {
			continue
		}
		if i <= l.length/2 { // 前一半保存到slice中
			s = append(s, cur.GetValue())
		} else {
			if s[l.length-i] != cur.GetValue() {
				return false
			}
		}
	}
	return true
}

// 方法2:
func isPalindrome2(l *LinearList) {

}

package main

/*21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
/**
 * Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路我们可以如下递归地定义两个链表里的 merge 操作（忽略边界情况，比如空链表等）：
 list1[0]+merge(list1[1:],l2)    list1[0]<list2[0]
 list2[0]+merge(list2,list1[1:]) otherwise
也就是说，两个链表头部值较小的一个节点与剩下元素的 merge 操作结果合并
 每次都会更新 list1[0]和list2[0]
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1 //更新list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

}

package day16

 // Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

//逆序存储，需要一个链表来接收值
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}//取地址符产生dummy指针

	curr, carry := dummy, 0//现在的指针和现在的进位 curr , carry
	for l1 != nil || l2 != nil {//循环条件  l1 l2  其中有一个不为空，就可以继续遍历
		x ,y :=0,0
		//检查非空，空节点以0值参与遍历
		if l1 != nil {
			x = l1.Val//x是l1节点值
		}
		if l2 != nil {
			y = l2.Val//y是l2节点值
		}

		total := x + y + carry   //当前进行计算和  x,y,carry
		curr.Next = &ListNode{Val: total % 10}//建立当前指针的下一个指针 ，取地址 赋值 total%10

		curr = curr.Next//当前计算完成节点赋值，指针移到下一个指针
		carry = total / 10//进位是为下一次计算准备的 ；这次进位是 上一次计算和/10  看代码 total计算在前，carry计算在后

		if l1 != nil {
			l1 = l1.Next//判断 l1不为空l1指向下一个节点
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
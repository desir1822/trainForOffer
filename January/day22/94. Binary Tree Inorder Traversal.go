package day22

/* Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

//中序遍历 中 左 右
//          左 变成 中  遍历 左的 左 ，左的右
//             右变成中 遍历 右的 左   右的右
//我们可以发现 递归  会出现一个 不变量 更新 ，这个不变量 就是 中
//循环条件 是遍历的这个中 为 空
//递归的操作 这里 只进行了 入栈操作
func inorderTraversal2(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//root保留栈末尾元素，之后 末尾元素出栈
		res = append(res, root.Val) //存入root的值
		root = root.Right           //更新root节点值
	}
	return
}

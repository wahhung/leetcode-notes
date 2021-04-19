 // 445. 两数相加 II
 // 利用栈后进先出原理（不修改结点）
 // 1. 先将两个链表的结点的值放进栈
 // 2. 定义 remain 表示上一个相加的进位，cur 表示当前结点
 // 3. 定义v1,v2,分别判断两个栈是否为空，不为空出栈赋值给 它们
 // 4. val = v1 + v2 + remain，结点的值=val%10， remian=val/10
 // 5. node.Next->cur cur = node
 // 6. 两个栈都为空，循环结束
 // 7.如果remain > 0，创建新的结点进行保存值，且node.Next->cur cur = node
 // Time:O(max(n, m)) Space:O(n+m)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    arr1 := []int{}
    arr2 := []int{}
    for l1 != nil {
        arr1 = append(arr1, l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        arr2 = append(arr2, l2.Val)
        l2 = l2.Next
    }
    var remain int
    var cur *ListNode
    for len(arr1) > 0 || len(arr2) > 0 {
        var v1, v2 int
        if len(arr1) > 0 {
            v1 = arr1[len(arr1)-1]
            arr1 = arr1[:len(arr1)-1]
        }
        if len(arr2) > 0 {
            v2 = arr2[len(arr2)-1]
            arr2 = arr2[:len(arr2)-1]
        }
        val := (v1 + v2 + remain) % 10
        remain = (v1 + v2 + remain) / 10
        node := &ListNode{Val:val, Next:cur}
        cur = node
    }
    if remain > 0 {
        node := &ListNode{Val:remain, Next:cur}
        cur = node
    }
    return cur
}

// 反转两个链表，通过指针方式遍历相加

package main

type ListNode struct {
	Val int
	Next *ListNode
}
// 2-4-3
// 5-6-0
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

   head := new(ListNode)
   cur := head
   carry := 0

   for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

	    if l1 != nil {
		    sum += l1.Val
		    l1 = l1.Next
	    }

	   if l2 != nil {
		   sum += l2.Val
		   l2 = l2.Next
	   }

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next

   }
   return head.Next
}

func main() {

	l1_1 := new(ListNode)
	l1_1.Val = 2

	l1_2 := new(ListNode)
    l1_2.Val = 4

	l1_3 := new(ListNode)
	l1_3.Val = 3

	l1_1.Next  = l1_2
	l1_2.Next = l1_3


	l2_1 := new(ListNode)
	l2_1.Val = 5

	l2_2 := new(ListNode)
	l2_2.Val = 6

	l2_3 := new(ListNode)
	l2_3.Val = 4

	l2_1.Next  = l2_2
	l2_2.Next = l2_3

	addTwoNumbers(l1_1,l2_1)


}
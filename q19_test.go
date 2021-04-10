package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	s, in, inpre := head, head, head // 哨兵先探路，步兵随其后，步兵上把手

	// 哨兵先走 n 步
	for i := 0; i < n; i++ {
		s = s.Next
	}

	// 哨兵和步兵一起前进，直到哨兵到头，步兵上把手要一直记录步兵位置
	for s != nil {
		s = s.Next
		inpre = in
		in = in.Next
	}

	// 步兵还在头部位置没有前进，去掉头结点
	if in == head {
		return head.Next
	}

	// 删除此步兵位置，将步兵上把手的 next 地址指向步兵的 next 地址
	inpre.Next = in.Next
	return head
}

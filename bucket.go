package GoSortAlgorithm

type ListNode struct {
	val  int
	next *ListNode
}

func BucketSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	bucketsort(arr)
}

func bucketsort(arr []int) {
	max, min := GetMaxAndMin(arr)
	bucketCount := (max-min)/len(arr) + 1
	buckets := make([]*ListNode, bucketCount)
	for i := 0; i < bucketCount; i++ {
		// fack head node
		buckets[i] = &ListNode{val: INT_MIN, next: nil}
	}
	for _, val := range arr {
		idx := (val - min) / len(arr)
		node := &ListNode{val: val, next: nil}
		InsertList(buckets[idx], node)
	}
	idx := 0
	for i := 0; i < bucketCount; i++ {
		for cur := buckets[i].next; cur != nil; cur = cur.next {
			arr[idx] = cur.val
			idx++
		}
	}
}

func InsertList(head, node *ListNode) {
	cur := head
	for cur.next != nil && node.val >= cur.next.val {
		cur = cur.next
	}
	node.next = cur.next
	cur.next = node
}

package cp

import (
	"testing"
)

type opts struct {
	list []int
	cIdx []int // which index have child
}

func makeMultilevelDoublyLinkedList(opts ...opts) (head *ListNode) {
	tmp := make([]*ListNode, 0)
	for _, o := range opts {
		mIdx := make(map[int]byte)
		for _, idx := range o.cIdx {
			mIdx[idx] = 1
		}

		var p *ListNode
		for idx, v := range o.list {
			n := &ListNode{Val: v}

			if head == nil {
				head = n
			}
			if p == nil {
				p = n
				if len(tmp) != 0 {
					tmp[0].Child = n
					tmp = tmp[1:]
				}
			} else {
				n.Prev = p
				p.Next = n
				p = p.Next
			}

			if mIdx[idx] == 1 {
				tmp = append(tmp, n)
			}
		}
	}
	return
}

func goThroughAllNodes(head *ListNode, out *[]int) {
	if out == nil {
		tmp := make([]int, 0)
		out = &tmp
	}
	for p := head; p != nil; p = p.Next {
		*out = append(*out, p.Val)
		if p.Child != nil {
			goThroughAllNodes(p.Child, out)
		}
	}
}

func Test_makeMultilevelDoublyLinkedList(t *testing.T) {
	tasks := []struct {
		input  []opts
		expect []int
	}{
		{
			input: []opts{
				{list: []int{1, 2, 3, 4, 5, 6}, cIdx: []int{2}},
				{list: []int{7, 8, 9, 10}, cIdx: []int{1}},
				{list: []int{11, 12}, cIdx: []int{}},
			},
			expect: []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6},
		},
	}

	for i, task := range tasks {
		var got []int
		head := makeMultilevelDoublyLinkedList(task.input...)
		goThroughAllNodes(head, &got)
		if !equal(got, task.expect) {
			t.Errorf("task #%d failed, output: %v, expect: %v", i, got, task.expect)
		}
	}
}

func Test_flatten(t *testing.T) {
	tasks := []struct {
		input  []opts
		expect []int
	}{
		{
			input: []opts{
				{list: []int{1, 2, 3}, cIdx: []int{}},
			},
			expect: []int{1, 2, 3},
		},
		{
			input: []opts{
				{list: []int{1, 2, 3}, cIdx: []int{1}},
				{list: []int{4, 5}, cIdx: []int{}},
			},
			expect: []int{1, 2, 4, 5, 3},
		},
		{
			input: []opts{
				{list: []int{1, 2, 3, 4, 5, 6}, cIdx: []int{2}},
				{list: []int{7, 8, 9, 10}, cIdx: []int{1}},
				{list: []int{11, 12}, cIdx: []int{}},
			},
			expect: []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6},
		},
	}

	for fIdx, f := range []func(*ListNode) *ListNode{flatten} {
		for i, task := range tasks {
			var got []int
			origin := makeMultilevelDoublyLinkedList(task.input...)
			ans := f(origin)
			for prev, p := ans, ans; p != nil; prev, p = p, p.Next {
				// check pointers
				if p == ans {
					if p.Prev != nil {
						t.Errorf("func #%d, head prev not nil", fIdx)
					}
				} else {
					if p.Prev != prev {
						t.Errorf("func #%d, doubly linked list node prev wrong, got: %v, expect: %v", fIdx, p.Prev, prev)
					}
				}
				if p.Child != nil {
					t.Errorf("func #%d, flatten doubly linked list child not nil", fIdx)
				}
				if ans != origin {
					t.Errorf("func #%d, origin head has been changed", fIdx)
				}
				got = append(got, p.Val)
			}
			if !equal(got, task.expect) {
				t.Errorf("func #%d, task #%d failed, output: %v, expect: %v", fIdx, i, got, task.expect)
			}
		}
	}
}
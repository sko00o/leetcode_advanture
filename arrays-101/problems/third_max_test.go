package problems

import (
	"fmt"
	"testing"
)

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},
		{
			name: "test 1",
			args: args{
				nums: []int{1, 2, -3},
			},
			want: -3,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{5, 2, 4, 1, 3, 6, 0},
			},
			want: 4,
		},
		{
			name: "test 3",
			args: args {
				nums: []int{1,2,2,5,3,5},
			},
			want: 2,
		},
	}

	for idx, f := range []func([]int) int{
		thirdMax,
		thirdMax1,
		thirdMax2,
		thirdMax3,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					var nums = make([]int, len(tt.args.nums))
					copy(nums, tt.args.nums)
					if got := f(nums); got != tt.want {
						t.Errorf("thirdMax() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}

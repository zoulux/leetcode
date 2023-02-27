package leetcode

import (
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				values: []int{2, 3, 4},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTreeNode(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				values: []int{1, 2, 3, 4, 5, 0, 6},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTreeNode(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	list "github.com/slonegd-otus-go/06_list"
)

func TestList_PushBack(t *testing.T) {
	var list list.List

	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"1", 1, "[1]"},
		{"2.5", 2.5, "[1 2.5]"},
		{"text", "text", "[1 2.5 text]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.PushBack(tt.value)
			got := list.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_PushFront(t *testing.T) {
	var list list.List

	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"1", 1, "[1]"},
		{"2.5", 2.5, "[2.5 1]"},
		{"text", "text", "[text 2.5 1]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.PushFront(tt.value)
			got := list.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_Len(t *testing.T) {
	tests := []struct {
		name   string
		action func(*list.List)
		want   int
	}{
		{
			name:   "empty",
			action: func(*list.List) {},
			want:   0,
		},
		{
			name: "add 1",
			action: func(list *list.List) {
				list.PushBack(100)
			},
			want: 1,
		},
		{
			name: "push back and push front",
			action: func(list *list.List) {
				list.PushBack(100)
				list.PushBack(200)
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list list.List
			tt.action(&list)
			got := list.Len()
			assert.Equal(t, tt.want, got)
		})
	}
}

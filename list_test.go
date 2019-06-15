package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	list "github.com/slonegd-otus-go/06_list"
)

func TestList_PushBack(t *testing.T) {
	var list list.Type

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
	var list list.Type

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
		action func(*list.Type)
		want   int
	}{
		{
			name:   "empty",
			action: func(*list.Type) {},
			want:   0,
		},
		{
			name: "add 1",
			action: func(list *list.Type) {
				list.PushBack(100)
			},
			want: 1,
		},
		{
			name: "push back and push front",
			action: func(list *list.Type) {
				list.PushBack(100)
				list.PushBack(200)
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list list.Type
			tt.action(&list)
			got := list.Len()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_First(t *testing.T) {
	tests := []struct {
		name      string
		action    func(*list.Type)
		wantValue interface{}
	}{
		{
			name:      "empty",
			action:    func(*list.Type) {},
			wantValue: nil,
		},
		{
			name: "add 1",
			action: func(list *list.Type) {
				list.PushBack(100)
			},
			wantValue: 100,
		},
		{
			name: "push back 2 values",
			action: func(list *list.Type) {
				list.PushBack(100)
				list.PushBack(200)
			},
			wantValue: 100,
		},
		{
			name: "push front 2 values",
			action: func(list *list.Type) {
				list.PushFront(100)
				list.PushFront(200)
			},
			wantValue: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list list.Type
			tt.action(&list)
			got := list.First()
			if tt.wantValue == nil {
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, tt.wantValue, got.Value)
		})
	}
}

func TestList_Last(t *testing.T) {
	tests := []struct {
		name      string
		action    func(*list.Type)
		wantValue interface{}
	}{
		{
			name:      "empty",
			action:    func(*list.Type) {},
			wantValue: nil,
		},
		{
			name: "add 1",
			action: func(list *list.Type) {
				list.PushBack(100)
			},
			wantValue: 100,
		},
		{
			name: "push back 2 values",
			action: func(list *list.Type) {
				list.PushBack(100)
				list.PushBack(200)
			},
			wantValue: 200,
		},
		{
			name: "push front 2 values",
			action: func(list *list.Type) {
				list.PushFront(100)
				list.PushFront(200)
			},
			wantValue: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list list.Type
			tt.action(&list)
			got := list.Last()
			if tt.wantValue == nil {
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, tt.wantValue, got.Value)
		})
	}
}

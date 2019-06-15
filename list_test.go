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
		{
			name: "push back and push front and remove",
			action: func(list *list.Type) {
				list.PushBack(100)
				list.PushBack(200)
				list.First().Remove()
			},
			want: 1,
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

func TestItem_Next(t *testing.T) {
	tests := []struct {
		name      string
		action    func(*list.Type)
		wantValue interface{}
	}{
		{
			name: "len 1: list.First().Next()",
			action: func(list *list.Type) {
				list.PushBack(100)
			},
			wantValue: nil,
		},
		{
			name: "len 2: list.First().Next()",
			action: func(list *list.Type) {
				list.PushBack(100)
				list.PushBack(200)
			},
			wantValue: 200,
		},
		{
			name: "len >2: list.First().Next()",
			action: func(list *list.Type) {
				list.PushBack(300)
				list.PushBack(400)
				list.PushBack(500)
				list.PushBack(600)
			},
			wantValue: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list list.Type
			tt.action(&list)
			got := list.First().Next()
			if tt.wantValue == nil {
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, tt.wantValue, got.Value)
		})
	}
}

func TestItem_Prev(t *testing.T) {
	tests := []struct {
		name      string
		action    func(*list.Type)
		wantValue interface{}
	}{
		{
			name: "len 1: list.Last().Prev()",
			action: func(list *list.Type) {
				list.PushBack(100)
			},
			wantValue: nil,
		},
		{
			name: "len 2: list.Last().Prev()",
			action: func(list *list.Type) {
				list.PushBack(100)
				list.PushBack(200)
			},
			wantValue: 100,
		},
		{
			name: "len >2: list.Last().Prev()",
			action: func(list *list.Type) {
				list.PushBack(600)
				list.PushBack(700)
				list.PushBack(800)
				list.PushBack(900)
			},
			wantValue: 800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list list.Type
			tt.action(&list)
			got := list.Last().Prev()
			if tt.wantValue == nil {
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, tt.wantValue, got.Value)
		})
	}
}

func TestItem_Remove(t *testing.T) {
	tests := []struct {
		name         string
		makeList     func() *list.Type
		removeAction func(*list.Type)
		wantBefore   string
		wantAfter    string
	}{
		{
			name: "remove first",
			makeList: func() *list.Type {
				list := &list.Type{}
				list.PushBack(1)
				list.PushBack(2)
				list.PushBack(3)
				list.PushBack(4)
				return list
			},
			removeAction: func(list *list.Type) {
				list.First().Remove()
			},
			wantBefore: "[1 2 3 4]",
			wantAfter:  "[2 3 4]",
		},
		{
			name: "remove last",
			makeList: func() *list.Type {
				list := &list.Type{}
				list.PushBack(1)
				list.PushBack(2)
				list.PushBack(3)
				list.PushBack(4)
				return list
			},
			removeAction: func(list *list.Type) {
				list.Last().Remove()
			},
			wantBefore: "[1 2 3 4]",
			wantAfter:  "[1 2 3]",
		},
		{
			name: "remove second",
			makeList: func() *list.Type {
				list := &list.Type{}
				list.PushBack(1)
				list.PushBack(2)
				list.PushBack(3)
				list.PushBack(4)
				return list
			},
			removeAction: func(list *list.Type) {
				list.First().Next().Remove()
			},
			wantBefore: "[1 2 3 4]",
			wantAfter:  "[1 3 4]",
		},
		{
			name: "remove prelast",
			makeList: func() *list.Type {
				list := &list.Type{}
				list.PushBack(1)
				list.PushBack(2)
				list.PushBack(3)
				list.PushBack(4)
				return list
			},
			removeAction: func(list *list.Type) {
				list.Last().Prev().Remove()
			},
			wantBefore: "[1 2 3 4]",
			wantAfter:  "[1 2 4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.makeList()
			assert.Equal(t, tt.wantBefore, list.String())
			tt.removeAction(list)
			assert.Equal(t, tt.wantAfter, list.String())
		})
	}
}

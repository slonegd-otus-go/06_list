package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_PushBack(t *testing.T) {
	var list List

	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"1", 1, "1 "},
		{"2.5", 2.5, "1 2.5 "},
		{"text", "text", "1 2.5 text "},
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
	var list List

	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"1", 1, "1 "},
		{"2.5", 2.5, "2.5 1 "},
		{"text", "text", "text 2.5 1 "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.PushFront(tt.value)
			got := list.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

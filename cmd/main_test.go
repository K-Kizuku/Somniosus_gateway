package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		input   []int
		wantErr bool
	}{
		{
			name:    "test001",
			want:    3,
			input:   []int{1, 2},
			wantErr: false,
		},
		{
			name:    "test002",
			want:    2,
			input:   []int{3, -1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			got := add(tt.input[0], tt.input[1])
			if tt.want != got {
				t.Errorf("main add() want = %v, got = %v", tt.want, got)
			}
		})
	}
}

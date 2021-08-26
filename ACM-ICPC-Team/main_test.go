package main

import (
	"reflect"
	"testing"
)

func Test_acmTeam(t *testing.T) {
	type args struct {
		topic []string
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Case 1",
			args: struct{ topic []string }{
				topic: []string{"10101", "11110", "00010"},
			},
			want: []int32{5, 1},
		}, {
			name: "Case 2",
			args: struct{ topic []string }{
				topic: []string{"10101", "11100", "11010", "00101"},
			},
			want: []int32{5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := acmTeam(tt.args.topic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("acmTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

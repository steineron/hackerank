package main

import "testing"

func Test_nextJump(t *testing.T) {
	type args struct {
		c        []int32
		location int
	}
	type test struct {
		name string
		args args
		want int32
	}

	tests := []test{
		{
			name: "empty input",
			args: args{
				c: []int32 {0,0},
				location: 0,
			},
			want: 1,
		},{
			name: "010 input",
			args: args{
				c: []int32 {0,1,0},
				location: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextJump(tt.args.c, tt.args.location); got != tt.want {
				t.Errorf("nextJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

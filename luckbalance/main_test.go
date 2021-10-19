package main

import "testing"

func Test_luckBalance(t *testing.T) {
	type args struct {
		k        int32
		contests [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test 0",
			args: args{
				k:        2,
				contests: [][]int32{{5, 1}, {1, 1}, {4, 0}},
			},
			want: 10,
		}, {
			name: "test 1",
			args: args{
				k:        1,
				contests: [][]int32{{5, 1}, {1, 1}, {4, 0}},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := luckBalance(tt.args.k, tt.args.contests); got != tt.want {
				t.Errorf("luckBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

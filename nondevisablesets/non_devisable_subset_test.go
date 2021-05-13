package main

import "testing"

func Test_nonDivisibleSubset(t *testing.T) {
	type args struct {
		k int32
		s []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "that empty set returns 0",
			args: args{k: 8, s: []int32{}},
			want: 0,
		},{
			name: "that 0 K returns 0",
			args: args{k: 0, s: []int32{1,2,3}},
			want: 0,
		},{
			name: "case 1",
			args: args{k: 3, s: []int32{1, 7, 2, 4}},
			want: 3,
		},{
			name: "case 2",
			args: args{k: 4, s: []int32{19,10,12,10,24,25,22}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nonDivisibleSubset(tt.args.k, tt.args.s); got != tt.want {
				t.Errorf("nonDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

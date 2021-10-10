package main

import "testing"

func Test_counterGame(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "6=Richard",
			args: args{6},
			want: "Richard",
		}, {
			name: "1560834904=Richard",
			args: args{1560834904},
			want: "Richard",
		},{
			name: "1768820483=Richard",
			args: args{1768820483},
			want: "Richard",
		},{
			name: "132=Louise",
			args: args{132},
			want: "Louise",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := counterGame(tt.args.n); got != tt.want {
				t.Errorf("counterGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

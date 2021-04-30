package main

import "testing"

func Test_minimumBribes(t *testing.T) {
	type testcase struct {
		q []int32
	}
	type test struct {
		name    string
		test    testcase
		want    int32
		wantErr bool
	}
	tests := []test{
		{
			name: "123456",
			test: testcase{
				q: []int32{1, 2, 3, 4, 5, 6},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "1234 6 5",
			test: testcase{
				q: []int32{1, 2, 3, 4, 6, 5},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "123 6 45",
			test: testcase{
				q: []int32{1, 2, 3, 6, 4, 5},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "12 6 345",
			test: testcase{
				q: []int32{1, 2, 6, 3, 4, 5},
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "123 6 5 4",
			test: testcase{
				q: []int32{1, 2, 3, 6, 5, 4},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "1 3 2 6 5 4",
			test: testcase{
				q: []int32{1, 3, 2, 6, 5, 4},
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "3 1 2 6 5 4",
			test: testcase{
				q: []int32{ 3,1, 2, 6, 5, 4},
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "3 2 1 6 5 4",
			test: testcase{
				q: []int32{ 3, 2, 1, 6, 5, 4},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := minimumBribes(tt.test.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("minimumBribes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("minimumBribes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

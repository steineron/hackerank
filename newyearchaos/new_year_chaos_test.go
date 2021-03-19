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
	tests := []test {
		{
			name:    "123456",
			test:    testcase{
				q: []int32{1,2,3,4,5,6},
			},
			want:    0,
			wantErr: false,
		},
		{
			name:    "123465",
			test:    testcase{
				q: []int32{1,2,3,4,6,5},
			},
			want:    1,
			wantErr: false,
		},
		{
			name:    "123645",
			test:    testcase{
				q: []int32{1,2,3,6,4,5},
			},
			want:    2,
			wantErr: false,
		},
		{
			name:    "126345",
			test:    testcase{
				q: []int32{1,2,3,4,5,6},
			},
			want:    -1,
			wantErr: true,
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

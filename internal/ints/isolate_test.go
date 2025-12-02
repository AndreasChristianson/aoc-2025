package ints

import "testing"

func TestIsolate(t *testing.T) {
	type args struct {
		target              int
		lowestDecimalDigit  int
		highestDecimalDigit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestIsolate",
			args: args{
				target:              1234567890,
				lowestDecimalDigit:  2,
				highestDecimalDigit: 6,
			},
			want: 5678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Isolate(tt.args.target, tt.args.lowestDecimalDigit, tt.args.highestDecimalDigit); got != tt.want {
				t.Errorf("Isolate() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
				target:              1234_5678_90,
				lowestDecimalDigit:  2,
				highestDecimalDigit: 6,
			},
			want: 5678,
		},
		{
			name: "go doc",
			args: args{
				target:              1098_76_54321,
				lowestDecimalDigit:  5,
				highestDecimalDigit: 7,
			},
			want: 76,
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

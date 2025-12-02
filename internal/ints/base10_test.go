package ints

import "testing"

func TestBase10Length(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "12345",
			args: args{input: 12345},
			want: 5,
		},
		{
			name: "10",
			args: args{input: 10},
			want: 2,
		},
		{
			name: "100",
			args: args{input: 100},
			want: 3,
		},
		{
			name: "0",
			args: args{input: 0},
			want: 1,
		},
		{
			name: "9",
			args: args{input: 9},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base10Length(tt.args.input); got != tt.want {
				t.Errorf("Base10Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

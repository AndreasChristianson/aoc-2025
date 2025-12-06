package ints

import (
	"reflect"
	"sort"
	"testing"
)

func TestFactors(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "twelve",
			args: args{
				input: 12,
			},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "zero",
			args: args{
				input: 0,
			},
			want: []int{},
		},
		{
			name: "thirteen",
			args: args{
				input: 13,
			},
			want: []int{1},
		},
		{
			name: "one",
			args: args{
				input: 1,
			},
			want: []int{},
		},
		{
			name: "four",
			args: args{
				input: 4,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Factors(tt.args.input)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factors() = %v, want %v", got, tt.want)
			}
		})
	}
}

package day_6

import (
	"aoc-2025/internal/disk_io"
	"testing"
)

func TestPart1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				lines: disk_io.ReadLines("example-input.txt"),
			},
			want: "4277556",
		},
		{
			name: "actual",
			args: args{
				lines: disk_io.ReadLines("input.txt"),
			},
			want: "4693159084994",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				lines: disk_io.ReadLines("example-input.txt"),
			},
			want: "3263827",
		},
		{
			name: "actual",
			args: args{
				lines: disk_io.ReadLines("input.txt"),
			},
			want: "11643736116335",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

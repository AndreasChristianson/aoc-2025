package day_9

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
			want: "50",
		},
		{
			name: "actual",
			args: args{
				lines: disk_io.ReadLines("input.txt"),
			},
			want: "4782268188",
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
			want: "24",
		},
		{
			name: "actual",
			args: args{
				lines: disk_io.ReadLines("input.txt"),
			},
			want: "1574717268",
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

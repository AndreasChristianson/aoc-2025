package day_8

import (
	"aoc-2025/internal/disk_io"
	"testing"
)

func TestPart1(t *testing.T) {
	type args struct {
		lines           []string
		connectionCount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				lines:           disk_io.ReadLines("example-input.txt"),
				connectionCount: 10,
			},
			want: "40",
		},
		{
			name: "actual",
			args: args{
				lines:           disk_io.ReadLines("input.txt"),
				connectionCount: 1000,
			},
			want: "52668",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines, tt.args.connectionCount); got != tt.want {
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
			want: "25272",
		},
		{
			name: "actual",
			args: args{
				lines: disk_io.ReadLines("input.txt"),
			},
			want: "1474050600",
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

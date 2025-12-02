package strings

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not palindrome",
			args: args{
				"testing",
			},
			want: false,
		},
		{
			name: "is palindrome odd",
			args: args{
				"civic",
			},
			want: true,
		},
		{
			name: "Is Palindrome even",
			args: args{
				"tattarrattat",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.input); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

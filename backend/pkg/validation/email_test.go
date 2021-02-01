package validation

import (
	"strings"
	"testing"
)

func TestIsEmailValid(t *testing.T) {
	type args struct {
		email string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validation.email.1",
			args: args{
				email: strings.Repeat("a", emailMinLen-2) + "@" + strings.Repeat("d", emailMinLen-2),
			},
			want: true,
		},
		{
			name: "validation.email.2",
			args: args{
				email: strings.Repeat("a", emailMaxLen-5) + "@a.co",
			},
			want: true,
		},
		{
			name: "validation.email.3",
			args: args{
				email: strings.Repeat("a", emailMaxLen-5+1) + "@a.co",
			},
			want: false,
		},
		{
			name: "validation.email.4",
			args: args{
				email: strings.Repeat("a", emailMinLen) + "@",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmailValid(tt.args.email); got != tt.want {
				t.Errorf("IsEmailValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

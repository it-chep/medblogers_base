package checker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CleanTelegramStringRegex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "https://t.me/readydoctor",
			want:  "readydoctor",
		},
		{
			input: "http://t.me/readydoctor",
			want:  "readydoctor",
		},
		{
			input: "t.me/readydoctor",
			want:  "readydoctor",
		},
		{
			input: "@readydoctor",
			want:  "readydoctor",
		},
		{
			input: "readydoctor",
			want:  "readydoctor",
		},
		{
			input: "t.me/readydoctor",
			want:  "readydoctor",
		},
		{
			input: "Брендмейкер врачей",
			want:  "брендмейкер врачей",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			res := cleanTelegramStringRegex(test.input)
			assert.Equal(t, test.want, res)
		})
	}
}

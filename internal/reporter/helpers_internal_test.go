package reporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processFormValue(t *testing.T) {
	type args struct {
		data string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "array with empty elements",
			args: args{
				data: "test\r\ntest2\r\n\r\n\r\ntest3",
			},
			want: []string{
				"test", "test2", "test3",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := processFormValue(tt.args.data)
			assert.Equal(t, tt.want, got)
		})
	}
}

package errpool

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {

	tests := []struct {
		err ErrList
		exp string
	}{
		{
			ErrList{errors.New("foo"), errors.New("bar"), errors.New("jazz")},
			"foo | bar | jazz",
		},
		{
			ErrList{errors.New("foo")},
			"foo",
		},
		{
			ErrList{},
			"",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.exp, func(t *testing.T) {

			assert.Equal(t, tt.exp, tt.err.Error())
		})
	}
}

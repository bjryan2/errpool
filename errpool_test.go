package errpool

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorAdd(t *testing.T) {

	tests := []struct {
		errs ErrList
		exp  string
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
			p := NewPool()

			for _, err := range tt.errs {
				err := err
				p.Add(1)
				go func() {
					p.Error(err)
					p.Done()
				}()
			}

			resErr := p.Wait()
			if tt.exp == "" {
				assert.Nil(t, resErr)
			} else {
				if assert.NotNil(t, resErr) {

					errParts := strings.Split(resErr.Error(), " | ")
					for _, err := range tt.errs {
						assert.Contains(t, errParts, err.Error())
					}
				}
			}
		})
	}
}

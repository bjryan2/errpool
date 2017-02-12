package errpool

import (
	"bytes"
)

// ErrList
type ErrList []error

// Error concatenates all of the errors contained in the ErrList into
// a single error.
func (e ErrList) Error() string {

	numErrs := len(e)

	if numErrs == 0 {
		return ""
	}

	var b bytes.Buffer
	b.WriteString("")

	for i, err := range e {
		b.WriteString(err.Error())

		// seperate consecutive errs with a delimiter
		if i != numErrs-1 {
			b.WriteString(" | ")
		}
	}
	return b.String()
}

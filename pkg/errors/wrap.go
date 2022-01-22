package errors

import "fmt"

func Wrap(errp *error, format string, args ...interface{}) error {
	if *errp != nil {
		s := fmt.Sprintf(format, args...)
		return fmt.Errorf("%s: %w", s, *errp)
	}
	return fmt.Errorf("%s", fmt.Sprintf(format, args...))
}

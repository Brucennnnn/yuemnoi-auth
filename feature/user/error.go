package user

import (
	"errors"
	"fmt"
)

func IsErrorUserNotFound(err error) bool {
	return errors.Is(err, ErrorUserNotFound)
}

func WrapErrorUserNotFound(err error) error {
	if err == nil {
		return ErrorUserNotFound
	}
	return fmt.Errorf("%w: %v", ErrorUserNotFound, err)
}

var ErrorUserNotFound = errors.New("user not found")

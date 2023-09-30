package errors

import (
	"errors"
	"fmt"
	"strings"
)

type Error struct {
	Inner           error
	FriendlyMessage string
	Misc            map[string]interface{}
}

func (e Error) Error() string {
	var msg []string
	unwrapped := errors.Unwrap(e)

	for unwrapped != nil {
		msg = append(msg, unwrapped.Error())
		unwrapped = errors.Unwrap(unwrapped)
	}

	if len(msg) > 0 {
		return e.FriendlyMessage + ": " + strings.Join(msg, ": ")
	}

	return e.FriendlyMessage
}

func (e Error) Unwrap() error {
	return e.Inner
}

func Wrap(e error, misc map[string]interface{}, messageF string, messageArgs ...interface{}) Error {
	return Error{
		Inner:           e,
		FriendlyMessage: fmt.Sprintf(messageF, messageArgs...),
		Misc:            misc,
	}
}

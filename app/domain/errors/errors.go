package errors

import (
	"errors"
	"strings"
)

type Error struct {
	Inner           error
	FriendlyMessage string
	Misc            map[string]interface{}
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

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

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Inner
}

func (e *Error) Is(target error) bool {
	tErr, ok := target.(*Error)
	if !ok {
		return false
	}

	if e == nil {
		return target == nil
	}

	if e.Error() == tErr.Error() {
		return true
	}

	inner := e.Inner
	for {
		if inner == nil {
			return false
		}

		if inner.Error() == tErr.Error() {
			return true
		}

		inner = errors.Unwrap(inner)
	}
}

func Wrap(e error, friendlyMessage string, misc map[string]interface{}) *Error {
	return &Error{
		Inner:           e,
		FriendlyMessage: friendlyMessage,
		Misc:            misc,
	}
}

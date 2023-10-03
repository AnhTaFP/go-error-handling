package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	t.Run("errors.Is", func(t *testing.T) {
		e1 := Wrap(nil, "foobar", map[string]interface{}{
			"scope": "scopeA",
		})

		e2 := Wrap(e1, "foobar 2", map[string]interface{}{
			"scope": "scopeB",
		})

		e3 := Wrap(e2, "foobar 3", map[string]interface{}{
			"scope": "scopeC",
		})

		e4 := Wrap(e2, "foobar 4", map[string]interface{}{
			"scope": "scopeD",
		})

		assert.True(t, errors.Is(e3, e1))
		assert.True(t, errors.Is(e3, e2))
		assert.True(t, errors.Is(e3, e3))
		assert.False(t, errors.Is(e3, e4))

		emptyErr := errors.New("")
		var x *Error
		var y *Error

		assert.True(t, errors.Is(y, x))
		assert.True(t, errors.Is(x, y))
		assert.False(t, errors.Is(emptyErr, x))
		assert.False(t, errors.Is(x, emptyErr))
	})

	t.Run("errors.As", func(t *testing.T) {
		e1 := Wrap(nil, "foobar", nil)
		e2 := Wrap(e1, "foobar 2", nil)
		e3 := fmt.Errorf("foobar 3: %w", e2)

		var e4 *Error
		assert.True(t, errors.As(e3, &e4))
		assert.Equal(t, "foobar 2: foobar", e4.Error())

		ce := customErr{msg: "custom foobar"}
		e5 := Wrap(ce, "foobar 5", nil)
		e6 := Wrap(e5, "foobar 56", nil)

		var ce2 customErr
		assert.True(t, errors.As(e6, &ce2))
		assert.Equal(t, "custom foobar", ce2.Error())
	})
}

type customErr struct {
	msg string
}

func (ce customErr) Error() string {
	return ce.msg
}

package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	t.Run("wrappable", func(t *testing.T) {
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
	})
}

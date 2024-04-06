package assert_test

import (
	"errors"
	"testing"

	"github.com/fabioelizandro/speech-to-text/assert"
	"github.com/stretchr/testify/require"
)

func TestAssertSuite(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert.NoErr(nil)

		f := func() {
			assert.NoErr(errors.New("foo"))
		}

		require.Panics(t, f)
	})

	t.Run("true", func(t *testing.T) {
		assert.True(true, "s")

		f := func() {
			assert.True(false, "t")
		}

		require.Panics(t, f)
	})

	t.Run("false", func(t *testing.T) {
		assert.False(false, "s")

		f := func() {
			assert.False(true, "t")
		}

		require.Panics(t, f)
	})

	t.Run("unreachable", func(t *testing.T) {
		f := func() {
			assert.Unreachable("t")
		}

		require.Panics(t, f)
	})

	t.Run("must", func(t *testing.T) {
		require.True(t, assert.Must(true, nil))
		require.Equal(t, 1, assert.Must(1, nil))

		require.PanicsWithError(t, "some error", func() {
			assert.Must(false, errors.New("some error"))
		})
	})
}

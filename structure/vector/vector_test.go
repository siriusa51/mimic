package vector

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func equal[T any](t *testing.T, v1, v2 vector[T]) {
	require.Equal(t, len(v1.data), len(v2.data))

}

func TestNew(t *testing.T) {
	{
		v := New()
		require.Equal(t, 0, len(v.data))
	}

	{
		v := New([]int{1, 2, 3, 4})

	}
}

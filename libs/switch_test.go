package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SwitchToMax(t *testing.T) {
	t.Run("Switch To Max when current < max", func(t *testing.T) {
		tab := SwitchToMax(1, 2)
		assert.Equal(t, 2, tab)
	})

	t.Run("Switch To Max when current > max", func(t *testing.T) {
		tab := SwitchToMax(2, 2)
		assert.Equal(t, 2, tab)
	})
}

func Test_SwitchToMin(t *testing.T) {
	t.Run("Switch To Min when current > min", func(t *testing.T) {
		tab := SwitchToMin(1, 0)
		assert.Equal(t, 0, tab)
	})

	t.Run("Switch To Min when current < min", func(t *testing.T) {
		tab := SwitchToMin(0, 0)
		assert.Equal(t, 0, tab)
	})
}

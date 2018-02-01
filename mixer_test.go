package synthia_test

import (
	"testing"

	"github.com/dalloriam/synthia"
)

func TestNewMixer(t *testing.T) {
	t.Run("initializes the correct number of channels", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			m := synthia.NewMixer(i)

			if len(m.Channels) != i {
				t.Errorf("newmixer initialized %d channels, expected %d", len(m.Channels), i)
			}
		}
	})
}

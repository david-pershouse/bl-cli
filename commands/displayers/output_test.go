package displayers

import (
	"bytes"
	"testing"

	"git.mammoth.com.au/github/bl-cli/do"

	"github.com/stretchr/testify/assert"
)

func TestDisplayerDisplay(t *testing.T) {
	emptyImages := make([]do.Image, 0)
	var nilImages []do.Image

	tests := []struct {
		name         string
		item         Displayable
		expectedJSON string
	}{
		{
			name:         "displaying a non-nil slice of Volumes should return an empty JSON array",
			item:         &Image{Images: emptyImages},
			expectedJSON: `[]`,
		},
		{
			name:         "displaying a nil slice of Volumes should return an empty JSON array",
			item:         &Image{Images: nilImages},
			expectedJSON: `[]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}

			displayer := Displayer{
				OutputType: "json",
				Item:       tt.item,
				Out:        out,
			}

			err := displayer.Display()
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedJSON, out.String())
		})
	}
}

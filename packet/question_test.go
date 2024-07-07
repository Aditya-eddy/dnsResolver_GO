package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {

	t.Run("should encode the dns Name", func(t *testing.T) {
		encodeDnsName := encodeDnsName([]byte("dns.google.com"))
		assert.Equal(t, []byte("\x03dns\x06google\x03com\x00"), encodeDnsName)

	})
}

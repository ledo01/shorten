package json

import (
	"fmt"
	"testing"
	"time"

	"github.com/ledo01/shorten/shorten"
	"github.com/stretchr/testify/assert"
)

var (
	now   = time.Now().UTC().Unix()
	input = &shorten.Redirect{
		Code:      "abc",
		URL:       "http://google.com",
		CreatedAt: now,
	}
	inputRaw = []byte(fmt.Sprintf(
		`{"code":"%s","url":"%s","CreatedAt":%d}`,
		input.Code,
		input.URL,
		input.CreatedAt,
	))
)

func TestDecode(t *testing.T) {
	redirect := &Redirect{}
	decoded, err := redirect.Decode(inputRaw)
	assert.Nil(t, err)
	assert.Equal(t, input, decoded)
}

func TestEncode(t *testing.T) {
	redirect := &Redirect{}
	raw, err := redirect.Encode(input)

	assert.Nil(t, err)
	assert.Equal(t, inputRaw, raw)
}

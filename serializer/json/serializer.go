package json

import (
	"encoding/json"

	"github.com/ledo01/shorten/shorten"
	"github.com/pkg/errors"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shorten.Redirect, error) {
	redirect := &shorten.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serizalizer.Redirect.Decode")
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *shorten.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)

	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}

	return rawMsg, nil
}

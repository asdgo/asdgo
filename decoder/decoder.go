package decoder

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/go-playground/form"
)

var Instance *Decoder

func Init() {
	d := form.NewDecoder()

	Instance = &Decoder{
		decoder: d,
	}
}

type Decoder struct {
	decoder *form.Decoder
}

func (d *Decoder) Struct(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	return d.UrlValues(r.PostForm, dst)
}

func (d *Decoder) UrlValues(v url.Values, dst any) error {
	err := d.decoder.Decode(dst, v)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
	}

	return err
}

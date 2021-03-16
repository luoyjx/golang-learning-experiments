package encode

import (
	"fmt"
	"strings"

	"github.com/francoispqt/gojay"
)

var testStr1 = strings.Repeat("abc", 10000)

type data map[string]interface{}

func (d data) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range d {
		fmt.Println(k, v)
		enc.AddInterfaceKey(k, v)
	}
}
func (d data) IsNil() bool {
	return d == nil
}

// Our MarshalerStream implementation
type StreamChan chan data

func (s StreamChan) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case <-enc.Done():
		return
	case o := <-s:
		_ = enc.Encode(o)
	}
}


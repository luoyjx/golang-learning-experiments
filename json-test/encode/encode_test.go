package encode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var m = map[string]interface{}{
	"a": 1,
	"b": strings.Repeat("abcde", 10),
	"c": int64(2),
	"d": map[string]string{"foo": strings.Repeat("bar12foo34", 100000)},
	"e": []int{1, 2},
}

func Test_StreamEncode(t *testing.T) {
	builder := &strings.Builder{}

	enc := json.NewEncoder(builder)
	_ = enc.Encode(m)

	//enc := gojay.Stream.BorrowEncoder(builder)
	//defer enc.Release()
	//// instantiate our MarshalerStream
	//s := StreamChan(make(chan data))
	//// start the stream encoder
	//// will block its goroutine until enc.Cancel(error) is called
	//// or until something is written to the channel
	//go enc.EncodeStream(s)
	//// write to our MarshalerStream
	//for i := 0; i < 1; i++ {
	//	s <- data{
	//		"a": 1,
	//		"b": testStr1,
	//		"c": int64(2),
	//		"d": map[string]string{"foo": "bar"},
	//	}
	//}
	//// Wait
	//<-enc.Done()
	fmt.Println("result", builder.String())

	b := &strings.Builder{}
	stream := jsoniter.ConfigFastest.BorrowStream(b)
	stream.WriteVal(m)
	_ = stream.Flush()
	fmt.Println("res ", b.String())
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(m)
	}
}

func BenchmarkJSONEncoder(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		enc := json.NewEncoder(ioutil.Discard)
		_ = enc.Encode(m)
	}
}

func BenchmarkJSONIterator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stream := jsoniter.ConfigFastest.BorrowStream(ioutil.Discard)
		stream.WriteVal(m)
		_ = stream.Flush()
	}
}

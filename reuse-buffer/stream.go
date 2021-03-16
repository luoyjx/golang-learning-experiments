package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"syscall"
	"testing"
)

var testStr1 = strings.Repeat("abc", 10000)
var testStr2 = strings.Repeat("cde", 10000)
var testStr3 = strings.Repeat("def", 10000)
var testStr4 = strings.Repeat("efg", 10000)
var testStr = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s", testStr1, testStr2, testStr3, testStr4, testStr1, testStr2, testStr3, testStr4)

type NopWriterCloser struct{}

func (nwc *NopWriterCloser) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (nwc *NopWriterCloser) Close() error {
	return nil
}

func (nwc *NopWriterCloser) Flush() {}

var bufioReaderPool sync.Pool
func newBufioReader(r io.Reader) *bufio.Reader {
	if v := bufioReaderPool.Get(); v != nil {
		br := v.(*bufio.Reader)
		br.Reset(r)
		return br
	}
	return bufio.NewReader(r)
}
func putBufioReader(br *bufio.Reader) {
	br.Reset(nil)
	bufioReaderPool.Put(br)
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4*1024)
	},
}

func testStreamWrite(w io.Writer, testStr string) {
	flusher := w.(http.Flusher)

	reader := newBufioReader(strings.NewReader(testStr))
	defer putBufioReader(reader)
	buf := bufPool.Get().([]byte)
	defer bufPool.Put(buf)
	isEOF := false
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				// 有可能 reader 已经关闭，这里跳出
				// logrus.Errorf("read response error %v", err)
				break
			} else {
				// EOF 时，可能还有读出的数据
				isEOF = true
			}
		}

		if n == 0 && !isEOF {
			continue
		}

		if n > 0 {
			wroteN, err := w.Write(buf[:n])
			if err != nil {
				if errors.Is(err, syscall.EPIPE) {
					break
				} else {
					// logrus.Errorf("write line with error : %+v ", err)
				}
			}

			if wroteN > 0 {
				flusher.Flush()
			}
		}

		if isEOF {
			break
		}
	}
}

func writeByFprint() {
	w := &NopWriterCloser{}
	_, _ = fmt.Fprint(w, testStr)
}

func writeByStream() {
	w := &NopWriterCloser{}
	testStreamWrite(w, testStr)
}
func main() {
	fmt.Println("AllocsPerRun fmt.Fprint ", testing.AllocsPerRun(100, writeByFprint))
	fmt.Println("AllocsPerRun testStreamWrite ", testing.AllocsPerRun(100, writeByStream))
}

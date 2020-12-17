package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"testing"
)

var testStr1 = strings.Repeat("abc", 10000)
var testStr2 = strings.Repeat("cde", 10000)
var testStr = fmt.Sprintf("%s\n%s\n%s\n%s", testStr1, testStr2, testStr1, testStr2)

func readBytesByLine() {
	rc := ioutil.NopCloser(strings.NewReader(testStr))
	reader := bufio.NewReader(rc)

	loopCount := 0

	for {
		_, err := reader.ReadBytes('\n')
		if err == io.EOF {
			rc.Close()
			break
		}

		loopCount++
	}
}

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
		return make([]byte, 10*1024)
	},
}

func readFromBuffer() {
	rc := ioutil.NopCloser(strings.NewReader(testStr))
	reader := newBufioReader(rc)

	buf := bufPool.Get().([]byte)

	loopCount := 0

	for {
		_, err := reader.Read(buf)

		if err == io.EOF {
			break
		}

		loopCount++
	}

	putBufioReader(reader)
	bufPool.Put(buf)
}

func main() {
	fmt.Println("AllocsPerRun readBytesByLine ", testing.AllocsPerRun(100, readBytesByLine))
	fmt.Println("AllocsPerRun readFromBuffer ", testing.AllocsPerRun(100, readFromBuffer))
}

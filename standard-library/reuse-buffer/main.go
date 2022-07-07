package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"testing"
)

var testStr1 = strings.Repeat("abc", 10000)
var testStr2 = strings.Repeat("cde", 10000)
var testStr3 = strings.Repeat("def", 10000)
var testStr4 = strings.Repeat("efg", 10000)
var testStr = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s", testStr1, testStr2, testStr3, testStr4, testStr1, testStr2, testStr3, testStr4)
var strBytesLen = len(testStr)

func readBytesByLine() {
	rc := ioutil.NopCloser(strings.NewReader(testStr))
	newBuf := bytes.NewBuffer(make([]byte, 0, strBytesLen))
	reader := bufio.NewReader(rc)

	c := 0
	for {
		bs, err := reader.ReadBytes('\n')
		if err == io.EOF {
			if len(bs) > 0 {
				c += len(bs)
				newBuf.Write(bs)
			}

			rc.Close()
			break
		}

		c += len(bs)
		newBuf.Write(bs)
	}
	// fmt.Println(string(newBuf.Bytes()) == testStr)
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
	newBuf := bytes.NewBuffer(make([]byte, 0, strBytesLen))
	reader := newBufioReader(rc)
	buf := bufPool.Get().([]byte)

	for {
		n, err := reader.Read(buf)

		if err == io.EOF {
			if n > 0 {
				newBuf.Write(buf[:n])
			}
			rc.Close()
			break
		}

		_, err = newBuf.Write(buf[:n])

		if err != nil {
			panic(err)
		}
	}

	putBufioReader(reader)
	bufPool.Put(buf)
	// fmt.Println(string(newBuf.Bytes()) == testStr)
}

func main() {
	fmt.Println("AllocsPerRun readBytesByLine ", testing.AllocsPerRun(100, readBytesByLine))
	fmt.Println("AllocsPerRun readFromBuffer ", testing.AllocsPerRun(100, readFromBuffer))
}

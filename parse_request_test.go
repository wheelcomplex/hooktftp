
package main

import (
	"fmt"
	"testing"
	"encoding/binary"
	"strconv"
)

func createRRQTestData(path string, mode string, blocksize int) ([]byte) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, RRQ)
	data = append(data, []byte(path)...)
	data = append(data, 0)
	data = append(data, []byte(mode)...)
	if blocksize != -1 {
		data = append(data, 0)
		data = append(data, []byte("blksize")...)
		data = append(data, 0)
		data = append(data, []byte(strconv.Itoa(blocksize))...)
		data = append(data, 0)
	}
	return data
}

func TestRequestDataWithCustomBlockSize(t *testing.T) {

	data := createRRQTestData("foo.txt", "octet", 5)

	rrq, err := ParseRequest(data)
	if err != nil {
		t.Errorf("Got error %v", err)
		return
	}

	if rrq.path != "foo.txt" {
		t.Errorf("Bad path: %v", rrq.path)
		return
	}

	if rrq.mode != OCTET {
		t.Errorf("Bad mode: %v", rrq.mode)
		return
	}

	if rrq.blocksize != 5 {
		t.Errorf("Bad blocksize: %v", rrq.blocksize)
		return
	}

	fmt.Println(rrq)

}

func TestRequestDataWithDefaultBlockSize(t *testing.T) {
	data := createRRQTestData("foo.txt", "octet", -1)

	rrq, err := ParseRequest(data)
	if err != nil {
		t.Errorf("Got error %v", err)
		return
	}

	if rrq.blocksize != DEFAULT_BLOCKSIZE {
		t.Errorf("Bad blocksize: %v", rrq.blocksize)
		return
	}
}

func TestIgnoreUnknownOptions(t *testing.T) {
	data := createRRQTestData("foo.txt", "octet", -1)
	data = append(data, 0)
	data = append(data, []byte("bad")...)
	data = append(data, 0)
	data = append(data, []byte("badvalue")...)
	data = append(data, 0)

	data = append(data, []byte("blksize")...)
	data = append(data, 0)
	data = append(data, []byte(strconv.Itoa(5))...)
	data = append(data, 0)

	rrq, err := ParseRequest(data)
	if err != nil {
		t.Errorf("Got error %v", err)
		return
	}

	if rrq.blocksize != 5 {
		t.Errorf("Bad blocksize: %v", rrq.blocksize)
		return
	}
}


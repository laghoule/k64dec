package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

const (
	yamlStringDataGold = "../testdata/secretStringData.yaml"
	yamlDataGold       = "../testdata/secretData.yaml"
	jsonDataGold       = "../testdata/secretData.json"
	bytesDataGold      = "../testdata/secretData.bytes"
	badSecret          = "../testdata/badSecret"
)

func captureConsoleOutput(f func()) []byte {
	var buf bytes.Buffer
	pterm.SetDefaultOutput(&buf)

	f()

	pterm.SetDefaultOutput(os.Stderr)
	return buf.Bytes()
}

func TestPrintVersion(t *testing.T) {
	captured := captureConsoleOutput(
		func() {
			if err := printVersion(); err != nil {
				t.Error(err)
				return
			}
		},
	)

	expected := []byte{0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x39, 0x36, 0x6d, 0x1b, 0x5b, 0x39, 0x36, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x39, 0x30, 0x6d, 0x1b, 0x5b, 0x39, 0x30, 0x6d, 0x20, 0x7c, 0x20, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x39, 0x36, 0x6d, 0x1b, 0x5b, 0x39, 0x36, 0x6d, 0x47, 0x69, 0x74, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x39, 0x30, 0x6d, 0x1b, 0x5b, 0x39, 0x30, 0x6d, 0x20, 0x7c, 0x20, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x39, 0x36, 0x6d, 0x1b, 0x5b, 0x39, 0x36, 0x6d, 0x47, 0x69, 0x74, 0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0xa, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x20, 0x20, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x39, 0x30, 0x6d, 0x1b, 0x5b, 0x39, 0x30, 0x6d, 0x20, 0x7c, 0x20, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x39, 0x30, 0x6d, 0x1b, 0x5b, 0x39, 0x30, 0x6d, 0x20, 0x7c, 0x20, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x1b, 0x5b, 0x33, 0x39, 0x6d, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0xa}
	assert.Equal(t, expected, captured)
}

func TestReadFromSTDIN(t *testing.T) {
	file, err := os.Open(yamlDataGold)
	if err != nil {
		t.Error(err)
		return
	}

	defer file.Close()
	os.Stdin = file

	_, err = file.Read([]byte{})
	if err != nil {
		t.Error(err)
		return
	}

	data, err := readFromSTDIN()
	if err != nil {
		t.Error(err)
		return
	}

	expected := []byte{0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x76, 0x31, 0xa, 0x64, 0x61, 0x74, 0x61, 0x3a, 0xa, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x20, 0x62, 0x58, 0x6b, 0x67, 0x63, 0x48, 0x4a, 0x6c, 0x59, 0x32, 0x6c, 0x76, 0x64, 0x58, 0x4d, 0x67, 0x59, 0x32, 0x39, 0x75, 0x5a, 0x6d, 0x6c, 0x6e, 0xa, 0x6b, 0x69, 0x6e, 0x64, 0x3a, 0x20, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0xa, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0xa, 0x20, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x39, 0x2d, 0x32, 0x34, 0x54, 0x30, 0x33, 0x3a, 0x35, 0x31, 0x3a, 0x34, 0x37, 0x5a, 0x22, 0xa, 0x20, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x61, 0x70, 0x70, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x62, 0x79, 0x3a, 0x20, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0xa, 0x20, 0x20, 0x20, 0x20, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x3a, 0x20, 0x70, 0x61, 0x63, 0x6d, 0x61, 0x6e, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa, 0x20, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x70, 0x61, 0x63, 0x6d, 0x61, 0x6e, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa, 0x20, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x20, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0xa, 0x20, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x22, 0x31, 0x33, 0x36, 0x34, 0x36, 0x37, 0x38, 0x22, 0xa, 0x20, 0x20, 0x75, 0x69, 0x64, 0x3a, 0x20, 0x33, 0x34, 0x36, 0x36, 0x66, 0x66, 0x37}
	assert.Equal(t, expected, data)
}

func createTempFile(size int) (string, error) {
	f, err := os.CreateTemp(os.TempDir(), "k64dec-")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file %s/%s", os.TempDir(), f.Name())
	}

	defer f.Close()

	var data []byte
	for i := 0; i <= size; i++ {
		data = append(data, 0)
	}

	_, err = f.Write(data)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

func TestReadFromSTDINMaxsize(t *testing.T) {
	fileName, err := createTempFile(maxSize)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if err := os.Remove(fileName); err != nil {
			t.Error(err)
			return
		}
	}()

	f, err := os.Open(fileName)
	if err != nil {
		t.Error(t)
		return
	}

	os.Stdin = f

	_, err = f.Read([]byte{})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = readFromSTDIN()
	assert.ErrorContains(t, err, "max read buffer reach")
}

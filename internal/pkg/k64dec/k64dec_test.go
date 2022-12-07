package k64dec

import (
	"bytes"
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

const (
	yamlStringDataGold = "testdata/secretStringData.yaml"
	yamlDataGold       = "testdata/secretData.yaml"
	jsonDataGold       = "testdata/secretData.json"
	bytesDataGold      = "testdata/secretData.bytes"
	badSecret          = "testdata/badSecret"
)

func TestDecode(t *testing.T) {
	file, err := os.ReadFile(yamlDataGold)
	if err != nil {
		t.Error(err)
		return
	}

	s, err := decode(file)
	if err != nil {
		t.Error(err)
		return
	}

	expected, err := os.ReadFile(bytesDataGold)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, string(expected), s.String())
}

func TestDecodeBadSecret(t *testing.T) {
	file, err := os.ReadFile(badSecret)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = decode(file)
	assert.Error(t, err)
}

func captureConsoleOutput(f func()) []byte {
	var buf bytes.Buffer
	pterm.SetDefaultOutput(&buf)

	f()

	pterm.SetDefaultOutput(os.Stderr)
	return buf.Bytes()
}

func TestPrintDecodedSecret(t *testing.T) {
	files := []string{yamlDataGold, yamlStringDataGold}

	for _, testFile := range files {
		file, err := os.ReadFile(testFile)
		if err != nil {
			t.Error(err)
			return
		}
	
		captured := captureConsoleOutput(
			func() {
				if err := PrintDecodedSecret(file); err != nil {
					t.Error(err)
					return
				}
			},
		)
	
		expected := []byte{0x1b, 0x5b, 0x34, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1b, 0x5b, 0x30, 0x6d, 0xa, 0x1b, 0x5b, 0x34, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x6d, 0x6d, 0x79, 0x20, 0x70, 0x72, 0x65, 0x63, 0x69, 0x6f, 0x75, 0x73, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1b, 0x5b, 0x30, 0x6d}
		assert.Equal(t, expected, captured)
	}
}

func TestPrintDecodedBadSecret(t *testing.T) {
	file, err := os.ReadFile(badSecret)
	if err != nil {
		t.Error(err)
		return
	}

	err = PrintDecodedSecret(file)
	assert.Error(t, err)
}

func TestPrint(t *testing.T) {
	captured := captureConsoleOutput(
		func() {
			print("key", "value")
		},
	)

	expected := []byte{0x1b, 0x5b, 0x34, 0x6d, 0x6b, 0x65, 0x79, 0x1b, 0x5b, 0x30, 0x6d, 0xa, 0x1b, 0x5b, 0x34, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x6d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1b, 0x5b, 0x30, 0x6d}
	assert.Equal(t, expected, captured)
}

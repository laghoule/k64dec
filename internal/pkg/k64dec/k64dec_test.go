package k64dec

import (
	"bytes"
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

const (
	yamlGold  = "testdata/secret.yaml"
	jsonGold  = "testdata/secret.json"
	bytesGold = "testdata/secret.bytes"
)

func TestDecode(t *testing.T) {
	file, err := os.ReadFile(yamlGold)
	if err != nil {
		t.Error(err)
		return
	}

	s, err := decode(file)
	if err != nil {
		t.Error(err)
		return
	}

	expected, err := os.ReadFile(bytesGold)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, string(expected), s.String())
}

func captureConsoleOutput(f func()) []byte {
	var buf bytes.Buffer
	pterm.SetDefaultOutput(&buf)

	f()

	pterm.SetDefaultOutput(os.Stderr)
	return buf.Bytes()
}

func TestPrint(t *testing.T) {
	captured := captureConsoleOutput(
		func() {
			print("key", "value")
		},
	)

	var expected = []byte{0x1b, 0x5b, 0x34, 0x6d, 0x6b, 0x65, 0x79, 0x1b, 0x5b, 0x30, 0x6d, 0xa, 0x1b, 0x5b, 0x34, 0x6d, 0x1b, 0x5b, 0x30, 0x6d, 0x1b, 0x5b, 0x33, 0x6d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1b, 0x5b, 0x30, 0x6d}

	assert.Equal(t, expected, captured)
}

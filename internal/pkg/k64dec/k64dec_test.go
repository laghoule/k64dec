package k64dec

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

const (
	testDataDir          = "../../../testdata"
	testDataExpectedDir  = testDataDir + "/expected"
	secretDataTLSYaml    = "secretDataTLSYaml"
	secretStringDataYaml = "secretStringDataYaml"
	secretDataYaml       = "secretDataYaml"
	secretDataJson       = "secretDataJson"
	badSecret            = "badSecret"
)

func resolvePath(path, file string) string {
	return fmt.Sprintf("%s/%s", path, file)
}

func readFile(t *testing.T, fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		t.Errorf("failed to read file %s", fileName)
		return nil
	}
	return file
}

func TestDecode(t *testing.T) {
	var files = []string{
		secretDataYaml,
		secretDataJson,
		secretDataTLSYaml,
	}

	for _, testFile := range files {
		f := readFile(t, resolvePath(testDataDir, testFile))

		s, err := decode(f)
		if err != nil {
			t.Error(err)
			return
		}

		expected := readFile(t, resolvePath(testDataExpectedDir, testFile+"-TestDecode"))
		assert.Equal(t, string(expected), s.String())
	}
}

func TestDecodeBadSecret(t *testing.T) {
	file := readFile(t, resolvePath(testDataDir, badSecret))

	_, err := decode(file)
	assert.Error(t, err)
}

func captureConsoleOutput(f func()) []byte {
	var buf bytes.Buffer
	pterm.SetDefaultOutput(&buf)
	pterm.DisableColor()
	pterm.DisableStyling()

	f()

	pterm.SetDefaultOutput(os.Stderr)
	return buf.Bytes()
}

func TestPrintDecodedSecret(t *testing.T) {
	files := []string{
		secretDataYaml,
		secretDataJson,
		secretDataTLSYaml,
		secretStringDataYaml,
	}

	for _, testFile := range files {
		file := readFile(t, resolvePath(testDataDir, testFile))

		captured := captureConsoleOutput(
			func() {
				if err := PrintDecodedSecret(file); err != nil {
					t.Error(err)
					return
				}
			},
		)

		expected := readFile(t, resolvePath(testDataExpectedDir, testFile+"-TestPrintDecodedSecret"))
		assert.Equal(t, string(expected), string(captured))
	}
}

func TestPrintDecodedBadSecret(t *testing.T) {
	file := readFile(t, resolvePath(testDataDir, badSecret))

	err := PrintDecodedSecret(file)
	assert.Error(t, err)
}

func TestPrint(t *testing.T) {
	captured := captureConsoleOutput(
		func() {
			print("key", "value")
		},
	)

	expected := "key\nvalue"
	assert.Equal(t, expected, string(captured))
}

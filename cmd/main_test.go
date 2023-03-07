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
	testDataDir          = "../testdata"
	testDataExpectedDir  = testDataDir + "/expected"
	secretDataTLSYaml    = "secretDataTLSYaml"
	secretStringDataYaml = "secretStringDataYaml"
	secretDataYaml       = "secretDataYaml"
	secretDataJson       = "secretDataJson"
	badSecret            = "badSecret"
)

func captureConsoleOutput(f func()) []byte {
	var buf bytes.Buffer
	pterm.SetDefaultOutput(&buf)
	pterm.DisableColor()

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

	expected := "Version | Git commit | Git reference\ndevel   |            | \n\n"
	assert.Equal(t, expected, string(captured))
}

func TestReadFromSTDIN(t *testing.T) {
	var testCase = []string{
		secretDataYaml,
		secretDataJson,
		secretDataTLSYaml,
	}

	for _, test := range testCase {
		secretFile := fmt.Sprintf("%s/%s", testDataDir, test)

		input, err := os.Open(secretFile)
		if err != nil {
			t.Error(err)
			return
		}

		defer input.Close()
		os.Stdin = input

		_, err = input.Read([]byte{})
		if err != nil {
			t.Error(err)
			return
		}

		data, err := readFromSTDIN()
		if err != nil {
			t.Error(err)
			return
		}

		expected, err := os.ReadFile(secretFile)
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, expected, data)
	}

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

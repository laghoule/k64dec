package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/laghoule/k64dec/internal/pkg/k64dec"

	"github.com/pterm/pterm"
	flag "github.com/spf13/pflag"
)

const (
	// 1MiB + 1KiB
	// secrets are limited to 1MiB
	bufSize = 1048576 + 1024
)

var (
	version   = "devel"
	gitCommit = ""
	gitRef    = ""
)

func main() {
	fileName := flag.String("file", "", "kubernetes secret in json or yaml file")
	version := flag.Bool("version", false, "print version")
	flag.Parse()

	if *version {
		if err := printVersion(); err != nil {
			exitOnError(err)
		}
		return
	}

	var data []byte
	var err error

	if *fileName == "" {
		data, err = readFromSTDIN()
		if err != nil {
			exitOnError(err)
		}
	} else {
		data, err = os.ReadFile(*fileName)
		if err != nil {
			exitOnError(err)
		}
	}

	if err := k64dec.PrintDecodedSecret(data); err != nil {
		exitOnError(err)
	}
}

// readFromSTDIN read data from standard input
func readFromSTDIN() ([]byte, error) {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, bufSize)

	var data []byte

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			return []byte{}, fmt.Errorf("read from STDIN failed: %s", err)
		}

		if err != nil && err != io.EOF {
			return []byte{}, fmt.Errorf("read from STDIO failed: %s", err)
		}

		data = buf
	}

	return data, nil
}

// printVersion print version of k64dec
func printVersion() error {
	var pdata = pterm.TableData{
		{"Version", "Git commit", "Git reference"},
		{version, gitCommit, gitRef},
	}
	if err := pterm.DefaultTable.WithHasHeader().WithData(pdata).Render(); err != nil {
		return fmt.Errorf("failed to print version: %s", err)
	}
	return nil
}

func exitOnError(err error) {
	fmt.Printf("error: %s\n", err.Error())
	os.Exit(1)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/laghoule/k64dec/internal/pkg/k64dec"

	"github.com/pterm/pterm"
	flag "github.com/spf13/pflag"
)

const (
	bufSize = 4 * 1024
)

var (
	version   = "devel"
	buildDate = time.Now().String()
	gitCommit = ""
	gitRef    = ""
)

func main() {
	fileName := flag.String("file", "", "kubernetes secret in json or yaml file")
	version := flag.Bool("version", false, "print version")
	flag.Parse()

	if *version {
		printVersion()
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
func printVersion() {
	var pdata = pterm.TableData{
		{"Version", "Build date", "Git commit", "Git reference"},
		{version, buildDate, gitCommit, gitRef},
	}
	if err := pterm.DefaultTable.WithHasHeader().WithData(pdata).Render(); err != nil {
		exitOnError(err)
	}
}

func exitOnError(err error) {
	fmt.Printf("error: %s\n", err.Error())
	os.Exit(1)
}

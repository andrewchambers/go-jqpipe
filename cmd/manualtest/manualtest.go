package main

import (
	"fmt"
	"io"
	"os"

	"github.com/andrewchambers/go-jqpipe"
)

func main() {
	jq, err := jqpipe.New(os.Stdin, os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	for {
		next, err := jq.Next()
		switch {
		case err == nil:
			_, _ = fmt.Fprintf(os.Stdout, "%s", next)
		case err == io.EOF:
			jq.Close()
			return
		default:
			_, _ = fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
	}
}

// Package mobingi is a command line interface client for Mobingi services.
package main

import (
	"log"

	"github.com/mobingi/mobingi/cmd"
	"github.com/mobingi/sdk-go/pkg/util/cmdline"
)

func main() {
	pfx := "[" + cmdline.Args0() + "]: "
	log.SetPrefix(pfx)
	log.SetFlags(0)
	cmd.Execute()
}

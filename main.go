// Package mobingi is a command line interface client for Mobingi services.
package main

import (
	"log"

	"github.com/mobingi/gosdk/pkg/util/cmdline"
	"github.com/mobingi/mobingi/cmd"
)

func main() {
	pfx := "[" + cmdline.Args0() + "]: "
	log.SetPrefix(pfx)
	log.SetFlags(0)
	cmd.Execute()
}

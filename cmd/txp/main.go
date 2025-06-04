package main

import (
	"flag"
	"path"

	"github.com/charmbracelet/log"
	"github.com/unhanded/txp/internal/txppack"
)

func main() {
	var isDebug bool
	flag.BoolVar(&isDebug, "d", false, "enable debug logging")
	flag.Parse()
	command := flag.Arg(0)
	target := flag.Arg(1)
	switch command {
	case "validate":
		validate(target, isDebug)
	}
}

func validate(dirpath string, isDebug bool) {
	if isDebug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	pack, err := txppack.Load(dirpath)
	if err != nil {
		log.Fatal("Failed to load", "err", err)
	}
	if err := pack.Validate(); err != nil {
		log.Error("Validation failed", "err", err)
	} else {
		log.Info("Validation successful!", "filecount", len(pack.FileList))
	}

	for _, file := range pack.FileList {
		filename := path.Base(file)
		log.Debugf("Got file \"%s\"", filename)
	}
}

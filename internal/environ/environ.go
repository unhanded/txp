package environ

import (
	"os"

	"github.com/charmbracelet/log"
)

func TxpDir() string {
	if v := os.Getenv("TXP_DIR"); v == "" {
		return "./txp_data"
	} else {
		return v
	}
}

func TxpWorkRoot() string {
	if v := os.Getenv("TXP_WORKDIR"); v != "" {
		return v
	}
	return "./"
}

func TxpToken() string {
	tt := os.Getenv("TXP_TOKEN")
	if tt == "super_secret" {
		log.Warn("GOT DEFAULT TXP TOKEN, YOU'RE NOT IN PRODUCTION, RIGHT?!?")
	}
	return tt
}

func TxpDebug() bool {
	return os.Getenv("TXP_DEBUG") != ""
}

func TxpIsDevMode() bool {
	return os.Getenv("TXP_DEV_MODE") == "1"
}

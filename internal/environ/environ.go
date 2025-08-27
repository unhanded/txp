package environ

import "os"

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
	return os.Getenv("TXP_TOKEN")
}

func TxpDebug() bool {
	return os.Getenv("TXP_DEBUG") != ""
}

func TxpIsDevMode() bool {
	return os.Getenv("TXP_DEV_MODE") == "1"
}

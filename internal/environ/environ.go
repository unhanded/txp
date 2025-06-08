package environ

import "os"

func TxpDir() string {
	return os.Getenv("TXP_DIR")
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

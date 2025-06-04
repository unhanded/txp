package dataman

import "os"

const TXP_DEFAULT_DIR = "/txp_data"

func GetTxpDir() string {
	fromEnv := os.Getenv("TXP_DIR")
	if fromEnv == "" {
		return TXP_DEFAULT_DIR
	} else {
		return fromEnv
	}
}

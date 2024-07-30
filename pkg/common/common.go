package common

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ChkFatal(err error, errMsg string) {
	if err != nil {
		log.Fatal(errMsg, " (", err, ")")
	}
}

func ChkWarn(err error, errMsg string) {
	if err != nil {
		log.Println(errMsg, " (", err, ")")
	}
}

func SysInterrupt() chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	return sigChan
}

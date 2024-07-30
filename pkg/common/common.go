package common

import "log"

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

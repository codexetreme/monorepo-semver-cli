package error_roster

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	RUNTIME_ERROR    = 1
	INVALID_INPUT    = 2
	NOT_IMPLEMENTED  = 3
	JSON_PARSE_ERROR = 4
	YAML_PARSE_ERROR = 5
	HTTP_ERROR       = 6
	SERVER_ERROR     = 7
)

var exitCode int = 0

func SetErrorCode(code int) {
	exitCode = code
}

// CheckErr prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is nil, it does nothing.
func CheckErr(msg interface{}) {
	if msg != nil {
		log.Error("Error:", msg)
		os.Exit(exitCode)
	}
}

func CheckForArgs(args []string) error {
	if len(args) == 0 {
		SetErrorCode(INVALID_INPUT)
		return errors.New("argument missing")
	} else {
		return nil
	}
}

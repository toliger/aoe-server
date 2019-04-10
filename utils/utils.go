package utils

import (
	rlog "log"
	"os"
	"strconv"
)

//Getenv : Get an environ var
func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

//Debug : Debug print
func Debug(msg string) {
	b, _ := strconv.ParseBool(Getenv("DEBUG", "false"))

	if b {
		rlog.Print(msg)
	}
}

package main

import (
    "log"
    "fmt"
    "time"
)

const (
    TAG = "ZKEY"
    PORT = ":50000"
)


func rep(char byte, count int) []byte {
	var result = make([]byte, count)
	for i, _ := range result {
		result[i] = char
	}

	return result
}


func _P(message string, params ...interface{}) {
    if len(params) > 0 {
        message = fmt.Sprintf(message, params...)
    }

    log.Printf(fmt.Sprintf("[%s] %s", TAG, message))
}


func GetTime() uint64 {
    return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}

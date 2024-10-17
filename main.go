package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type Log struct {
	Level LogLevel
}

func (l Log) Log(message string) {
	fmt.Printf("%s: %s", l.Level, message)
}

type LogLevel string

const (
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
)

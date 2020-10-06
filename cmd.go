package main

import (
	"os"

	"github.com/toeydevelopment/golang-traning-2020/channel"
	"github.com/toeydevelopment/golang-traning-2020/concurrency"
	"github.com/toeydevelopment/golang-traning-2020/datastructure"
	"github.com/toeydevelopment/golang-traning-2020/datatype"
	"github.com/toeydevelopment/golang-traning-2020/go_context"
	"github.com/toeydevelopment/golang-traning-2020/go_extension"
	"github.com/toeydevelopment/golang-traning-2020/go_testing"
	"github.com/toeydevelopment/golang-traning-2020/goroutine"
	"github.com/toeydevelopment/golang-traning-2020/loop"
)

func cmd() {
	arg := os.Args[1]

	switch arg {
	case "channel":
		channel.Run()
	case "concurrency":
		concurrency.Run()
	case "datastructure":
		datastructure.Run()
	case "datatype":
		datatype.Run()
	case "go_context":
		go_context.Run()
	case "goroutine":
		goroutine.Run()
	case "loop":
		loop.Run()
	case "testing":
		go_testing.Run()
	case "go_extension":
		go_extension.Run()
	default:
		panic("Unknow Command")
	}
}

package main

import (
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/tipsio/tips"
	"github.com/tipsio/tips/conf"
)

var addr = "127.0.0.1:12345"

var url = "http://127.0.0.1:12345"

func TestMain(m *testing.M) {
	conf := &conf.Server{}
	pubsub, _ := tips.MockTips()
	server := NewServer(conf, pubsub)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	go server.Serve(lis)
	time.Sleep(time.Second)
	v := m.Run()
	os.Exit(v)
}

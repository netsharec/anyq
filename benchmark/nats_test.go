package main

import (
	"github.com/jaehue/qlib"
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkNatsProduce(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	q, err := qlib.Setup("nats", "nats://192.168.81.43:4222")
	if err != nil {
		b.Error(err)
	}
	consumeBenchmark(q, b)
}

func BenchmarkNatsConsume(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	q, err := qlib.Setup("nats", "nats://192.168.81.43:4222")
	if err != nil {
		b.Error(err)
	}
	produceBenchmark(q, b)
}
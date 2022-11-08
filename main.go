package main

import (
	"context"
	"fmt"
	"github.com/open-policy-agent/opa/logging"
	"github.com/open-policy-agent/opa/storage/disk"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/term"
	"syscall"
)

func opa(input []byte) (*disk.Store, error) {
	opts, err := disk.OptionsFromConfig(input, "bar")
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &disk.Options{}
	}
	logger := logging.New()
	return disk.New(context.Background(), logger, prometheus.DefaultRegisterer, *opts)
}

func main() {
	input, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		panic(err)
	}
	opt, err := opa(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("opt: %#v\n", opt)
}

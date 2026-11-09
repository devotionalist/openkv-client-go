// Copyright 2026 devotionalist
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/devotionalist/openkv-client-go/kv"
)

func main() {
	addr := "127.0.0.1:19000"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli, err := kv.Dial(ctx, addr)
	if err != nil {
		fatal(err)
	}
	defer cli.Close()

	key := []byte("clients/go/demo")
	if _, err := cli.PutTTL(ctx, key, []byte("hello"), 0); err != nil {
		fatal(err)
	}
	val, found, rev, err := cli.Get(ctx, key)
	if err != nil {
		fatal(err)
	}
	fmt.Printf("get found=%v rev=%d val=%s\n", found, rev, val)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

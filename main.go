package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"

	"codeberg.org/peterzam/socks5"
	"golang.org/x/net/proxy"
)

var (
	bindAddr    = flag.String("bind", "0.0.0.0:1080", "Bind Address for SOCKS5")
	csvfilepath = flag.String("csv", "socks.csv", "SOCKS5 server list csv file")
)

func main() {
	flag.Parse()

	var t []proxy.Dialer
	f, err := os.ReadFile(*csvfilepath)
	if err != nil {
		fmt.Println("----------------")
		fmt.Println("CSV file read error")
		fmt.Println("----------------")
		panic(err)
	}
	sockstrings := bytes.Split(f, []byte("\n"))

	for _, i := range sockstrings {
		j, _ := proxy.SOCKS5("tcp", string(i), &proxy.Auth{}, nil)
		t = append(t, j)
	}

	var i int
	server, _ := socks5.New(&socks5.Config{
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			if ctx.Done() == nil {
				i = rand.Intn(len(t))
			}
			return t[i].Dial(network, addr)
		},
	})
	if err := server.ListenAndServe("tcp", *bindAddr); err != nil {
		fmt.Println("----------------")
		fmt.Println("SOCKS5 server cannot be started at : ", *bindAddr)
		fmt.Println("----------------")
		panic(err)
	}
}

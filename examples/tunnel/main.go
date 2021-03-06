package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/at-wat/mqtt-go"
	"github.com/seqsense/aws-iot-device-sdk-go/v4"
	"github.com/seqsense/aws-iot-device-sdk-go/v4/tunnel"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if len(os.Args) != 3 {
		println("usage: tunnel AWS_IOT_ENDPOINT THING_NAME")
		println("")
		println("This example creates AWS IoT Secure Tunneling destination client")
		println("to the local SSH port.")
		println("THING_NAME must be registered to your account of AWS IoT beforehand.")
		println("")
		println("Following files must be placed under the current working directory:")
		println("         root-CA.crt: root CA certificate")
		println(" certificate.pem.crt: client certificate associated to THING_NAME")
		println("     private.pem.key: private key associated to THING_NAME")
		os.Exit(1)
	}
	host := os.Args[1]
	thingName := os.Args[2]

	for _, file := range []string{
		"root-CA.crt",
		"certificate.pem.crt",
		"private.pem.key",
	} {
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			println(file, "not found")
			os.Exit(1)
		}
	}

	cli, err := awsiotdev.New(
		thingName,
		&mqtt.URLDialer{
			URL: fmt.Sprintf("mqtts://%s:8883", host),
			Options: []mqtt.DialOption{
				mqtt.WithTLSCertFiles(
					host,
					"root-CA.crt",
					"certificate.pem.crt",
					"private.pem.key",
				),
				mqtt.WithConnStateHandler(func(s mqtt.ConnState, err error) {
					fmt.Printf("%s: %v\n", s, err)
				}),
			},
		},
		mqtt.WithReconnectWait(500*time.Millisecond, 2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	if _, err := cli.Connect(ctx,
		thingName,
		mqtt.WithKeepAlive(30),
	); err != nil {
		panic(err)
	}

	t, err := tunnel.New(ctx, cli, map[string]tunnel.Dialer{
		"ssh": func() (io.ReadWriteCloser, error) {
			return net.Dial("tcp", "localhost:22")
		},
	})
	if err != nil {
		panic(err)
	}
	cli.Handle(t)

	select {}
}

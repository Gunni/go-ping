package main

import (
	"os"
	"fmt"
	"net"
	"flag"
	"time"
)

func main() {
	// Get the important arguments
	var port    string
	var count   uint64
	var delay   float64
	var timeout float64

	flag. StringVar(&port, "p", "80", "Port to ping")
	flag. Uint64Var(&count, "c", 5, "Number of connections to perform (0 = infinite)")
	flag.Float64Var(&delay, "d", 1, "Delay between sending each connection (in seconds)")
	flag.Float64Var(&timeout, "t", 5, "Timeout (in seconds)")

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Printf("Argument \"Host\" is required!\n")
		os.Exit(-1)
	}

	host := flag.Arg(0)
	if host == "" {
		fmt.Printf("Host \"%s\" is invalid\n", flag.Arg(0))
		os.Exit(-1)
	}
	// Arguments done

	var i uint64
	for i = 1; count == 0 || i <= count; i++ {
		// Show the iteration number
		fmt.Printf("%d;", i)

		timeoutVal := time.Duration(float64(time.Second) * timeout)

		// Send the ping
		duration, err := ping(host, port, timeoutVal)

		if err != nil {
			errn := err.(net.Error)
			if errn.Timeout() {
				fmt.Printf("timeout\n")
			} else {
				// Unknown error
				fmt.Printf("%s\n", err)
			}

			continue
		}

		// Print out the time taken, in seconds
		fmt.Printf("%f\n", duration.Seconds())

		// Delay if not last round
		if i != count {
			time.Sleep(time.Duration(float64(time.Second) * delay))
		}
	}
}

func ping(addr string, port string, timeout time.Duration) (duration time.Duration, err error) {
	// Start timer
	start := time.Now()

	// Connect to destination host:port
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(addr, port), timeout)
	if err != nil {
		return
	}
	fmt.Printf("%s;", conn.RemoteAddr().String())
	defer conn.Close()

	// Return time taken
	return time.Since(start), nil
}


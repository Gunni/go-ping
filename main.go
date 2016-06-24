package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"time"
)

func main() {
	// Get the important arguments
	var port string
	var count uint64
	var delay float64
	var timeout float64

	flag.StringVar(&port, "p", "80", "Port to ping")
	flag.Uint64Var(&count, "c", 5, "Number of connections to perform (0 = infinite)")
	flag.Float64Var(&delay, "d", 1, "Delay between sending each connection (in seconds)")
	flag.Float64Var(&timeout, "t", 5, "Timeout (in seconds)")

	flag.Parse()

	host := flag.Arg(0)

	if host == "" || flag.NArg() != 1 {
		fmt.Printf("The host argument is not optional\n", flag.Arg(0))
		os.Exit(-1)
	}
	// Arguments done

	// Prepare values for main process
	tcphost, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)

		os.Exit(1)
	}

	// Run the main process
	res := process(count, delay, timeout, tcphost)

	// Handle bad exit states
	if res != 0 {
		// nothing
	}

	os.Exit(res)
}

func process(count uint64, delay float64, timeout float64, tcphost *net.TCPAddr) int {
	var i uint64
	ret := 0

	for i = 1; count == 0 || i <= count; i++ {
		// Show the iteration number
		fmt.Printf("%d;", i)

		timeoutduration := time.Duration(float64(time.Second) * timeout)

		// Send the ping
		duration, err := ping(timeoutduration, tcphost)

		if err != nil {
			errn := err.(net.Error)
			if errn.Timeout() {
				// Timeouts are NaNs
				fmt.Printf("%f\n", math.NaN())
				ret = 1
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

	return ret
}

func ping(timeout time.Duration, tcphost *net.TCPAddr) (duration time.Duration, err error) {
	d := net.Dialer{Timeout: timeout, LocalAddr: nil, DualStack: false, KeepAlive: 0}

	// Start timer
	start := time.Now()

	// Print out the destination
	fmt.Printf("%s;", tcphost.String())

	// Connect to destination host:port
	conn, err := d.Dial("tcp", tcphost.String())
	if err != nil {
		return
	}
	defer conn.Close()

	// Return time taken
	return time.Since(start), nil
}

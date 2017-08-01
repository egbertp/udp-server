package main

import (
	"fmt"
	"net"
	"os"
	"flag"
	"log"
	"time"
	"github.com/egbertp/udp-server/helpers"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}


// Variables to identiy the build
var (
	CommitHash string
	VersionTag string
	BuildTime  string
)

func main() {

	var (
		port = flag.String("port", "9005", "UDP server port")
	)
	flag.Parse()


	// The port number (e.g. 9005) needs to be appended with an :
	parsedPort := ":" + *port

	log.Printf("The version is: %s; the commit hash is: %s. Build time is: %s", VersionTag, CommitHash, helpers.ParseBuildTime(BuildTime).Format(time.RFC1123))

	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", parsedPort)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	log.Printf("UDP sever running on port: %s", *port)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		log.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			log.Println("Error: ", err)
		}
	}
}

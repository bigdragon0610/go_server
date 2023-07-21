package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	path := "./"
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := os.Stat(absPath)
	if err != nil {
		log.Fatalf("Error retrieving the path %s: %v", absPath, err)
	} else if !fi.IsDir() {
		log.Fatalf("The path %s exists but is not a directory", absPath)
	}
	startServer(8080, absPath)
}

func startServer(startPort int, absPath string) {
	for i := startPort; i < startPort+100; i++ {
		serverAddress := fmt.Sprintf(":%d", i)
		listener, err := net.Listen("tcp", serverAddress)

		if err != nil {
			continue
		}

		fmt.Printf("Displaying %s at http://localhost:%d\n", absPath, i)

		log.Fatal(http.Serve(listener, http.FileServer(http.Dir(absPath))))
	}
	log.Fatalf("Could not start server on ports %d through %d", startPort, startPort+99)
}

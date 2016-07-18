// main.go

// written for the PROXY server Revision 1
package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"net/http"
	
)

// 1. Listen for connections.
// 2. Accept connections.
// 3. Read requests from the client.
// 4. Connect to the backend web server.
// 5. Send the request to the backend.
// 6. Read the response from the backend.
// 7. Send the response to the client, making sure to close it.

func main() {
	fmt.Println("Hello Start of the ProxServer V1.0 !")
	
	// 1. Listen to connections 
	if ln, err := net.Listen("tcp",":8080"); err == nil {
		
		//2 Accept Connections 
		
		for conn , err := ln.Accept(); err == nil{
			
			// Create a Buffered reader obj 
			
			reader := bufio.NewReader(conn)
			
			//3. Read requests from clients
			
			if req , err := http.ReadRequest(reader); err == nil{
				
				// 4.connect to the Backend web services and 
				
				if backend, err := net.Dial("tcp", "127.0.0.1:8081"); err == nil {
						backend_reader := bufio.NewReader(backend)
						// 5. Send the request to the backend.
						if err := req.Write(backend); err == nil {
							// 6. Read the response from the backend.
							if resp, err := http.ReadResponse(backend_reader, req); err == nil {
								// 7. Send the response to the client, making sure to close it.
								resp.Close = true
								if err := resp.Write(conn); err == nil {
									log.Printf("proxied %s: got %d", req.URL.Path, resp.StatusCode)
								}
								conn.Close()
								// Repeat back at 2: accept the next connection.
							}
						}
					}

			}
		}
		
	}
}

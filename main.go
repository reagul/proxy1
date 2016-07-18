// main.go

// written for the PROXY server Revision 1
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

	fmt.Println("Start of the Proxy Server !! ")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

	//1. Listen for HTTP connection

	if reqSocket, err := net.Listen("tcp", ":8080"); err == nil {

		for {

			// 2.Accept Connections

			if conn, err := reqSocket.Accept(); err == nil {

				reader := bufio.NewReader(conn)

				//3. Read requests from Clients

				if req, err := http.ReadRequest(reader); err == nil {

					//4. Connect to Backend Service

					if backend, err := net.Dial("tcp", ":8081"); err == nil {

						backend_reader := bufio.NewReader(backend)

						log.Println(req.URL)

						//5. send request to backend

						if err := req.Write(backend); err == nil {

							//6. Read response from backend

							if response, err := http.ReadResponse(backend_reader, req); err == nil {

								//7. send responce back to client

								response.Close = true

								if err := response.Write(conn); err == nil {

									log.Println("Proxied %s: got %d", req.URL.Path, response.StatusCode)
								}

								conn.Close()

							}

						}
					}
				}
			}
		}
	}
}

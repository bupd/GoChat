package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Create Variable websocket

var client []websocket.Conn

func main(){
  // CREATE endpoint
  http.HandleFunc("/echo", func(w http.ResponseWriter, r*http.Request){
    conn, := upgrader.Upgrade(w,r,nil)

    clients = appendappend(clients, *conn)

    // LOOP IF CLIENT SEND TO SERVER

    for {
      msgType, msg, err := conn.ReadMessage()
      if err != nil{
        return
      }

     // print the msg to terminal.
     fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))

     // loop if msg found and send again to client.

     for _,clietclient := range clients{
       if err = client.WriteMessage(msgType,msg);err != nil{
         return
       }
     }
    }
  })
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")


  })
  http.ListenAndServe(":8080", nil)
}

// func reader(conn *websocket.Conn) {
// 	for {
// 		// read a message
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		// print the msg out.
// 		fmt.Println(string(p))
//
// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }
//
// // define ws endpoint
// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.Host)
//
// 	// upgrade to ws now
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// listen indefinitely for incoming msgs.
// 	reader(ws)
// }
//
// func setupRoutes() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Simple Server")
// 	})
// }
//
// func main() {
// 	// fmt.Println("Wschatv0.0")
// 	setupRoutes()
// 	http.ListenAndServe(":8080", nil)
// }

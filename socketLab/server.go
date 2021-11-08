package socketLab

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func Set() {

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected")
		return nil
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.Emit("msg", "TEST")
		return msg
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

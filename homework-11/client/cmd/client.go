package main

const defaultProtocol = "tcp4"
const defaultSocket = "localhost:8000"

func main() {
	conn, err := net.Dial(defaultProtocol, defaultSocket)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

}
}

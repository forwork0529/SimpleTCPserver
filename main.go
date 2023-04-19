package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const addr = "0.0.0.0:5432"

// Протокол сетевой службы.
const proto = "tcp4"

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Вызов обработчика подключения.
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// Закрытие соединения.
	defer conn.Close()
	// Чтение сообщения от клиента.
	reader := bufio.NewReader(conn)
	for{
		b, err := reader.ReadString('.')
		if err != nil {
			log.Println(err)
			return
		}

		// Удаление символов конца строки.
		msg := strings.TrimSuffix(string(b), "\n")
		msg = strings.TrimSuffix(msg, "\r")
		fmt.Printf("Read meassage: %v\n", msg)
	}
}

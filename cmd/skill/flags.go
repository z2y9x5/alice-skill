package main

import (
	"flag"
)

// неэкспортируемая переменная flagRunAddr содержит адрес и порт для запуска сервера
var flagRunAddr string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", "localhost:80", "address and port to run server")
	flag.Parse()
}

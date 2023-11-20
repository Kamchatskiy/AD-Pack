package main

import (
	"C"
	_ "errors"
	_ "fmt"
	_ "net/http"
	"os"
	_ "time"
)

func UNUSED(x ...interface{}) {}

func main() {
	var host string = os.Args[1]
	var port int = 0
	UNUSED(host, port)

}

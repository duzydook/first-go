//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
)

func main() {
	// declaring the variable using the var keyword

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {

		fmt.Print(">> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			lineMessage(line)
		}

		select {
		case <-interrupt:
			log.Println("interrupt")
			return
		default:
			continue
		}
	}
}

func lineMessage(text string) {

	if text != "" {
		data := url.Values{
			"text": {text},
		}

		http.PostForm("http://localhost:8888/line", data)
	}

}

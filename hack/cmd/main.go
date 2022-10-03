package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	cli := http.Client{}
	ch := make(chan bool, 5)

	for {
		ch <- true
		go func() {
			r := rand.Intn(10)
			resp, _ := cli.Get("https://www.hostgator.com.br/")
			fmt.Printf("%v, %s\n", r, resp.Status)
			<-ch
		}()
	}
}

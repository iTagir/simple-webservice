package main

import (
	"log"
	"net/http"
	"time"

	timerandnum "github.com/iTagir/simple-webservice/time-rand-num"
)

func topUpQueue(nq *timerandnum.NumQueue) {

	for {
		nq.Add(time.Now(), timerandnum.GenerateRandom())
		time.Sleep(1 * time.Second)
	}
}

func main() {
	nq := timerandnum.NumQueue{Queue: []timerandnum.NumElem{}, Max: 10, Len: 0}

	go topUpQueue(&nq)

	dir := "C:\\Users\\Tagir\\Documents\\golangspace\\go_workspace\\src\\github.com\\iTagir\\simple-webservice"
	http.HandleFunc("/queue", nq.HTTPServerResponse)
	res := http.ListenAndServeTLS(":10443", dir+"\\cert.pem", dir+"\\key.pem", nil)
	log.Fatal(res)
}

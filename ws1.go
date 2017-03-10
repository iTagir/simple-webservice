package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	timerandnum "github.com/iTagir/simple-webservice/time-rand-num"
)

func topUpQueue(nq *timerandnum.NumQueue) {

	nnn := 123

	for {
		ttt := time.Now()
		nq.Add(ttt, nnn)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	nq := timerandnum.NumQueue{Queue: []timerandnum.NumElem{}, Max: 10, Len: 0}

	go topUpQueue(&nq)

	dir := "C:\\Users\\Tagir\\Documents\\golangspace\\go_workspace\\src\\github.com\\iTagir\\simple-webservice"
	http.HandleFunc("/j", nq.Handle)
	t := time.Now()
	fmt.Println("time type: ", reflect.TypeOf(t))
	fmt.Println("time: ", t)
	res := http.ListenAndServeTLS(":10443", dir+"\\cert.pem", dir+"\\key.pem", nil)
	log.Fatal(res)
	println("finished.", res)
}

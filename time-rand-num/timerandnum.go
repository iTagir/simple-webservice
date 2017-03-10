package timerandnum

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type NumElem struct {
	Date time.Time `json:date`
	Num  int       `json:num`
}

type NumQueue struct {
	Queue []NumElem
	Max   int
	Len   int
}

//Add element to queue
func (q *NumQueue) Add(t time.Time, n int) {
	newElem := NumElem{Date: t, Num: n}
	if q.Len < q.Max {
		tempQ := q.Queue
		q.Queue = append(tempQ, newElem)
		q.Len++
	} else {
		tempQ := q.Queue[1:]
		q.Queue = append(tempQ, newElem)
	}
}

func (q *NumQueue) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q.Queue)
}

//GenerateRandom number
func GenerateRandom() int {
	rand.Seed(100)
	return rand.Intn(10)
}
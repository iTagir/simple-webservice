package timerandnum_test

import (
	"testing"
	"time"

	timerandnum "github.com/iTagir/simple-webservice/time-rand-num"
)

func TestGenerateRandom(t *testing.T) {
	rn := timerandnum.GenerateRandom()
	t.Log("Random number:", rn)
	if !(rn >= 0 && rn <= 10) {
		t.Fatal("Failed to generate random number:", rn)
	}
}

func TestAdd(t *testing.T) {
	ttt := time.Now()
	nnn := timerandnum.GenerateRandom()
	nq := timerandnum.NumQueue{Queue: []timerandnum.NumElem{}, Max: 10, Len: 0}
	nq.Add(ttt, nnn)
	if nq.Len != 1 {
		t.Fatal("Failed Queue len 1: ", nq.Len)
	}
	nq.Add(ttt, nnn)
	nq.Add(ttt, nnn)
	nq.Add(ttt, nnn)
	nq.Add(ttt, nnn)
	if nq.Len != 5 {
		t.Fatal("Failed Queue len 5: ", nq.Len)
	}
	nq.Add(ttt, nnn)
	nq.Add(ttt, nnn)
	nq.Add(ttt, nnn)
	nq.Add(ttt, nnn)
	nq.Add(ttt, nnn)
	nq.Add(ttt, nnn) //eleventh element in the queue
	if nq.Len != 10 {
		t.Fatal("Failed Queue len 10: ", nq.Len)
	}
}

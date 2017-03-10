package timerandnum_test

import (
	"testing"
	"time"

	timerandnum "github.com/iTagir/simple-webservice/time-rand-num"
)

func TestGenerateRandom(t *testing.T) {
	i := timerandnum.GenerateRandom()
	t.Log("returned something:", i)
}

func TestAdd(t *testing.T) {
	ttt := time.Now()
	nnn := 123
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
	nq.Add(ttt, nnn)
	if nq.Len != 10 {
		t.Fatal("Failed Queue len 10: ", nq.Len)
	}
}

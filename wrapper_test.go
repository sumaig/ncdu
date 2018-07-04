package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestWalker(t *testing.T) {
	a := `[1,1,{"progname":"ncdu","progver":"1.13","timestamp":1530676862},
[{"name":"/home/gaoguangpeng/test","asize":4096,"dsize":4096,"dev":2049,"ino":268784},
{"name":"tb","asize":14,"dsize":4096,"ino":268947},
[{"name":"l2","asize":4096,"dsize":4096,"ino":268949},
{"name":"la","asize":13,"dsize":4096,"ino":268950},
{"name":"lc","asize":24,"dsize":4096,"ino":268952},
[{"name":"l3","asize":4096,"dsize":4096,"ino":268953},
{"name":"d","asize":10,"dsize":4096,"ino":268954}],
{"name":"lb","asize":17,"dsize":4096,"ino":268951}],
{"name":"tc","asize":12,"dsize":4096,"ino":268948},
{"name":"ta","asize":9,"dsize":4096,"ino":268946}]]`

	err := json.Unmarshal([]byte(a), &message)
	if err != nil {
		t.Error(err)
	}

	walker(message[3])
}

func TestPathJoin(t *testing.T) {
	p1 := "/Users/gaoguangpeng/test"
	p2 := "tb"

	pj := pathJoin()
	fmt.Println(pj(p1))
	fmt.Println(pj(p2))
}

func TestQuickSort(t *testing.T) {
	a := []uint64{10, 13, 14, 17, 24, 12, 9}
	QuickSort(a)
	fmt.Println(a)
}

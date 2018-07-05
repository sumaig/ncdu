package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestWalker(t *testing.T) {
	a := `[1,1,{"progname":"ncdu","progver":"1.13","timestamp":1530704292},
[{"name":"/Users/gaoguangpeng/test","asize":4096,"dsize":4096,"dev":2049,"ino":268784},
{"name":"tb","asize":14,"dsize":4096,"ino":268947},
[{"name":"l2","asize":4096,"dsize":4096,"ino":268949},
{"name":"la","asize":13,"dsize":4096,"ino":268950},
{"name":"lc","asize":24,"dsize":4096,"ino":268952},
[{"name":"l3","asize":4096,"dsize":4096,"ino":268953},
{"name":"d","asize":10,"dsize":4096,"ino":268954}],
{"name":"lb","asize":17,"dsize":4096,"ino":268951}],
{"name":"tc","asize":12,"dsize":4096,"ino":268948},
[{"name":"l21","asize":4096,"dsize":4096,"ino":268956},
{"name":"oiiasdfii","asize":8,"dsize":4096,"ino":268957},
[{"name":"l3","asize":4096,"dsize":4096,"ino":268959},
{"name":"sjdflkj","asize":6,"dsize":4096,"ino":268962},
{"name":"2df","ino":268960},
{"name":"sdfwkjhf","ino":268961},
{"name":"sjdsd","asize":6,"dsize":4096,"ino":268963}],
{"name":"oiiii","asize":8,"dsize":4096,"ino":268945},
{"name":"oiiasafii","asize":8,"dsize":4096,"ino":268958}],
{"name":"ta","asize":9,"dsize":4096,"ino":268946}]]`

	err := json.Unmarshal([]byte(a), &message)
	if err != nil {
		t.Error(err)
	}

	walker(message[3], baseDir)
}

func TestQuickSort(t *testing.T) {
	a := []uint64{10, 13, 14, 17, 24, 12, 9}
	QuickSort(a)
	fmt.Println(a)
}

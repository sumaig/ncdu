package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"
)

var message = make([]interface{}, 4)

type meta struct {
	Name  string `json:"name"`
	Asize uint64 `json:"asize"`
	Ino   uint64 `json:"ino"`
}

type Result struct {
	msg map[uint64]*meta
	top []uint64
}

func (r *Result) OutPut() {
	QuickSort(result.top)
	count := len(result.top)
	limit := 10
	if count < limit {
		limit = count
	}

	message := fmt.Sprintf("\t%-20s%-20s%-20s\r\n", "path", "size", "inode")
	for i := 0; i < limit; i++ {
		message += fmt.Sprintf("\t%-20s%-20s%-20s\r\n", r.msg[r.top[i]].Name, r.msg[r.top[i]].Asize, r.msg[r.top[i]].Ino)
	}

	fmt.Println(message)
}

func newResult() *Result {
	return &Result{
		msg: make(map[uint64]*meta),
		top: make([]uint64, 0),
	}
}

var result = newResult()

// ncdu wrapper
func Wrapper(cmd string, timeout int) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command := exec.Command("/bin/bash", "-c", cmd)
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	command.Stdout = &stdout
	command.Stderr = &stderr
	command.Start()

	err, isTimeout := CmdRunWithTimeout(command, time.Duration(timeout)*time.Second)

	errStr := stderr.String()
	if errStr != "" {
		fmt.Println(cmd, "failed: ", strings.TrimSpace(errStr))
		return
	}

	if isTimeout {
		if err == nil {
			fmt.Println("timeout and kill process", cmd, "successfully")
		}

		if err != nil {
			fmt.Println("kill process", cmd, "occur error: ", err)
		}
		return
	}

	if err != nil {
		fmt.Println(cmd, "failed: ", err)
	}

	err = json.Unmarshal(stdout.Bytes(), &message)
	if err != nil {
		fmt.Println("json unmarshal error: ", err)
	}

	fmt.Println(message[3])
	walker(message[3])
}

// run command with timeout
func CmdRunWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	var err error
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout):
		log.Printf("timeout, process:%s will be killed", cmd.Path)

		go func() {
			<-done // allow goroutine to exit
		}()

		err = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		if err != nil {
			log.Println("kill failed, error:", err)
		}

		return err, true
	case err = <-done:
		return err, false
	}
}

var basePath string

// walk json
func walker(i interface{}) {
	copyPath := basePath
	itemA := pathJoin()
	itemB := pathJoin()
	itemA(basePath)
	itemB(copyPath)
	switch i := i.(type) {
	case []interface{}:
		for t, v := range i {
			x, ok := v.(map[string]interface{})
			if ok {
				p := x["name"].(string)
				a := x["asize"].(float64)
				in := x["ino"].(float64)
				// TODO: Use file stat
				if a == 4096 {
					if t == 0 {
						basePath = itemA(p)
						// fmt.Println("base", basePath)
					}
					copyPath = itemB(p)
					// fmt.Println("copy", copyPath)
				} else {
					p = path.Join(copyPath, p)
					fmt.Println(p)
					result.msg[uint64(a)] = &meta{
						Name:  p,
						Asize: uint64(a),
						Ino:   uint64(in),
					}
					result.top = append(result.top, uint64(a))
				}

			} else {
				walker(v)
			}
		}
		fmt.Println(result.top)
		QuickSort(result.top)
		fmt.Println(result.top)
	default:
		// fmt.Println("default")
	}
}

// join path
func pathJoin() func(p string) string {
	var base string
	return func(p string) string {
		if strings.HasPrefix(p, "/") {
			base = p
		} else {
			base = path.Join(base, p)
		}
		return base
	}

}

func QuickSort(values []uint64) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		// fmt.Println(values)
		if values[i] < mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	QuickSort(values[:head])
	QuickSort(values[head+1:])
}

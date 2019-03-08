package main

import (
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	outC := make(chan string)
	defer close(outC)
	go func() {
		var buf strings.Builder
		io.Copy(&buf, r)
		r.Close()
		outC <- buf.String()
	}()

	f()

	os.Stdout = stdout
	w.Close()

	return <-outC
}

func Test_main(t *testing.T) {
	tests := []struct {
		arg string
		ans string
	}{
		{"1,2", "3"},
		{"3,3", "6"},
		{"3,3", "0"}, // This case will return FAIL.
		// TODO: Add test cases.
	}
	argsbuf := os.Args
	for i, tt := range tests {
		si := strconv.Itoa(i)
		t.Run("Case "+si, func(t *testing.T) {
			os.Args = argsbuf
			os.Args = append(os.Args, tt.arg)
			ret := captureStdout(main)
			if ret != tt.ans {
				t.Errorf("Unexpected output: %s Need: %s", ret, tt.ans)
			}
		})
	}
}

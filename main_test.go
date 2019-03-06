package main

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"testing"
)

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func Test_main(t *testing.T) {
	tests := []struct {
		arg string
		ans string
	}{
		{"1,2", "3"},
		{"3,3", "6"},
		{"3,3", "0"},	// This case will return FAIL.
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


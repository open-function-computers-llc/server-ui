package server

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/gorilla/websocket"
)

func runScript(action string, domain string, ws *websocket.Conn) {
	scripts := map[string]string{
		"unlock": "ofco-unlock-site.sh",
		"lock":   "ofco-lock-site-production-modx.sh",
	}
	script, ok := scripts[action]

	// prep our output message
	output := map[string]interface{}{
		"message": "",
	}
	if !ok {
		output["message"] = "Unknow action " + action
		ws.WriteJSON(output)
		return
	}

	scriptPath := os.Getenv("SCRIPTS_ROOT") + "/" + script
	output["message"] = "About to run `" + scriptPath + "`"
	ws.WriteJSON(output)

	// run the script
	output["message"] = string("Script output:")
	ws.WriteJSON(output)

	// do it again but line by line
	cmd := exec.Command(scriptPath, domain)

	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := NewCapturingPassThroughWriter(os.Stdout, ws)
	stderr := NewCapturingPassThroughWriter(os.Stderr, ws)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr\n")
	}
}

// below is some modified IO code that came from
// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
type CapturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
	ws  *websocket.Conn
}

// NewCapturingPassThroughWriter creates new CapturingPassThroughWriter
func NewCapturingPassThroughWriter(w io.Writer, ws *websocket.Conn) *CapturingPassThroughWriter {
	return &CapturingPassThroughWriter{
		w:  w,
		ws: ws,
	}
}

func (w *CapturingPassThroughWriter) Write(d []byte) (int, error) {
	w.buf.Write(d)

	output := map[string]interface{}{
		"message": string(d),
	}
	w.ws.WriteJSON(output)
	return w.w.Write(d)
}

// Bytes returns bytes written to the writer
func (w *CapturingPassThroughWriter) Bytes() []byte {
	return w.buf.Bytes()
}

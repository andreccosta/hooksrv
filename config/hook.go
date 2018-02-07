package config

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

// Hook represents a hook
type Hook struct {
	Headers  map[string]string `json:"headers"`
	Commands []string          `json:"commands"`
}

// Handler handles an incoming request to a hook
func (hook *Hook) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Invalid HTTP method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for header, value := range hook.Headers {
		if r.Header.Get(header) != value {
			log.Print("Header mismatch")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	log.Print("Received valid request")
	log.Print("Executing commands ...")

	var out bytes.Buffer

	for _, command := range hook.Commands {
		log.Printf("Executing\t%s", command)
		err := execCommand(&out, command)

		if err != nil {
			log.Printf("Command failed\t%s", err)
			return
		}
	}

	log.Printf("Output\t%s", out.String())
	fmt.Fprint(w, out.String())
}

func execCommand(out *bytes.Buffer, command string) error {
	var err error
	cmdParts := strings.Split(command, " ")

	log.Print("Executable\t" + cmdParts[0])
	log.Print("Params\t" + strings.Join(cmdParts[1:], ","))

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	cmd.Stdout = out
	cmd.Stderr = out
	err = cmd.Run()

	return err
}

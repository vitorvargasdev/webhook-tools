package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/vitorvargasdev/webhook-tools/pkgs/printlog"
)

type output struct {
	header string
	body   string
}

func StartServer(flags Flags) error {
	port := strconv.Itoa(flags.port)

	http.HandleFunc("/listen", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(flags.httpCode)
		if _, err := w.Write([]byte(`{"message": "received"}`)); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		output := printOutput(r)

		printLog.Print(
			output.header,
			output.body,
		)
	})

	fmt.Printf("Starting HTTP server on %s:%s%s\n", "http://localhost", port, "/listen")
	return http.ListenAndServe(":"+port, nil)
}

func printOutput(r *http.Request) output {
	output := output{}

	headers := strings.Builder{}

	for k, v := range r.Header {
		headers.WriteString(k + ": " + v[0] + "\n")
	}

	body, _ := io.ReadAll(r.Body)
	var prettyBody bytes.Buffer
	if err := json.Indent(&prettyBody, body, "", "  "); err != nil {
		panic(err)
	}

	output.header = headers.String()
	output.body = prettyBody.String()

	return output
}

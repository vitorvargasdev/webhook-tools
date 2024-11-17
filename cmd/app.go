package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func StartServer(flags Flags) error {
	port := strconv.Itoa(flags.port)

	http.HandleFunc("/listen", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(flags.httpCode)
		if _, err := w.Write([]byte(`{"message": "received"}`)); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})

	fmt.Printf("Starting HTTP server on port %s\n", port)
	return http.ListenAndServe(":"+port, nil)
}

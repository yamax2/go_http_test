package main

import (
    "net/http"
    "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server.\n"))
}

func main() {
    http.HandleFunc("/", HelloServer)

    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

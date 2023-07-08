package main

import (
	"log"

	"github.com/Aner-Git/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
	/*
		resp := server.ProduceResponse{Offset: uint64(333)}
		json.NewEncoder(os.Stdout).Encode(resp)

		cresp := server.ConsumeResponse{Record: server.Record{}}
		json.NewEncoder(os.Stdout).Encode(cresp)
	*/

}

// This repository serves as the basis for a new project. The project provides
// interaction with a graph-based management structure.
//
// There are up-to ten TODO items to complete.
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/opentickettech/graph-exercise/env"
)

var (
	booted       = time.Now()
	jsonContents = `[
	{"manager": "mart", "report": "peter"},
	{"manager": "joost", "report": "chris"},
	{"manager": "dennis", "report": "peter"},
	{"manager": "chris", "report": "mart"},
	{"manager": "chris", "report": "peter"},
	{"manager": "dennis", "report": "marta"},
	{"manager": "joost", "report": "peter"},
	{"manager": "joost", "report": "jos"},
	{"manager": "jos", "report": "dennis"}
]`
)

func main() {
	fmt.Println("booting")
	http.HandleFunc("/", rootRoute)
	http.HandleFunc("/store", storeRoute)
	http.HandleFunc("/retrieve", retrieveRoute)
	http.HandleFunc("/retrieve/ordered", orderedRoute)
	http.HandleFunc("/stack", stackRoute)

	fmt.Println("listening")

	// TODO ensure LISTEN_PORT can be set from the environment
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", env.PortFb("LISTEN_PORT", 8080)), http.DefaultServeMux)
	must(err)

	// TODO implement graceful shutdown
}

// rootRoute is a http.Handler simply writing the time the server was booted to
// the response.
func rootRoute(writer http.ResponseWriter, _ *http.Request) {
	// TODO fix printing of the date
	_, err := fmt.Fprintf(writer, "up since %d", booted)
	must(err)
}

func storeRoute(writer http.ResponseWriter, request *http.Request) {
	// TODO
	//  1) validate the posted contents
	//  2) marshal the posted contents into graph.Node's
	//  3) store the entire graph in `/tmp/important_database`.
	_, err := fmt.Fprintf(writer, "storeRoute")
	must(err)

	_, err = io.Copy(writer, request.Body)
	must(err)
}

func retrieveRoute(writer http.ResponseWriter, request *http.Request) {
	// TODO
	//  1) Retrieve the graph stored in `/tmp/important_database`
	//  2) return the entire graph in the response.
	_, err := fmt.Fprintf(writer, "retrieveRoute")
	must(err)
}

func orderedRoute(writer http.ResponseWriter, request *http.Request) {
	// TODO
	//  1) Retrieve the graph stored in `/tmp/important_database`
	//  2) Return the results of Graph.Ordered(rootNode) where rootNode
	//     is a pointer to the root of the graph.
	_, err := fmt.Fprintf(writer, "orderedRoute")
	must(err)
}

func stackRoute(writer http.ResponseWriter, request *http.Request) {
	// TODO write a single string: 'up', or 'down' representing the direction to
	//  which the stack grows. Programmatically determine which direction the stack
	//  grows.
	_, err := fmt.Fprintf(writer, "storeRoute")
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

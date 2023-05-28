package main

import (
	//"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	CORSmux := middlewareCORS(mux)
	svr := http.Server{
		Handler: CORSmux,
		Addr: "localhost:8080",
	}
	svr.ListenAndServe()
}

func middlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*
serves a file called index.html
[x] Add the HTML code above to a file called index.html in the same root directory as your server
[x] Use the http.NewServeMux's .Handle() method to add a handler for the root path (/).
[ ] Use a standard http.FileServer as the handler
[ ] Use http.Dir to convert a filepath, (in our case a dot: . which indicates the current directory) to a directory for the http.FileServer.

[ ] Re-build and run your server
[ ] Test your server by visiting http://localhost:8080 in your browser
[ ] Run the tests in the window on the right

----------------------------------------------------
run a server that binds to localhost:8080 and always responds with a 404 Not Found response.
[x] Create a new http.ServeMux
[x] Wrap that mux in a custom middleware function that adds CORS headers to the response
	(see the tip below on how to do that).
[x] Create a new http.Server and use the corsMux as the handler
[x] Use the server's ListenAndServe method to start the server
[x] Build and run your server (e.g. go build -o out && ./out)
[x] Open http://localhost:8080 in your browser. You should see a 404 error because we haven't connected any handler logic yet. Don't worry, that's what is expected for the tests to pass for now.
[x] Run the tests in the window on the right.
you'll need to add a few headers to all of your responses:
Access-Control-Allow-Origin: *
You'll also need to handle requests with the HTTP OPTIONS method
and simply return a 200 OK response with the appropriate headers.
*/
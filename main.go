package main

import (
	//"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
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
run a server that binds to localhost:8080 and always responds with a 404 Not Found response.
[x] Create a new http.ServeMux
[x] Wrap that mux in a custom middleware function that adds CORS headers to the response
	(see the tip below on how to do that).
[x] Create a new http.Server and use the corsMux as the handler
[x] Use the server's ListenAndServe method to start the server
[ ] Build and run your server (e.g. go build -o out && ./out)
[ ] Open http://localhost:8080 in your browser. You should see a 404 error because we haven't connected any handler logic yet. Don't worry, that's what is expected for the tests to pass for now.
[ ] Run the tests in the window on the right.

You might have a browser setting that blocks JS/WASM requests to localhost.
If so, you'll need to disable it or use a different browser.
Known issues exist with:
--Brave browser's "shield"
--Google Chrome > settings > advanced > privacy and security > site settings >
	insecure content > ensure "block insecure private network requests" is disabled

or the Boot.dev test suite to be able to make requests to your fileserver,
you'll need to add a few headers to all of your responses:
Access-Control-Allow-Origin: *
You'll also need to handle requests with the HTTP OPTIONS method
and simply return a 200 OK response with the appropriate headers.
*/
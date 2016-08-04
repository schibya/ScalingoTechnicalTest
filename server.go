package main

import ("fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html>\n" +
		"<html>\n" +
		"  <head>\n" +
		"    <title>Technical test</title>\n" +
		"  </head>\n" +
		"  <body>\n" +
		"    <form method='get' action='search'>\n" +
		"      <input type='search' name='q' id='q' autofocus />\n" +
		"      <input type='submit' value='search' />\n" +
		"    </form>\n" +
		"  </body>\n" +
		"</html>")
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "<!DOCTYPE html>\n" +
		"<html>\n" +
		"  <head>\n" +
		"    <title>Search results</title>\n" +
		"  </head>\n" +
		"  <body>\n" +
		"    " + r.Form.Get("q") + "\n" +
		"  </body>\n" +
		"</html>")
}

package html

/*
 * Handle index.html action
 */

import (
	"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	"strings"
)

// HandleIndex is process functions of index.html
// TODO:
// Cookie
// Config
func HandleIndex(w http.ResponseWriter, req *http.Request) {
	//http.ServeFile(w, req, "html/")
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("path: ", req.URL.Path)
	fmt.Println("scheme: ", req.URL.Scheme)
	fmt.Println("method: ", req.Method)

	// First time access server
	// GET / URL == '/' for index page
	if req.Method == "GET" {
		fmt.Println("handle method get")
		//indexHTML, _ := ioutil.ReadFile("html/index.html")
		//fmt.Fprintf(w, string(indexHTML))
		http.ServeFile(w, req, "html/index.html")
	}

	// Form will require by POST
	for i, value := range req.Form {
		fmt.Println("key: ", i)
		fmt.Println("value: ", strings.Join(value, ""))
	}
	if req.Method == "POST" {
		// FIXME: potential fail
		fmt.Println(req.Form["password"])
		http.ServeFile(w, req, "html/shell.html")
	}
}

// TODO: base64 decode

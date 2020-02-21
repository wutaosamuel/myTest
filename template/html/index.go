package html

/*
 * Handle index.html action
 */

import (
	"encoding/base64"
	"fmt"
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
	PrintHTMLInfo(req)

	// First time access server
	// GET for responsing index.html
	if req.Method == "GET" {
		fmt.Println("handle index method get")
		http.ServeFile(w, req, "html/index.html")
	}

	// Form will require by POST
	if req.Method == "POST" {
		password, _ := base64.StdEncoding.DecodeString(FormToString(req, "password"))
		fmt.Println(FormToString(req, "password"))
		fmt.Println(string(password))
		// TODO:
		// Read Config and allow to sign in
		// Cookie for Auth
		http.ServeFile(w, req, "html/shell.html")
	}
}

// PrintHTMLInfo infomation
func PrintHTMLInfo(req *http.Request) {
	fmt.Println(req.Form)
	fmt.Println("path: ", req.URL.Path)
	fmt.Println("scheme: ", req.URL.Scheme)
	fmt.Println("method: ", req.Method)
}

// FormToString to string
func FormToString(req *http.Request, attribute string) string {
	return strings.Join(req.Form[attribute], "")
}

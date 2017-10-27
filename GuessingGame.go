// Web Application - Guessing Game
// Hugh Brady - G00338260 - 23/10/17

package main 

import "net/http"

func main() { 
	// handles request for index.html by serving files in "static" folder
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// serves on local host
    http.ListenAndServe(":8080", nil)
}
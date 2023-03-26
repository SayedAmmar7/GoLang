package handlerfunc

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method == "POST" {
		// Parse the form data from the request body
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// Get the form data from the request
		name := r.Form.Get("name")
		email := r.Form.Get("email")

		// Do something with the form data
		fmt.Fprintf(w, "Hello, %s! Your email address is %s", name, email)
	} else {
		// If the request method is not POST, display the form
		fmt.Fprint(w, "<form method='POST'><input type='text' name='name'><input type='text' name='email'><input type='submit' value='Submit'></form>")
	}
}

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	// Handle request
	r := mux.NewRouter()

	r.HandleFunc("/cal/{num1}/{operation}/{num2}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		num1Str := vars["num1"]
		operation := vars["operation"]
		num2Str := vars["num2"]

		// Convert num1 and num2 to integers
		num1, err := strconv.Atoi(num1Str)
		if err != nil {
			http.Error(w, "Invalid number: "+num1Str, http.StatusBadRequest)
			return
		}

		num2, err := strconv.Atoi(num2Str)
		if err != nil {
			http.Error(w, "Invalid number: "+num2Str, http.StatusBadRequest)
			return
		}

		// Perform the requested operation
		result := calculate(num1, operation, num2)

		// Create response string
		response := fmt.Sprintf(`
    <html>
        <head>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f4f4f4;
                    text-align: center;
                    padding: 50px;
                }
                h1 {
                    color: #333;
                }
                p {
                    font-size: 18px;
                    color: #666;
                }
            </style>
        </head>
        <body>
            <h1>Result</h1>
            <p>The result of the operation is: <strong>%d</strong></p>
        </body>
    </html>
`, result)

		// Write the response to the client
		w.Write([]byte(response))
	})

	// "/" is the route path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!") // Write data to response
	})

	// "/about" is the route path
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HaHa") // Write data to response
	})

	// Listen and serve
	http.ListenAndServe(":8080", r)
}

// Function to calculate the result based on the operation
func calculate(num1 int, operation string, num2 int) int {
	switch strings.ToLower(operation) {
	case "plus":
		return num1 + num2
	case "minus":
		return num1 - num2
	case "multiply":
		return num1 * num2
	case "divide":
		if num2 != 0 {
			return num1 / num2
		}
		return 0
	default:
		return 0
	}
}

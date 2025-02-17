// A simple Wikipedia to open files and view their content

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	// format filename
	// Call the os.Write function to append files
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) // My guess is this would return an error type
}

func loadPage(title string) (*Page, error) {
	// format filename
	// call the ReadFile function
	// Hnadle errors, like file not found
	// Return an instance of the struct
	// This is basically a constructor for struct Page
	filename := title + ".txt"           // prepare the file name
	body, error := os.ReadFile(filename) // read the files in filename and put it in body
	page := &Page{
		Title: title,
		Body:  body,
	}
	if error != nil {
		return page, error
	}
	return page, nil

	// When a part of a return is not needed, we can substitute it with nil
}

func input() (string, bool) {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			fmt.Println("Input Terminated")
			return "", false
		}
	} else {
		fmt.Println("Input Successful!")
	}
	return scanner.Text(), true
}

// Responsible for handling http requests
// Derive title from request
// Use the title to derive the file from our source directory
// Do some basic formatting with the content of the file

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)

	println(string(p.Body))
	fmt.Fprintf(w, "<h1>%s<h1><div>%s<div>", p.Title, p.Body)
}

func main() {
	// Define a new scanner method
	// Collect users input
	// Call the loadPage method
	// Handle all errors appropriately
	// Return the content of the file

	// Program prompt

	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	/*fmt.Print("Filename: ")

	// input() takes in inputs and handle empty input errors
	filenameinput, _ := input()

	page, error := loadPage(filenameinput)

	// Handle error (if filenameinput) does not exists
	if error != nil {
		fmt.Println("Error Opening File: ", error)
		fmt.Println("Would you like to create a new file? (Leave blank for no): ")

		// Let's see if we can create a new file for the user
		_, status := input()
		if status {
			file, err := os.Create(filenameinput)
			if err != nil {
				// Update log for errors opening file
				log.Fatal(err)
			}

			// Close file after main has exited successfully
			defer file.Close()
		} else {
			return // Close program if no input was detected
		}
		// page.Save()
	}

	// Output the content of the file
	body := string(page.Body) // since the program return nil if an error occurs we have to exit the program.
	fmt.Println(body)*/
}

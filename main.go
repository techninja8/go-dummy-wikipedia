// A simple Wikipedia to open files and view their content

package main

import (
	"fmt"
	"html/template"
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
	filename := title + ".txt"         // prepare the file name
	body, err := os.ReadFile(filename) // read the files in filename and put it in body
	page := &Page{
		Title: title,
		Body:  body,
	}
	if err != nil {
		return nil, err
	}
	return page, nil

	// When a part of a return is not needed, we can substitute it with nil
}

/*func input() (string, bool) {
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
}*/

// Responsible for handling http requests
// Derive title from request
// Use the title to derive the file from our source directory
// Do some basic formatting with the content of the file

// This function is reponsible for redering our HTML templates
// This function would also handle any errors that arise when we try to do this
func renderTemplate(w http.ResponseWriter, templ string, p *Page) {
	t, err := template.ParseFiles(templ + ".html")
	// If template,ParseFiles fail
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If t.Execute fails, let's handle the error accordingly
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// This function is responsible for /view/ endpoints
// This function handles errors related to loading a page
// If a certain file does not exist, then redirect to edit and save.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	// The requested page does not exists
	if err != nil {
		fmt.Printf("Editing new file %s...", title)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	// Render the HTML file for reponse
	renderTemplate(w, "view", p)
}

// This function is reposnible for saving files
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}

	// Call the Save() method and save file to root directory
	p.Save()

	// Handle errors if we cannot save a file
	err := p.Save()
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	// Once the file has been saved, go ahead to redirect to view the saved document
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// This function is responsible for editing new files
// This function also handles error related to opening the file for editing
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	// If page struct does not exists in this context
	// Go ahead and create it
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	// Bring all the handlers to scope
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/edit/", editHandler)

	// Listen for requests on port :8080
	log.Fatal(http.ListenAndServe(":8080", nil))

	// This is a previous command-line implementation
	// I removed some functions in future updates, pls refer to revision history to view them
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

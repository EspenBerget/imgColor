package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func getNames(fs []os.FileInfo) []string {
	names := make([]string, len(fs))
	for i, f := range fs {
		names[i] = f.Name()
	}
	return names[:]
}

type options struct {
	Images []string
	// TODO
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	static, err := os.Open("static")
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	fileInfo, err := static.Readdir(0)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	images := getNames(fileInfo)
	o := options{Images: images}
	t := template.Must(template.ParseFiles("options.html"))
	t.Execute(w, o)
}

func main() {
	http.HandleFunc("/", indexHandler)
	//	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":4545", nil))
}

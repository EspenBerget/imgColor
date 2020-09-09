package main

import (
	"html/template"
	"imgcolor/img"
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
	Active string
}

var o options

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		o.Active = r.FormValue("image")
	}
	if o.Active != "" {
		img.Hist(o.Active)
	}
	t := template.Must(template.ParseFiles("options.html"))
	t.Execute(w, o)
}

func main() {
	static, err := os.Open("static")
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := static.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	images := getNames(fileInfo)
	o = options{Images: images}

	http.HandleFunc("/", indexHandler)
	http.Handle("/bin/", http.StripPrefix("/bin", http.FileServer(http.Dir("bin"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":4545", nil))
}

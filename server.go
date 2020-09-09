package main

import (
	"html/template"
	"imgcolor/img"
	"io/ioutil"
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

type upload struct {
	Upload bool
}

var u upload

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20) // 10MB filecap
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		tempFile, err := os.Create("./static/tmp/upload.jpg")
		if err != nil {
			log.Panic(err)
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Panic(err)
		}
		tempFile.Write(fileBytes)
		u.Upload = true
	}
	if u.Upload {
		img.Hist("/tmp/upload.jpg")
	}

	t := template.Must(template.ParseFiles("upload.html"))
	t.Execute(w, u)
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
	u = upload{Upload: false}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/bin/", http.StripPrefix("/bin", http.FileServer(http.Dir("bin"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":4545", nil))
}

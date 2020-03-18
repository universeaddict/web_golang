package main

import (
    // "os"
	"fmt"
	"net/http"
	"html/template"
)

type M map[string]interface{}

func main() {
    // port := os.Getenv("PORT")
    http.Handle("/static/", 
    	http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

    http.HandleFunc("/", home)
    http.HandleFunc("/index", home)
    
    fmt.Println("server started at localhost:8080")
    http.ListenAndServe(":8080", nil)
    // fmt.Println("server started at localhost:"+port)
    // http.ListenAndServe(":"+port, nil)
}

func home(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/index" && r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound, r.URL.Path)
        return
    }
    var tmpl = template.Must(template.ParseFiles(
        "views/index.html",
        "views/_header.html",
        "views/_footer.html",
    ))
	var data = M{"about": "Saya Leonardus Fernando Purba, umur masih muda XD, pernah kuliah (lulus walaupun telat), seorang web developer yang berkerja di salah satu perusahaan IT di Pontianak, bahasa pemrograman yang sering digunakan yaitu PHP, JavaScript, Go, dan MySql. Website ini adalah salah satu yang dikembangkan dengan Go."}
    err := tmpl.ExecuteTemplate(w, "index", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, url_path string) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "404 page not found bro. ", url_path)
    }
}
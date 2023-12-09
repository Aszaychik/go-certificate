package main

import (
	"go-certificate/models"
	"go-certificate/pkg"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	certificate := models.Certificate{
    ID: uuid.NewString(),
    Name: "AsZaychik",
    CourseName: "Go Backend",
    Date: time.Now(),
  }

	htmlTemplate := `<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<title>SERTIFIKAT KOPETENSI KELULUSAN</title>
		</head>
		<body>
			<h1>No.{{.ID}}</h1>
			<h2>Nama : {{.Name}}</h2>
			<h2>Course : {{.CourseName}}</h2>
			<h3>Date : {{.Date}}</h3>
		</body>
	</html>
	`

	// render the HTML template on the index route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err := template.Must(template.New("certificate").
			Parse(htmlTemplate)).
			Execute(w, certificate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// serve the PDF on the download route
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		pdf, err := pkg.GenerateHtmlToPDF(htmlTemplate, certificate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=certificate.pdf")
		w.Write(pdf)
	})
	println("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
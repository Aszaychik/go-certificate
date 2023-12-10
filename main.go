package main

import (
	"go-certificate/models"
	"go-certificate/pkg"
	certificateTemplate "go-certificate/templates/certificate"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	certificate := models.Certificate{
		ID: uuid.NewString(),
    Name: "AsZaychik",
    CourseName: "Go Backend",
    CreatedAt: time.Now().UTC(),
  }
	
	tookDate := strings.Split(certificate.CreatedAt.String(), " ")
	certificate.IssuedAt = tookDate[0]
	
	embedTemplate := certificateTemplate.CertificateBlankContent

	// render the HTML template on the index route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		certificateTemplate, err := template.ParseFS(embedTemplate, "certificate.html")
		err = certificateTemplate.Execute(w, certificate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// serve the PDF on the download route
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		pdf, err := pkg.GenerateHtmlToPDF(embedTemplate, certificate)
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
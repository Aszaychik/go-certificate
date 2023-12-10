package pkg

import (
	"bytes"
	"embed"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GenerateHtmlToPDF(htmlTemplate embed.FS, model interface{}) ([]byte, error) {
	var htmlBuffer bytes.Buffer
	certificateTemplate, err := template.ParseFS(htmlTemplate, "certificate.html")
	err = certificateTemplate.Execute(&htmlBuffer, model)
	if err != nil {
		return nil, err
	}

	// Initialize the converter
	pdfGen, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlBuffer.Bytes()))
	pdfGen.AddPage(page)
	pdfGen.MarginBottom.Set(0)
	pdfGen.MarginLeft.Set(0)
	pdfGen.MarginRight.Set(0)
	pdfGen.MarginTop.Set(0)
	pdfGen.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfGen.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// Generate the PDF
	err = pdfGen.Create()
	if err != nil {
		return nil, err
	}
	return pdfGen.Bytes(), nil
}
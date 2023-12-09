package pkg

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GenerateHtmlToPDF(htmlTemplate string, model interface{}) ([]byte, error) {
	var htmlBuffer bytes.Buffer
	err := template.Must(template.New("template").
		Parse(htmlTemplate)).
		Execute(&htmlBuffer, model)
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
	pdfGen.MarginLeft.Set(10)
	pdfGen.MarginRight.Set(10)
	pdfGen.Dpi.Set(300)
	pdfGen.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfGen.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	// Generate the PDF
	err = pdfGen.Create()
	if err != nil {
		return nil, err
	}
	return pdfGen.Bytes(), nil
}
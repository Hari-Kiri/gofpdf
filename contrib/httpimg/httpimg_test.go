package httpimg_test

import (
	"github.com/Hari-Kiri/gofpdf"
	"github.com/Hari-Kiri/gofpdf/contrib/httpimg"
	"github.com/Hari-Kiri/gofpdf/internal/example"
)

func ExampleRegister() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPage()

	url := "https://github.com/jung-kurt/gofpdf/raw/master/image/logo_gofpdf.jpg?raw=true"
	httpimg.Register(pdf, url, "")
	pdf.Image(url, 15, 15, 267, 0, false, "", 0, "")
	fileStr := example.Filename("contrib_httpimg_Register")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_httpimg_Register.pdf
}

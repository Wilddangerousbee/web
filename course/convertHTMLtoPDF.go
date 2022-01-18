package main

import (
	"os"
	"os/exec"
)

const converterName = "C:/Program Files/wkhtmltopdf/bin/wkhtmltopdf.exe"

func convertHTMLtoPDF(fileNameHTML string, fileNamePDF string) {
	cmd := exec.Command(converterName, fileNameHTML, fileNamePDF)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

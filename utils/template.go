package utils

import (
	"os"
	"text/template"
)

func CreateFile(filePath string) *os.File {
	file, err := os.Create(filePath)
	ErrorHandler(err)
	return file
}

func ReadSchema(filePath string) *os.File {
	file, err := os.Open(filePath)
	ErrorHandler(err)
	return file
}

func ReadTemplate(filePath string) *template.Template {
	return template.Must(template.ParseFiles(filePath, "./templates/mui/input.tmpl"))
}

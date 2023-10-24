package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Field struct {
	Name     string
	Type     string
	Optional bool
}

func CreateSchemaFields(schema *os.File) struct {
	Fields []Field
	Model  string
} {
	var isModel bool
	var fields []Field

	fileScanner := bufio.NewScanner(schema)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		if line[0] == "model" && os.Args[1] == line[1] {
			isModel = true
			continue
		}
		if isModel && line[0] == "}" {
			break
		}
		if isModel && line[0] != "id" {
			fmt.Println(line)
			fields = append(fields, Field{Name: line[0], Type: assignFieldType(line[1]), Optional: optionalField(line[1])})
		}
	}

	data := struct {
		Fields []Field
		Model  string
	}{fields, strings.ToLower(os.Args[1])}

	return data
}

func optionalField(SchemaType string) bool {
	return strings.Contains(SchemaType, "?")
}

func assignFieldType(SchemaType string) string {
	switch {
	case strings.Contains(SchemaType, "String"):
		return "string"
	case strings.Contains(SchemaType, "Int") || strings.Contains(SchemaType, "Float"):
		return "number"
	default:
		return "string"
	}
}

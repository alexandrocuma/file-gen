package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Field struct {
	Name     string
	Type     string
	Optional bool
}

type Dict struct {
	Models map[string]DictModelField
	Enums map[string]DictEnumField
}

type DictModelField struct {
	Fields map[string][]string
}

type DictEnumField struct {
	Fields map[string]string
}

func CreateSchemaFields(schema *os.File) struct {
	Fields []Field
	Model  string
} {
	var isInsideModel bool
	var isInsideEnum bool
	var fields []Field
	var content string
	var dict Dict
	var currentModel string
	var currentEnum string

	pattern := `(?m)(^(?:enum|model).+)\{(?:.|\n)*?\}`
	fileScanner := bufio.NewScanner(schema)

	for fileScanner.Scan() {
		content += fileScanner.Text() + "\n"
	}

    regex, err := regexp.Compile(pattern)

    if err != nil {
			panic(err)
    }

    matches := regex.FindAllString(content, -1)
    dict.Models = make(map[string]DictModelField)
    dict.Enums = make(map[string]DictEnumField)

    for _, match := range matches {
			line := strings.Split(match, "\n")
			for _, l := range line {
				lineField := strings.Fields(l)

				if lineField[0] == "model" || lineField[0] == "enum" {
					if lineField[0] == "model" {
						currentModel = lineField[1]
						dict.Models[currentModel] = DictModelField{Fields: make(map[string][]string)} 
						isInsideEnum = false
						isInsideModel = true
					}

					if lineField[0] == "enum" {
						currentEnum = lineField[1]
						dict.Enums[currentEnum] = DictEnumField{Fields: map[string]string{}} 
						isInsideModel = false
						isInsideEnum = true
					}
					continue
				}
				
				if isInsideModel && lineField[0] == "}" {
					// reach the end of a block
					isInsideModel = false
					isInsideEnum = false
					currentModel = ""
					currentEnum = ""
					break
				}
				
				if isInsideModel && len(lineField) > 1 && currentModel != "" {
					dict.Models[currentModel].Fields[lineField[0]] = lineField[1:]
					
				}
			
				if isInsideEnum && len(lineField) >= 1 && currentEnum != "" {
					dict.Enums[currentEnum].Fields[lineField[0]] = lineField[0]
				}
			}

			/*
			*/
    }

    fmt.Println("Models: ", dict.Models)
    fmt.Println("Enums: ", dict.Enums)

    /*
		if isModel && line[0] != "id" {
			fmt.Println(line[1])
			fields = append(fields, Field{Name: line[0], Type: assignFieldType(line[1]), Optional: optionalField(line[1])})
		}
    */

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

package main

import (
	"fmt"
	"os"
	"project-creator/utils"
	"strings"
)

func main() {
	fmt.Println("Running project creator")

	schema := utils.ReadSchema("schema.prisma")

	file := utils.CreateFile("./projects/page.tsx")
	router := utils.CreateFile(fmt.Sprintf("./projects/%s.ts", strings.ToLower(os.Args[1])))

	pageTemplate := utils.ReadTemplate("./templates/page.tmpl")
	routerTemplate := utils.ReadTemplate("./templates/router.tmpl")

	data := utils.CreateSchemaFields(schema)
	pageErr := pageTemplate.Execute(file, data)
	utils.ErrorHandler(pageErr)
	routerErr := routerTemplate.Execute(router, data)
	utils.ErrorHandler(routerErr)

	defer file.Close()
	defer router.Close()
	defer schema.Close()
}
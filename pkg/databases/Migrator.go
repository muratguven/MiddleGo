package databases

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	model "middle/pkg/models"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func GetModels() {

	baseEntityType := reflect.TypeOf(model.BaseEntity{})

	var structTypes []reflect.Type

	filepath.Walk("pkg/models/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Println(info)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if filepath.Ext(path) == ".go" {
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				log.Fatal(err)
			}
			for _, decl := range node.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if structType, ok := typeSpec.Type.(*ast.StructType); ok {
								for _, field := range structType.Fields.List {
									fieldType := fset.Position(field.Type.Pos()).String()
									if fieldType == baseEntityType.String() {
										structTypes = append(structTypes, reflect.TypeOf(model.BaseEntity{}))
										break
									}
								}
							}
						}
					}
				}
			}
		}
		return nil
	})

	var entities []interface{}

	for _, t := range structTypes {
		entity := reflect.New(t).Interface()
		entities = append(entities, entity)
	}

	fmt.Println(entities)

	getStructTypes()
}

func getStructTypes() ([]reflect.Type, error) {
	structTypes := []reflect.Type{}
	// baseEntityType := reflect.TypeOf(model.BaseEntity{})

	err := filepath.Walk("pkg/models", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "type ") && strings.Contains(line, " struct {") {
				typeName := strings.TrimSpace(strings.TrimSuffix(strings.Split(line, "type ")[1], " struct {"))
				typeObj := reflect.TypeOf(model.BaseEntity{})
				if _, ok := typeObj.FieldByName("BaseEntity"); !ok {
					continue
				}
				obj := reflect.New(typeObj).Elem().Interface()
				if reflect.TypeOf(obj).Name() == typeName {
					structTypes = append(structTypes, reflect.TypeOf(obj))
					break
				}
			}
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return structTypes, nil
}

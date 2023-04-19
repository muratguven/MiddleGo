package databases

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"go/ast"
    "go/token" 
)

type packageInfo struct {
	name    string
	structs []reflect.Type
}

func GetModels() {
	modelsDir := "./models"
	filepaths, err := filepath.Glob(modelsDir + "/*.go")

	if err != nil {
		log.Fatal(err)
		panic("Models direction reading fail!")
		os.Exit(1)
	}

	// for _, filepath := range filepaths {

	// }

}

func parseFile(filepath string) (*packageInfo, error) {
    fileSet := token.NewFileSet()

    // Parse the file and extract information about its package and structs
    file, err := parser.ParseFile(fileSet, filepath, nil, parser.ParseComments)
    if err != nil {
        return nil, err
    }

    packageInfo := &packageInfo{}
    packageInfo.name = file.Name.Name

    // for _, decl := range file.Decls {
    //     if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
    //         for _, spec := range genDecl.Specs {
    //             if typeSpec, ok := spec.(*ast.TypeSpec); ok {
    //                 if structType, ok := typeSpec.Type.(*ast.StructType); ok {
    //                     packageInfo.structs = append(packageInfo.structs, reflect.StructOf([]reflect.StructField{
    //                         {Name: typeSpec.Name.Name, Type: reflect.TypeOf(structType)},
    //                     }))
    //                 }
    //             }
    //         }
    //     }
    // }

    return packageInfo, nil
}


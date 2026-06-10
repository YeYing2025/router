package model

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestBoolFieldsDoNotUseGormDefaultTrue(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatalf("glob go files: %v", err)
	}
	fset := token.NewFileSet()
	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		}
		source, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("read %s: %v", file, err)
		}
		parsed, err := parser.ParseFile(fset, file, source, parser.ParseComments)
		if err != nil {
			t.Fatalf("parse %s: %v", file, err)
		}
		ast.Inspect(parsed, func(node ast.Node) bool {
			field, ok := node.(*ast.Field)
			if !ok || field.Tag == nil {
				return true
			}
			ident, ok := field.Type.(*ast.Ident)
			if !ok || ident.Name != "bool" {
				return true
			}
			tag, err := strconv.Unquote(field.Tag.Value)
			if err != nil {
				t.Fatalf("unquote tag in %s: %v", file, err)
			}
			gormTag := reflect.StructTag(tag).Get("gorm")
			if strings.Contains(gormTag, "default:true") {
				t.Fatalf("%s: bool field must not use gorm default:true because false can be rewritten during create", fset.Position(field.Pos()))
			}
			return true
		})
	}
}

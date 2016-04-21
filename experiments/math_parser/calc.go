package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"strconv"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func main() {
	variables := make(map[string]float64)
	variables["x"] = 120

	tr, _ := parser.ParseExpr("x-100")

	fmt.Println(typeof(tr))

	X := variables[tr.(*ast.BinaryExpr).X.(*ast.Ident).Name]
	Y, _ := strconv.ParseFloat(tr.(*ast.BinaryExpr).Y.(*ast.BasicLit).Value, 0)
	if tr.(*ast.BinaryExpr).Op.String() == "-" {
		fmt.Println(X - Y)
	} else if tr.(*ast.BinaryExpr).Op.String() == "+" {
		fmt.Println(X + Y)
	}
}

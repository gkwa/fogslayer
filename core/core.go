package core

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/parser"
	"github.com/go-logr/logr"
)

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")
}

func MarshalUnmarshalCUE(input string) (string, error) {
	f, err := parser.ParseFile("", input, parser.ParseComments)
	if err != nil {
		return "", err
	}

	ast.Walk(f, func(n ast.Node) bool {
		if f, ok := n.(*ast.Field); ok && f.Label != nil {
			if ident, ok := f.Label.(*ast.Ident); ok && ident.Name == "b" {
				f.Value = ast.NewString("updated")
				return false
			}
		}
		return true
	}, nil)

	opts := []format.Option{
		format.Simplify(),
		format.UseSpaces(4),
		format.TabIndent(false),
	}

	bytes, err := format.Node(f, opts...)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CountListItems(input, listName string) (int, error) {
	ctx := cuecontext.New()

	value := ctx.CompileString(input)
	if value.Err() != nil {
		return 0, value.Err()
	}

	list := value.LookupPath(cue.ParsePath(listName))
	if !list.Exists() {
		return 0, nil
	}

	iter, err := list.List()
	if err != nil {
		return 0, err
	}

	count := 0
	for iter.Next() {
		count++
	}

	return count, nil
}

func IsItemInList(input, listName, item string) (bool, error) {
	ctx := cuecontext.New()

	value := ctx.CompileString(input)
	if value.Err() != nil {
		return false, value.Err()
	}

	list := value.LookupPath(cue.ParsePath(listName))
	if !list.Exists() {
		return false, nil
	}

	iter, err := list.List()
	if err != nil {
		return false, err
	}

	for iter.Next() {
		val, err := iter.Value().String()
		if err != nil {
			return false, err
		}
		if val == item {
			return true, nil
		}
	}

	return false, nil
}

package evaluator

import (
	"fmt"

	"github.com/devasherr/lambda/object"
)

var builtins = map[string]*object.BuiltIn{
	"len": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			if len(array.Elements) == 0 {
				return NULL
			}

			return array.Elements[0]
		},
	},
	"last": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			if len(array.Elements) == 0 {
				return NULL
			}

			return array.Elements[len(array.Elements)-1]
		},
	},
	"rest": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			if len(array.Elements) == 0 {
				return NULL
			}

			return &object.Array{Elements: array.Elements[1:]}
		},
	},
	"push": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			newArray := make([]object.Object, 0)
			newArray = append(newArray, array.Elements...)
			newArray = append(newArray, args[1])
			return &object.Array{Elements: newArray}
		},
	},
	"puts": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}

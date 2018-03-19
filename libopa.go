package main

import (
  "C"
  "context"
  "github.com/open-policy-agent/opa/ast"
  "github.com/open-policy-agent/opa/rego"
  "encoding/json"
)

var compiler = ast.NewCompiler()

var b = true

//export compile
func compile(m *C.char, f *C.char) {
  var module = C.GoString(m)
  var filename = C.GoString(f)

  parsed, err := ast.ParseModule(filename, module)

  if err != nil {
    // Handle error.
  }

  compiler.Compile(map[string]*ast.Module{filename: parsed})

  if compiler.Failed() {
    // Handle error. Compilation errors are stored on the compiler.
    panic(compiler.Errors)
  }
}

//export query
func query(q *C.char, i *C.char) bool {
  var query = C.GoString(q)
  var input = C.GoString(i)
  var x map[string]interface{}
  json.Unmarshal([]byte(input), &x)

  ctx := context.Background()

  regO := rego.New(
    rego.Query(query),
    rego.Compiler(compiler),
    rego.Input(x),
  )

  // Run evaluation.
  rs, err := regO.Eval(ctx)

  if err != nil {
    // Handle error.
  }

  return rs[0].Expressions[0].Value == true
}

// main function is required, don't know why!
func main() {}

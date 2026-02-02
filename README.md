# Go Language Notes — Complete Quick Reference ⚡️

> Concise, practical notes for learning Go (Golang). Use this as a single-file reference while working through the examples in this repo.

---

## Table of Contents 📚

1. About this Repo
2. Quick Setup (Windows / macOS / Linux)
3. Quick Commands & Tooling
4. Language Basics
5. Composite Types: Arrays, Slices, Maps
6. Structs, Methods & Interfaces
7. Concurrency: Goroutines, Channels, Select, sync
8. Error Handling & Logging
9. Testing, Benchmarks & Examples
10. Modules & Dependency Management
11. Building, Packaging & Deployment
12. Best Practices & Style
13. Useful Links & Resources
14. Repo Examples Map

---

## 1) About this Repo

This repository contains small, focused Go examples (01_hello_world ... 10_concurrency) to practice core language concepts, concurrency, testing, and basic server code.

---

## 2) Quick Setup (Windows / macOS / Linux) 🛠️

- Download from https://go.dev/dl/ and follow the OS installer.
- Verify installation:

```bash
go version
```

- Configure modules (recommended):

```bash
mkdir myproj && cd myproj
go mod init example.com/myproj
```

Notes for Windows: Use MSI installer and ensure `C:\Go\bin` is in PATH. On macOS: use the pkg installer or `brew install go`.

---

## 3) Quick Commands & Tooling 🔧

- go run: `go run main.go`
- go build: `go build ./...` or `go build -o bin/app ./cmd/myapp`
- go test: `go test ./...` (use `-v` for verbose)
- go fmt / goimports: `gofmt -w .` or `go fmt ./...`
- go vet: `go vet ./...`
- godoc / go doc: `go doc fmt` or use `pkg.go.dev` for docs
- go get / go install: `go get module@version`, `go install pkg@latest`
- Environment: `go env`

---

## 4) Language Basics — Syntax & Types ✅

- Variables:

```go
var s string = "hello"
// shorthand
n := 42
```

- Basic types: `bool`, `int`, `int64`, `uint`, `float32`, `float64`, `complex64`, `string`, `byte`, `rune`.
- Constants:

```go
const Pi = 3.14
```

- Control flow: `if`, `for` (only loop), `switch` (can be expression-less), `break`, `continue`.

- Functions:

```go
func add(a, b int) int { return a + b }
```

- Multiple return values & named returns:

```go
func divmod(a, b int) (q, r int) { q = a/b; r = a%b; return }
```

- Pointers:

```go
var x int = 1
p := &x
*p = 5
```

---

## 5) Composite Types: Arrays, Slices, Maps

- Array: fixed size: `var a [3]int`
- Slice: dynamic view over an array:

```go
s := []int{1,2,3}
s = append(s, 4)
```

- Map:

```go
m := map[string]int{"a":1}
val, ok := m["b"]
```

- Range loop:

```go
for i, v := range s { _ = i; _ = v }
```

---

## 6) Structs, Methods & Interfaces

- Structs:

```go
type User struct { ID int; Name string }
```

- Methods (value vs pointer receiver):

```go
func (u *User) Promote() { u.ID++ }
```

- Interfaces & duck-typing:

```go
type Stringer interface { String() string }
```

- Empty interface `interface{}` holds any value.

---

## 7) Concurrency: Goroutines, Channels, Select, sync 🔁

- Goroutine:

```go
go doWork()
```

- Channel:

```go
ch := make(chan int)
ch <- 42    // send
v := <-ch    // receive
```

- Buffered channel: `make(chan int, 10)`
- Select (multiple channel operations):

```go
select {
case v := <-ch1: _ = v
case ch2 <- 1:
default:
}
```

- sync primitives: `sync.Mutex`, `sync.WaitGroup`, `sync.Once`, `atomic` operations and `context.Context` for cancellation.

---

## 8) Error Handling & Logging ⚠️

- Idiomatic errors: return `error` as last return value.

```go
func readFile() (string, error) {
  if err != nil { return "", err }
}
```

- Wrap errors with `fmt.Errorf("%w", err)` and use `errors.Is`/`errors.As`.
- Logging: use standard `log` or structured loggers like `zap` or `logrus` for production.

---

## 9) Testing, Benchmarks & Examples ✅

- Test file: `foo_test.go` with `func TestXxx(t *testing.T)`.
- Run tests: `go test ./...` and `go test -bench=.` for benchmarks.
- Table-driven tests recommended.
- Use `httptest` for testing HTTP handlers.

Example:

```go
func TestAdd(t *testing.T) {
  got := add(1,2)
  if got != 3 { t.Fatalf("want 3 got %d", got) }
}
```

---

## 10) Modules & Dependency Management

- Initialize: `go mod init example.com/myproj`
- Add a dependency: `go get github.com/pkg/errors@v0.9.1`
- Tidy modules: `go mod tidy`
- Replace / pin versions via `go.mod`.

---

## 11) Building, Packaging & Deployment 🚀

- Build for target OS/ARCH:

```bash
GOOS=linux GOARCH=amd64 go build -o myapp ./cmd/myapp
```

- Cross-compilation supported via env vars.
- Use `go install` to place binaries into `$GOPATH/bin` (or `GOBIN`).

---

## 12) Best Practices & Style ✅

- Keep functions small, prefer composition over inheritance.
- Favor slices over arrays unless fixed-size semantics required.
- Name packages clearly; prefer `pkg` subpackages when appropriate.
- Handle errors immediately and explicitly.
- Use `context.Context` for cancellation and request-scoped values.
- Run `gofmt`/`goimports` and `go vet` regularly.

> Tip: Read effective Go and go blog posts frequently for idiomatic patterns.

---

## 13) Useful Links & Resources 🔗

- Official: https://golang.org, https://pkg.go.dev
- Tutorials: https://tour.golang.org
- Blog: https://go.dev/blog
- Style guide & patterns: "Effective Go" and "Go Proverbs"

---

## 14) Repo Examples Map 🔍

- `01_hello_world/` — Hello world and `go run` basics ✅
- `02_variables_constants/` — Variables, constants, types ✅
- `03_data_types/` — Primitive types and conversions ✅
- `04_arrays_slices/` — Arrays, slices, append, capacity ✅
- `05_loops_conditions/` — for, if, switch ✅
- `06_functions/` — functions, multiple return values, variadic ✅
- `07_structs/` — structs, methods ✅
- `08_maps/` — maps, lookups, delete ✅
- `09_error_handling/` — idiomatic error returns ✅
- `10_concurrency/` — goroutines, channels, worker pattern ✅

---

## Appendix — Short Examples

- HTTP server (minimal):

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
})
http.ListenAndServe(":8080", nil)
```

- JSON encode/decode:

```go
type User struct { Name string `json:"name"` }
json.NewEncoder(w).Encode(user)
```

---

## Contributing

Feel free to add more examples or improve explanations. Open a PR with clear description and tests when applicable. ✅

---

If you'd like, I can expand any section with more examples (channel patterns, advanced testing, http middleware, database examples, or build/release CI snippets). 💡

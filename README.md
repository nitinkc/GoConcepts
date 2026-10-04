# GoLang Learning

An incremental Go programming course — docs site + runnable code.

- [https://app.gointerview.dev/quiz](https://app.gointerview.dev/quiz)

- [https://app.gointerview.dev/cheatsheet](https://app.gointerview.dev/cheatsheet)

## Docs
[https://nitinkc.github.io/GoConcepts/](https://nitinkc.github.io/GoConcepts/)


## Run the code

```bash
cd code
go run .           # list lessons
go run . basics    # or . 1 … . 12, . scan for the interactive demo
```

## Build the docs

```bash
uvx --with mkdocs-material mkdocs serve

# OR
pip install -r requirements.txt
mkdocs serve       # live preview at http://127.0.0.1:8000
mkdocs build       # static site in site/
```

### Runnable examples

To make a Go snippet editable and runnable, place it in a tab named exactly `Runnable example`:

````markdown
=== "Runnable example"

    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, Go!")
    }
    ```
````

Regular `go` code fences remain read-only. Runnable snippets may also be short fragments; the runner adds `package main`, `func main()`, and an `fmt` import when needed.

Short fragments are also supported:

````markdown
=== "Runnable example"

    ```go
    name := "Go"
    fmt.Println("Hello,", name)
    ```


## Conventions

Follows idiomatic Go: lowercase package/file names, exported identifiers capitalized,
`go fmt` before committing. See [Effective Go](https://golang.org/doc/effective_go).

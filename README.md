# Resume renderer

A small Go application that loads resume content from YAML and renders a responsive, print-friendly HTML document.

## Run

```sh
go run ./cmd/cli > resume.html
```

Serve the resume over HTTP:

```sh
go run ./cmd/http -http :8080
```

The commands expect to run from the repository root, where `data/resume.yaml` and `html/content.html` are located. Use `-prefix /path/to/checkout` to load those files from another working directory.

## Project structure

- `data/resume.yaml` — resume content, separate from presentation.
- `internal/model` — typed resume schema and application interfaces.
- `internal/yaml` — YAML file loader.
- `internal/html` — HTML template renderer.
- `internal/app` — coordinates loading and rendering.
- `cmd/cli` and `cmd/http` — command-line and HTTP entry points.

Keep experience bullets focused on outcomes, scale, technical decisions, and measurable impact. Only include metrics and claims you can confidently substantiate. The HTML template includes responsive screen styling and print rules for US Letter paper; use your browser's print preview before sending a PDF.

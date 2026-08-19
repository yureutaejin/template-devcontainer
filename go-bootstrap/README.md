# Go bootstrap

Minimal Go application template. Replace `example.com/go-bootstrap` in `go.mod`
with the repository module path before adding imports.

```bash
make check
go run ./cmd/app
```

Keep executable entry points in `cmd/` and application-private packages in
`internal/`. For a reusable library, omit both directories and expose the
library package from the repository root.

# Python bootstrap

Python 3.13 application template managed by [uv](https://docs.astral.sh/uv/).
Dependencies and development tools are declared in `pyproject.toml`; commit the
generated `uv.lock` file.

```bash
uv sync
make check
uv run python-bootstrap
```

The `src/` layout prevents tests from accidentally importing code directly from
the repository root.

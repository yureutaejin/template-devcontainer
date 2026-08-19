.PHONY: format format-check lint typecheck test check

format:
	uv run ruff format .

format-check:
	uv run --locked ruff format --check .

lint:
	uv run --locked ruff check .

typecheck:
	uv run --locked pyright

test:
	uv run --locked pytest

check: format-check lint typecheck test

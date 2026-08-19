.PHONY: format format-check lint test build check

format:
	cargo fmt

format-check:
	cargo fmt --check

lint:
	cargo clippy --all-targets --all-features -- -D warnings

test:
	cargo test --all-features

build:
	cargo build --locked

check: format-check lint test build

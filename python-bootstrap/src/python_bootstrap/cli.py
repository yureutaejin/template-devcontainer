"""Command-line entry point."""


def message(name: str) -> str:
    """Return a greeting for name."""
    return f"hello, {name}"


def main() -> None:
    """Run the application."""
    print(message("from python-bootstrap"))

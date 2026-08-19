from python_bootstrap.cli import message


def test_message() -> None:
    assert message("Ada") == "hello, Ada"

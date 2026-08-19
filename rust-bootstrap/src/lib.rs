//! Domain logic for the application.

/// Returns a greeting for `name`.
#[must_use]
pub fn message(name: &str) -> String {
    format!("hello, {name}")
}

#[cfg(test)]
mod tests {
    use super::message;

    #[test]
    fn formats_a_greeting() {
        assert_eq!(message("Ada"), "hello, Ada");
    }
}

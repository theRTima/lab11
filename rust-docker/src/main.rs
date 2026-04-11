fn greeting(name: &str) -> String {
    format!("Hello, {}!", name)
}

fn main() {
    println!("{}", greeting("World"));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_greeting_world() {
        assert_eq!(greeting("World"), "Hello, World!");
    }

    #[test]
    fn test_greeting_custom_name() {
        assert_eq!(greeting("Rust"), "Hello, Rust!");
    }

    #[test]
    fn test_greeting_is_not_empty() {
        assert!(!greeting("World").is_empty());
    }

    #[test]
    fn test_greeting_contains_name() {
        let name = "Alice";
        assert!(greeting(name).contains(name));
    }
}
fn main() {
    let mut s = String::from("hello world");
    let word = first_word(&s);
    println!("{}", s);
    println!("{}", word);
    s.clear();
    println!("{}", s);

    let mut ss = String::from("hello world");
    let word = first_word2(&ss);
    // ss.clear();
    println!("{}",word)
}

fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}

fn first_word2(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}

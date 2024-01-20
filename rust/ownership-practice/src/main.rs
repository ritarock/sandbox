fn main() {
    let mut s = String::from("hello");
    s.push_str(", world");
    println!("{}", s);

    let s1 = String::from("hello");
    let s2 = s1;

    // println!("{}", s1); // ERROR
    println!("{}", s2);

    let x = 5;
    let y = x;

    println!("{}, {}", x, y);

    let ss = String::from("hello");
    takes_ownership(ss);
    // println!("{}", ss); // ERROR
    let xx = 5;
    makes_copy(xx);
    println!("{}", xx);

    let st1 = gives_ownership();
    println!("{}", st1);
    let st2 = String::from("helleo");
    println!("{}", st2);
    let st3 = takes_and_gives_back(st2);
    // println!("{}", st2); // ERROR
    println!("{}", st3);

    let str1 = String::from("hello");
    let (str2, len) = calculate_length(str1);
    // println!("{}, {}", str1, len); // ERROR
    println!("{}, {}", str2, len);

    let str3 = String::from("hello");
    let len = calculate_length2(&st3);
    println!("{}, {}", str3, len);

    mutfn();
    slicefn();
}

fn takes_ownership(some_string: String) {
    println!("{}", some_string);
}

fn makes_copy(some_string: i32) {
    println!("{}", some_string);
}

fn gives_ownership() -> String {
    let some_string = String::from("hello");
    some_string
}

fn takes_and_gives_back(a_string: String) -> String {
    a_string
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len();
    (s, length)
}

fn calculate_length2(s: &String) -> usize {
    s.len()
}

fn mutfn() {
    let mut s = String::from("hello");
    change(&mut s);
    println!("{}", s);

    let mut st = String::from("hello");
    {
        let r1 = &mut st;
        println!("{}", r1);
    }
    let r2 = &mut st;
    // let r3 = &mut st; // ERROR
    println!("{}", r2);

    let mut ss = String::from("hello");
    let rr1 = &ss;
    let rr2 = &ss;
    // let rr3 = &mut ss; // ERROR
    println!("{}", rr1);
    println!("{}", rr2);
    {
        let rr3 = &mut ss;
        println!("{}", rr3);
    }
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}

fn slicefn() {
    let mut s = String::from("hello world");
    let word = first_word(&s);
    println!("{}", word);
    s.clear();
    println!("{}", s);
    println!("{}", word);

    let st = String::from("hello world");
    let hello = &st[0..5];
    let world = &st[6..11];
    println!("{}", hello);
    println!("{}", world);

    let mut s2 = String::from("hello world");
    let word2 = first_word2(&s2);
    println!("{}", word2);
    s2.clear();
    println!("{}", s);
    // println!("{}", word2); // ERROR
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

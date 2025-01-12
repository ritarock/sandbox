use std::collections::HashMap;

fn main() {
    let v = vec![1, 2, 3, 4, 5];

    let third: &i32 = &v[2];
    println!("The third element is {}", third);

    match v.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element.")
    }

    for i in &v {
        println!("{}", i)
    }

    let mut vv = vec![1, 2, 3, 4, 5];
    for i in &mut vv {
        *i += 10;
    }

    let s = String::new(); // String
    let data = "initial contents"; // &str
    let s = data.to_string(); // String
    let s = String::from("initial contents"); // String

    println!("{}", s);

    let mut s1 = String::from("foo");
    let s2 = "bar";
    s1.push_str(s2);
    println!("s2 is {}", s2);

    let mut s1 = String::from("foo");
    s1.push('b');
    println!("s2 is {}", s1);

    let a = String::from("a");
    let b = String::from("b");
    let c = String::from("c");
    let s = format!("{}-{}-{}",a, b, c);
    println!("{}", s);

    let mut scores = HashMap::new();
    scores.insert(String::from("blue"), 10);
    scores.insert(String::from("yellow"), 50);

    let team_name = String::from("blue");
    let score = scores.get(&team_name);

    println!("{:?}", score);

}

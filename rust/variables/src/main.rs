fn main() {
    let mut x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    println!("The value of x is: {}", x);

    let y = 5;
    let y = y + 1;
    {
        let y = y * 2;
        println!("The value of y is: {}", y);
    }
    println!("The value of y is: {}", y);

    let y = {
        let x = 3;
        x + 1
    };
    println!("The value of y is: {}", y);

    let f = five();
    println!("The value of f is: {}", f);

    let number = 3;
    if number < 5 {
        println!("codition was true");
    } else {
        println!("codition was false");
    }
}

fn five() -> i32 {
    5
}

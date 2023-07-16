fn main() {
    let s = String::from("hello");
    takes_ownership(s); // sが関数にムーブされるので以降でsは使えない
    // println!("{}", s);  これは NG

    let x = 5;
    makes_copy(x); // xが関数にムーブされるが、 i32 は Copy なので以降でもxは使える
    // println!("{}", x); これは OK

    let s1 = gives_ownership(); // 戻り値をs1にムーブする
    let s2 = String::from("hello");
    let s3 = takes_and_gives_back(s2); // s2 が関数にムーブされ、戻り値がs3にムーブされる

    println!("s1 = {}, s3 = {}",s1, s3);

}

fn takes_ownership(some_string: String) {
    println!("{}", some_string);
}

fn makes_copy(some_integer: i32) {
    println!("{}", some_integer);
}

fn gives_ownership() -> String {
    let some_string = String::from("hello");
    some_string
}

fn takes_and_gives_back(a_string: String) -> String {
    a_string
}

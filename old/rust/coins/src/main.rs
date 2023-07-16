#[derive(Debug)]
enum UsState {
    Alabama,
    // Alaska,
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter(UsState),
}
fn main() {
    let penny = value_in_cents(Coin::Penny);
    let nickel = value_in_cents(Coin::Nickel);
    let dime = value_in_cents(Coin::Dime);
    let quarter = value_in_cents(Coin::Quarter(UsState::Alabama));
    println!("{}", penny);
    println!("{}", nickel);
    println!("{}", dime);
    println!("{}", quarter);
}

fn value_in_cents(coin: Coin) -> u32 {
    match coin {
        Coin::Penny => {
            println!("Lucky penny!");
            1
        },
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(state) => {
            println!("State quarter from {:?}!", state);
            25
        },
    }
}

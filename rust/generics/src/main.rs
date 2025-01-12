struct Point<T, U> {
    x: T,
    y: U,
}

impl<T, U> Point<T, U> {
    fn mixup<V, W>(self, other: Point<V, W>) -> Point<T, W> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}

fn main() {
    let number_list = vec![1, 2, 3, 5, 8, 0];
    let result = largest(&number_list);
    println!("{}", result);


    let char_list = vec!['y', 'm', 'a', 'q'];
    let result = largest(&char_list);
    println!("{}", result);

    let p1 = Point { x: 5, y: 10.0};
    let p2 = Point { x: "hello", y: 'c'};
    let p3 = p1.mixup(p2);
    println!("p3.x = {}, p3.y = {}", p3.x, p3.y)
}

// 関数に渡したスライスの型が PartialOrd と Copy を実装できれば良い
fn largest<T: std::cmp::PartialOrd + Copy>(list: &[T]) -> T {
    let mut largest = list[0];

    for &item in list {
        if item > largest {
            largest = item;
        }
    }

    largest
}

pub trait Summary {
    fn summarize(&self) -> String {
        String::from("(Read more)")
    }
}

pub struct NewsArticle {
	pub headline: String,
	pub location: String,
	pub author: String,
	pub content: String,
}

impl Summary for NewsArticle {
	fn summarize(&self) -> String {
		format!("{}, by {} ({})", self.headline, self.author, self.location)
	}
}

pub struct Tweet {
	pub username: String,
	pub content: String,
	pub reply: bool,
	pub retweet: bool,
}

impl Summary for Tweet {
	fn summarize(&self) -> String {
		format!("{}, by {}", self.username, self.content)
	}
}

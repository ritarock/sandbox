use traits::{NewArticle, Summary, Tweet};

fn main() {
    let tweet = Tweet{
        username: String::from("username"),
        content: String::from("content"),
        reply: false,
        retweet: false,
    };

    println!("1 new tweet: {}", tweet.summarize());

    let article = NewArticle{
        headline: String::from("headline"),
        location: String::from("location"),
        author: String::from("author"),
        content: String::from("content"),
    };

    println!("New article available! {}", article.summarize());
}


use clap::{Parser, Subcommand};
use serde::{Deserialize, Serialize};
use std::fs;
use std::path::PathBuf;

#[derive(Debug, Serialize, Deserialize)]
struct Todo {
    id: usize,
    title: String,
    completed: bool,
}

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Add a new todo
    Add {
        /// The title of the todo
        title: String,
    },
    /// List all todos
    List,
    /// Complete a todo
    Complete {
        /// The ID of the todo to complete
        id: usize,
    },
    /// Delete a todo
    Delete {
        /// The ID of the todo to delete
        id: usize,
    },
}

fn load_todos() -> Vec<Todo> {
    let home = std::env::var("HOME").unwrap();
    let todo_path = PathBuf::from(home).join(".todo.json");
    
    if !todo_path.exists() {
        return Vec::new();
    }

    let content = fs::read_to_string(todo_path).unwrap();
    serde_json::from_str(&content).unwrap_or_else(|_| Vec::new())
}

fn save_todos(todos: &[Todo]) {
    let home = std::env::var("HOME").unwrap();
    let todo_path = PathBuf::from(home).join(".todo.json");
    
    let content = serde_json::to_string_pretty(todos).unwrap();
    fs::write(todo_path, content).unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::fs;
    use tempfile::tempdir;
    use std::env;

    fn setup() -> tempfile::TempDir {
        let dir = tempdir().unwrap();
        env::set_var("HOME", dir.path());
        dir
    }

    #[test]
    fn test_add_todo() {
        let _dir = setup();
        let mut todos = Vec::new();
        
        let title = String::from("Test todo");
        let id = todos.len() + 1;
        todos.push(Todo {
            id,
            title: title.clone(),
            completed: false,
        });

        assert_eq!(todos.len(), 1);
        assert_eq!(todos[0].title, title);
        assert_eq!(todos[0].completed, false);
    }

    #[test]
    fn test_complete_todo() {
        let _dir = setup();
        let mut todos = Vec::new();
        
        todos.push(Todo {
            id: 1,
            title: String::from("Test todo"),
            completed: false,
        });

        if let Some(todo) = todos.iter_mut().find(|t| t.id == 1) {
            todo.completed = true;
        }

        assert!(todos[0].completed);
    }

    #[test]
    fn test_delete_todo() {
        let _dir = setup();
        let mut todos = Vec::new();
        
        todos.push(Todo {
            id: 1,
            title: String::from("Test todo"),
            completed: false,
        });

        if let Some(pos) = todos.iter().position(|t| t.id == 1) {
            todos.remove(pos);
        }

        assert!(todos.is_empty());
    }

    #[test]
    fn test_save_and_load_todos() {
        let dir = setup();
        let mut todos = Vec::new();
        
        todos.push(Todo {
            id: 1,
            title: String::from("Test todo"),
            completed: false,
        });

        save_todos(&todos);
        let loaded_todos = load_todos();

        assert_eq!(todos.len(), loaded_todos.len());
        assert_eq!(todos[0].id, loaded_todos[0].id);
        assert_eq!(todos[0].title, loaded_todos[0].title);
        assert_eq!(todos[0].completed, loaded_todos[0].completed);

        fs::remove_file(dir.path().join(".todo.json")).unwrap();
    }
}

fn main() {
    let cli = Cli::parse();
    let mut todos = load_todos();

    match cli.command {
        Commands::Add { title } => {
            let id = todos.len() + 1;
            todos.push(Todo {
                id,
                title,
                completed: false,
            });
            println!("Added todo #{}", id);
        }
        Commands::List => {
            if todos.is_empty() {
                println!("No todos found");
                return;
            }
            
            for todo in &todos {
                let status = if todo.completed { "✓" } else { " " };
                println!("[{}] #{}: {}", status, todo.id, todo.title);
            }
        }
        Commands::Complete { id } => {
            if let Some(todo) = todos.iter_mut().find(|t| t.id == id) {
                todo.completed = true;
                println!("Completed todo #{}", id);
            } else {
                println!("Todo #{} not found", id);
            }
        }
        Commands::Delete { id } => {
            if let Some(pos) = todos.iter().position(|t| t.id == id) {
                todos.remove(pos);
                println!("Deleted todo #{}", id);
            } else {
                println!("Todo #{} not found", id);
            }
        }
    }

    save_todos(&todos);
}

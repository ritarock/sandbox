// Learn more about Tauri commands at https://tauri.app/develop/calling-rust/
#[tauri::command]
fn my_custom_command() {
    println!("I was invoked from JavaScript!");
}

#[tauri::command]
fn my_custom_command2(invoke_message: String) {
    println!("I was invoked from JavaScript, with this message: {}", invoke_message);
}

#[tauri::command]
fn my_custom_command3() -> String {
    "Hello from Rust".into()
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_opener::init())
        .invoke_handler(tauri::generate_handler![
            my_custom_command,
            my_custom_command2,
            my_custom_command3,
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}

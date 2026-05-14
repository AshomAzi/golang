# 📞 Basic Phonebook CLI
A lightweight Command Line Interface (CLI) application built in Go to manage contacts. This project demonstrates the practical use of Go Maps for efficient data storage, retrieval, and state management.

## 🚀 Features
This application provides a simple interface to manage a digital phonebook with the following capabilities:

1. Create: Add new contacts (Name and Phone Number).

2. Read: Look up a phone number by entering a name.

3. Update: Change the phone number for an existing contact.

4. Delete: Remove a contact from the phonebook.

5. List All: Display all stored contacts in the current session.

## 🛠️ Technical Concepts Used
To build this, the following Go concepts were implemented:

Maps (map[string]string): Used as the primary in-memory database where the name is the key and the phone number is the value.

Control Flow: for loops and switch statements to handle the CLI menu.

User Input: Using the fmt or bufio package to capture user commands.

Error Handling: Checking if a contact exists before attempting to delete or update it.

📋 Prerequisites
Go (version 1.18 or higher recommended)

🔧 How to run
Bash
Run the application:

Bash
   go run main.go
📖 Usage Example
When you run the app, you will be prompted with a menu:

Plaintext
--- Phonebook Menu ---
1. Add Contact
2. Search Contact
3. Update Contact
4. Delete Contact
5. List All
6. Exit

Choose an option: 1
Enter Name: Alice
Enter Phone: 555-0102
Contact added successfully!
📝 Implementation Details
In this project, the map is defined as:

Go
phonebook := make(map[string]string)
The Search function utilizes the "comma ok" idiom to check if a name exists:

Go
number, exists := phonebook[name]
if !exists {
    fmt.Println("Contact not found.")
}
🤝 Contributing
U can decide to saving to a .txt or .json file

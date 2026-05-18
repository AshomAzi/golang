package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddContact(name, number string, contacts map[string]string) {

	for key := range contacts {
		if key == name {
			fmt.Println("key already exist with the number:", contacts[key])
		}
	}
	contacts[name] = number

	data, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		writer := bufio.NewWriter(data)
		writer.WriteString(name + ":" + number + "\n")
		fmt.Println("User created successgully.")
		writer.Flush()
	} else {
		fmt.Print("Error: ", err)
	}
}

func SearchName(input, file string) {

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Print("Error:", err)
	}

	newContent := strings.Fields(string(content))
	for _, v := range newContent {
		if strings.Contains(v, input) {
			fmt.Println(v)
			break
		} else {
			fmt.Println("Contact does not exist!")
			break
		}
	}
}

func TotalContacts(contacts map[string]string) int {
	return len(contacts)
}

func main() {
	fmt.Println("---Phonebook Menu---")
	fmt.Println("1. Add contact")
	fmt.Println("2. Search Contact")
	fmt.Println("3. Update contact")
	fmt.Println("4. Delete Contact")
	fmt.Println("5. List all contacts")
	fmt.Println("6. Exit")

	contacts := make(map[string]string)

	for {
		fmt.Print("Chose any number from the options above: ")
		option := bufio.NewScanner(os.Stdin)
		option.Scan()
		each := strings.Fields(option.Text())
		if len(each) == 1 {
			if each[0] == "1" {
				fmt.Print("Input the contacts name: ")
				name := bufio.NewScanner(os.Stdin)
				name.Scan()
				mName := strings.Fields(name.Text())
				if len(mName) == 1 {
					fmt.Print("Input the contacts number: ")
					number := bufio.NewScanner(os.Stdin)
					number.Scan()
					num := strings.Fields(number.Text())
					if len(num) == 1 {
						AddContact(mName[0], num[0], contacts)
					} else {
						fmt.Println("Invalid Number!")
					}
				} else {
					fmt.Println("Invalid name!")
				}
			} else if each[0] == "2" {
				fmt.Print("Input the name (key) to search for: ")
				rawname := bufio.NewScanner(os.Stdin)
				rawname.Scan()
				name := rawname.Text()
				newName := strings.Fields(name)
				if !strings.Contains(newName[0], " ") {
					SearchName(newName[0], "data.txt")
				} else {
					fmt.Println("Invalid key, name must not contain space!")
				}
			} else if each[0] == "6" {
				os.Exit(1)
			} else {
				fmt.Println("Invalid Option!")
			}
		} else {
			fmt.Println("Invalid option!")
		}
	}

}

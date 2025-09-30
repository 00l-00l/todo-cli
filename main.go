package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	todos := []string{} // Храним задачи в памяти
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🎯 TODO-CLI v1.1")
	fmt.Println("Команды: add, list, delete, exit")

	for {
		fmt.Print("\nВведите команду: ")
		scanner.Scan()
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch command {
		case "add":
			fmt.Print("Введите задачу: ")
			scanner.Scan()
			task := scanner.Text()
			if task != "" {
				todos = append(todos, task)
				fmt.Printf("✅ Добавлено: %s\n", task)
			}

		case "list":
			if len(todos) == 0 {
				fmt.Println("📝 Список пуст")
			} else {
				fmt.Println("📝 Ваши задачи:")
				for i, task := range todos {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case "delete":
			if len(todos) == 0 {
				fmt.Println("📝 Список пуст, нечего удалять")
			} else {
				fmt.Print("Введите номер задачи для удаления: ")
				scanner.Scan()
				var index int
				_, err := fmt.Sscanf(scanner.Text(), "%d", &index)
				if err == nil && index > 0 && index <= len(todos) {
					removed := todos[index-1]
					todos = append(todos[:index-1], todos[index:]...)
					fmt.Printf("❌ Удалено: %s\n", removed)
				} else {
					fmt.Println("❌ Неверный номер задачи")
				}
			}

		case "exit":
			fmt.Println("👋 Выход...")
			return

		default:
			fmt.Println("❌ Неизвестная команда. Доступные: add, list, delete, exit")
		}
	}
}

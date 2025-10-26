package main

import (
	"bashnya-hw3/internal/tree"
	"bufio"
	"fmt"
	"os"

	"strconv"
	"strings"
)

func runTreeModification(t *tree.Tree, command string, number int) string {
	switch command {
	case "insert":
		res := t.Insert(number)
		if res {
			return fmt.Sprintf("Элемент %d добавлен", number)
		} else {
			return fmt.Sprintf("Элемент %d уже существует", number)
		}
	case "remove":
		res := t.Remove(number)
		if res {
			return fmt.Sprintf("Элемент %d удален", number)
		} else {
			return fmt.Sprintf("Элемент %d не найден", number)
		}
	default:
		return "Неизвестная команда"
	}

}

func queryTreeData(t *tree.Tree, num int) (string, *tree.Node) {
	node := t.Find(num)
	if node != nil {
		return fmt.Sprintf("Элемент %d найден", num), node
	}
	return fmt.Sprintf("Элемент %d не найден", num), nil
}

func handleSimpleCommand(tr *tree.Tree, command string) {
	switch command {
	case "depth":
		fmt.Printf("Глубина дерева: %d\n", tr.Depth())
	case "print":
		tr.Print()
	default:
		fmt.Println("Неизвестная команда")
	}
}

func main() {
	tr := &tree.Tree{}
	fmt.Println("Введите одну из команд, где X - целое число int: insert X, remove X, find X, depth, print, exit")
	scanner := bufio.NewScanner(os.Stdin)
mainLoop:
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		value := scanner.Text()
		parts := strings.Split(value, " ")
		if len(parts) == 0 {
			continue
		}
		if parts[0] == "exit" {
			fmt.Println("Выход из программы")
			os.Exit(0)
		}
		switch parts[0] {
		case "insert", "remove", "find":
			if len(parts) < 2 {
				fmt.Printf("Ошибка: команда '%s' требует аргумент\n", parts[0])
				continue mainLoop
			}
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Ошибка: введите целое число")
				continue mainLoop
			}
			if parts[0] == "find" {
				text, _ := queryTreeData(tr, num)
				fmt.Println(text)
			} else {
				result := runTreeModification(tr, parts[0], num)
				fmt.Println(result)
			}
		case "print", "depth":
			handleSimpleCommand(tr, parts[0])

		default:
			fmt.Println("Неизвестная команда")
		}
	}

}

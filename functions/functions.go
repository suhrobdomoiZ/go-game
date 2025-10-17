package functions

import (
	"fmt"
	. "go-game/structures"
	"strings"
)

var player *Player
var world map[string]*Room
var doorLocked bool

func InitGame() {

	fmt.Println("Начинаем игру......")
	fmt.Println("Пенис активирован.......")
	player = &Player{"кухня", false, []string{}}
	world = map[string]*Room{
		"кухня":   {"кухня", "ты находишься на кухне.", map[string][]string{"на столе": {"чай"}}, []string{"коридор"}, "надо собрать рюкзак и идти в универ."},
		"коридор": {"коридор", "ничего интересного.", make(map[string][]string), []string{"кухня", "комната", "улица"}, ""},
		"комната": {"комната", "ты в своей комнате.", map[string][]string{"на столе": {"ключи", "конспекты"}, "на стуле": {"рюкзак"}}, []string{"коридор"}, ""},
		"улица":   {"улица", "на улице весна.", make(map[string][]string), []string{"домой"}, ""},
	}
	doorLocked = true

}

func itemsDisplay(currentRoom *Room) string {
	if currentRoom.Items == nil {
		return ""
	} else {
		var locationItems []string
		var counter = 0
		for location, items := range currentRoom.Items {
			counter++
			locationItems = append(locationItems, location)
			locationItems = append(locationItems, ": ")
			for i := 0; i < len(items); i++ {
				if i != len(items)-1 {
					locationItems = append(locationItems, items[i])
					locationItems = append(locationItems, ", ")
				} else {
					if counter < len(currentRoom.Items) {
						locationItems = append(locationItems, items[i])
						locationItems = append(locationItems, ", ")
					} else {
						locationItems = append(locationItems, items[i])
						locationItems = append(locationItems, ".")
					}
				}
			}
		}
		return strings.Join(locationItems, "")
	}
}

func waysDisplay(currentRoom *Room) string {
	if currentRoom.Ways == nil {
		return ""
	} else {
		ways := []string{"можно пройти - "}
		for i := 0; i < len(currentRoom.Ways); i++ {
			ways = append(ways, currentRoom.Ways[i])
			if i != len(currentRoom.Ways)-1 {
				ways = append(ways, ", ")
			} else {
				ways = append(ways, ".")
			}
		}
		return strings.Join(ways, "")
	}
}

func messageDisplay(currentRoom *Room) string {
	switch currentRoom.Name {
	case "кухня":
		if !player.HasBackpack {
			currentRoom.Message = "надо собрать рюкзак и идти в универ."
			return currentRoom.Message
		} else {
			return ""
		}
	default:
		return currentRoom.Message
	}

}

func grabDisplay(currentRoom *Room, item string) string {
	for location, items := range currentRoom.Items {
		for i := 0; i < len(items); i++ {
			if items[i] == item && player.HasBackpack {
				currentRoom.Items[location] = append(items[:i], items[i+1:]...)
				player.Inventory = append(player.Inventory, item)
				return fmt.Sprintf("предмет добавлен в инвентарь: %s", item)
			} else if items[i] == item && !player.HasBackpack {
				return "некуда класть"
			}
		}
	}
	return "нет такого"
}

func goDisplay(word string) string {
	for _, way := range world[player.Room].Ways {
		if way == word {
			if word != "улица" {
				player.Room = world[word].Name
				return fmt.Sprintf("%s %s %s %s", world[player.Room].Description, itemsDisplay(world[player.Room]), messageDisplay(world[player.Room]), waysDisplay(world[player.Room]))
			} else {
				if doorLocked {
					return "дверь закрыта"
				} else {
					player.Room = "улица"
					return fmt.Sprintf("%s %s %s %s", world[player.Room].Description, itemsDisplay(world[player.Room]), messageDisplay(world[player.Room]), waysDisplay(world[player.Room]))
				}
			}

		}
	}
	return fmt.Sprintf("нет пути в %s", word)
}

func doDisplay(item string, place string) string {
	for i := 0; i < len(player.Inventory); i++ {
		if player.Inventory[i] == item {
			if item == "ключи" && place == "дверь" && player.Room == "коридор" {
				doorLocked = false
				return "дверь открыта"
			} else {
				return "не к чему применить"
			}
		}
	}
	return fmt.Sprintf("нет предмета в инвентаре - %s", item)
}

func HandleCommand(command string) string {
	var commandParsed = strings.Split(command, " ")
	switch commandParsed[0] {
	case "осмотреться":
		if len(commandParsed) == 1 {
			return fmt.Sprintf("%s %s %s %s", world[player.Room].Description, itemsDisplay(world[player.Room]), messageDisplay(world[player.Room]), waysDisplay(world[player.Room]))
		}
	case "взять":
		if len(commandParsed) == 2 {
			return grabDisplay(world[player.Room], commandParsed[1])
		}
	case "надеть":
		if len(commandParsed) == 2 {
			if player.Room == "комната" && !player.HasBackpack && commandParsed[1] == "рюкзак" {
				delete(world[player.Room].Items, "на стуле")
				player.HasBackpack = true
				return "вы надели: рюкзак"
			} else {
				return "нет такого"
			}
		}
	case "идти":
		if len(commandParsed) == 2 {
			return goDisplay(commandParsed[1])
		}
	case "применить":
		if len(commandParsed) == 3 {
			return doDisplay(commandParsed[1], commandParsed[2])
		}
	}
	return "неизвестная команда"
}

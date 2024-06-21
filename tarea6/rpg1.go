package main

import (
	"fmt"
	"time"


)

type Player struct {
	Name     string
	Level    int
	Hp       int
	Attack   int
	Defense  int
	Location Location
	Gold     int
}

type Enemy struct {
	Name     string
	Hp       int
	Attack   int
	Defense  int
	Location Location
	GoldDrop int
}

type Location string

const (
	The_Island     Location = "The Island"
	Scorched_Earth Location = "Scorched Earth"
	The_Center     Location = "The Center"
	Abberation     Location = "Abberation"
	Extinction     Location = "Extinction"
	Genesis        Location = "Genesis"
)


type Item struct {
	Name         string
	Description  string
	Price        int
	AttackBoost  int
	DefenseBoost int
	HealthBoost  int
}

var shopItems = map[int]Item{
	1: {
		Name:        "Tek Gloves",
		Description: "Increase attack by 20",
		Price:       125,
		AttackBoost: 20,
	},
	2: {
		Name:         "Tek Helmet",
		Description:  "Increase defense by 30",
		Price:        150,
		DefenseBoost: 30,
	},
	3: {
		Name:        "Tek Boots",
		Description: "Increase health by 100",
		Price:       200,
		HealthBoost: 100,
	},
}

func createPlayer(name string) *Player {
	player := &Player{
		Name:     name,
		Level:    1,
		Hp:       100,
		Attack:   20,
		Defense:  20,
		Location: The_Island,
		Gold:     0,
	}
	return player
}

func movePlayer(player *Player, location Location) {
	player.Location = location
}

func createEnemy(enemy_map map[Location]Enemy, locations *[]Location) {
	manticore := Enemy{
		Name:     "Manticore",
		Hp:       100,
		Attack:   20,
		Defense:  20,
		Location: Scorched_Earth,
		GoldDrop: 50,
	}
	enemy_map[Scorched_Earth] = manticore
	*locations = append(*locations, Scorched_Earth)

	dragon := Enemy{
		Name:     "Dragon",
		Hp:       110,
		Attack:   40,
		Defense:  30,
		Location: The_Center,
		GoldDrop: 75,
	}
	enemy_map[The_Center] = dragon
	*locations = append(*locations, The_Center)

	rockwell := Enemy{
		Name:     "Rockwell",
		Hp:       120,
		Attack:   50,
		Defense:  40,
		Location: Abberation,
		GoldDrop: 100,
	}
	enemy_map[Abberation] = rockwell
	*locations = append(*locations, Abberation)

	kingTitan := Enemy{
		Name:     "King Titan",
		Hp:       130,
		Attack:   60,
		Defense:  50,
		Location: Extinction,
		GoldDrop: 150,
	}
	enemy_map[Extinction] = kingTitan
	*locations = append(*locations, Extinction)

	corruptedMasterController := Enemy{
		Name:     "Corrupted Master Controller",
		Hp:       140,
		Attack:   70,
		Defense:  60,
		Location: Genesis,
		GoldDrop: 300,
	}
	enemy_map[Genesis] = corruptedMasterController
	*locations = append(*locations, Genesis)
}

func battle(player *Player, location Location, enemy_map map[Location]Enemy) {
	enemy, up := enemy_map[location]
	if !up {
		fmt.Println("The boss is not in the arena")
		return
	}
	fmt.Printf("The boss %s has been summoned in the %s arena\n", enemy.Name, location)
	fmt.Printf("The boss has %d hp, %d attack, and %d defense\n", enemy.Hp, enemy.Attack, enemy.Defense)
	fmt.Printf("You have %d hp, %d attack, and %d defense\n", player.Hp, player.Attack, player.Defense)
	fmt.Println("What do you want to do?\n 1: Attack\n 2: Visit HLNA's shop\n 3: Leave")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		startBattle(player, location, enemy_map, enemy)
	case 2:
		visitShop(player)
	case 3:
		return
	default:
		break
	}
}

func startBattle(player *Player, location Location, enemy_map map[Location]Enemy, enemy Enemy) {
	turn := 0
	for player.Hp > 0 && enemy.Hp > 0 {
		if turn%2 == 0 {
			fmt.Println("\nPlayer's turn")
			if attack(player.Attack, &enemy.Defense, &enemy.Hp) {
				fmt.Printf("\nYou have defeated the arena boss %s\n", enemy.Name)
				statusPlayer(player)
				delete(enemy_map, location)
				fmt.Printf("\nYou have overcome the boss on the map %s!\n", location)
				player.Gold += enemy.GoldDrop 
				if player.Hp < 100 && player.Defense < 50 {
					choiseReward(player)
				}
				if player.Defense < 20 {
					choiseItem(player, location)
				}
				time.Sleep(2 * time.Second)
				checkShopAvailability(player)
				break
			}
		} else {
			fmt.Println("\nEnemy's turn")
			if attack(enemy.Attack, &player.Defense, &player.Hp) {
				fmt.Printf("\nThe boss %s has defeated you.\n", enemy.Name)
				enemy_map = make(map[Location]Enemy)
				break
			}
		}
		turn++
		time.Sleep(2 * time.Second)
	}
}

func attack(attack int, defense *int, hp *int) bool {
	hit := *defense - attack
	if hit > 0 {
		*defense = hit
	} else {
		*defense = 0
		*hp += hit
	}
	fmt.Printf("You have done %d damage. Remaining health: %d\n", attack, *hp)
	return *hp <= 0
}

func statusPlayer(player *Player) {
	fmt.Printf("\nYour hero %s has %d health, %d attack and %d defense\n", player.Name, player.Hp, player.Attack, player.Defense)
	time.Sleep(2 * time.Second)
}

func transmitterAvailable(locations []Location, enemy_map map[Location]Enemy) {
	for i, el := range locations {
		if _, ok := enemy_map[el]; ok {
			fmt.Printf("\n%d: %s\n", i+1, el)
		}
	}
}

func choiseReward(player *Player) {
	fmt.Println("\nYou have won a great battle and as a reward we offer you the following rewards. Choose wisely:\n")
	fmt.Println(" 1. Enduro Stew: +10 attack +110 HP\n 2. Focal Chili: +10 defense +110 HP")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		player.Attack += 10
		player.Hp += 110
		fmt.Println("You have received an Enduro Stew")
	case 2:
		player.Defense += 10
		player.Hp += 110
		fmt.Println("You have received a Focal Chili")
	default:
		fmt.Println("You have not chosen any reward")
	}
	time.Sleep(3 * time.Second)
	statusPlayer(player)
}

func choiseItem(player *Player, location Location) {
	fmt.Printf("\nYou have found a treasure in %s containing some items inside. Which one do you want to take?\n", location)
	fmt.Println("\n1. Tek Leggings: +40 attack\n")
	fmt.Println("\n2. Tek Chestpiece: +50 HP\n")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		player.Attack += 40
	case 2:
		player.Hp += 50
		break
	}
	statusPlayer(player)
}

func visitShop(player *Player) {
	fmt.Println("\nWelcome to HLNA's shop! Here are the items available for purchase:")
	for i := 1; i <= len(shopItems); i++ {
		item := shopItems[i]
		fmt.Printf("%d: %s - %s, Price: %d Gold\n", i, item.Name, item.Description, item.Price)
	}
	fmt.Println("\nYou currently have", player.Gold, "Gold.")
	fmt.Println("\nEnter the number of the item you want to purchase or enter '0' to return:")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	item, exists := shopItems[choice]
	if !exists {
		fmt.Println("Invalid choice. Returning to map selection.")
		return
	}
	if player.Gold >= item.Price {
		player.Gold -= item.Price
		player.Attack += item.AttackBoost
		player.Defense += item.DefenseBoost
		player.Hp += item.HealthBoost
		fmt.Printf("You have purchased %s from HLNA's shop!\n", item.Name)
		statusPlayer(player)
	} else {
		fmt.Println("Insufficient Gold. Please defeat more enemies to earn enough Gold.")
		time.Sleep(2 * time.Second)
		visitShop(player)
	}
}

func checkShopAvailability(player *Player) {
	if (player.Level-1)%2 == 0 && player.Level > 1 {
		fmt.Println("\nCongratulations! You have won 2 battles in a row. Visit HLNA's shop to upgrade your gear!")
		time.Sleep(2 * time.Second)
		visitShop(player)
		player.Level++
	}
}

func main() {
	enemy_map := make(map[Location]Enemy)
	var locations []Location
	createEnemy(enemy_map, &locations)
	fmt.Println("Locations:", locations)
	fmt.Println("Welcome to ARK GO\nEnter your name:")
	var namePlayer string
	fmt.Scanln(&namePlayer)

	player := createPlayer(namePlayer)
	fmt.Println("\nChoose the map:")
	for {
		if len(enemy_map) == 0 {
			fmt.Println("\nYou have managed to defeat all the bosses!")
			break
		}
		if player.Hp < 1 {
			fmt.Println("You have been defeated, a raptor ate your body. Try again")
			break
		}
		transmitterAvailable(locations, enemy_map)
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			movePlayer(player, Scorched_Earth)
			battle(player, Scorched_Earth, enemy_map)
			continue
		case 2:
			movePlayer(player, The_Center)
			battle(player, The_Center, enemy_map)
			continue
		case 3:
			movePlayer(player, Abberation)
			battle(player, Abberation, enemy_map)
			continue
		case 4:
			movePlayer(player, Extinction)
			battle(player, Extinction, enemy_map)
			continue
		case 5:
			movePlayer(player, Genesis)
			battle(player, Genesis, enemy_map)
		case 6:
			movePlayer(player, The_Island)
			fmt.Println("\nHas llegado a la casa, el juego ha terminado!")
			break
		default:
			break
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	userInput := bufio.NewReader(os.Stdin)

	populateDeck()
	shuffleDeck()


	fmt.Println("Lets play blackjack")
	total := 0
	c := hitMe()
	total += getTrueFace(c)
	fmt.Println(nameCard(c))
	c = hitMe()
	total += getTrueFace(c)
	fmt.Println(nameCard(c))

	for{
		if total > 21 {
			fmt.Println("YOU BUSTED")
			return
		}
		fmt.Printf("you have %d on top\nHIT or PASS?\n<<", total)
		input,_ := userInput.ReadString('\n')
		
		input = strings.TrimSpace(input)
		if input == "HIT"{
			c = hitMe()
			fmt.Println(nameCard(c))
			total += c.face
		}else if input == "PASS"{
			break;
		}else{
			fmt.Println("Bad Input")
		}
	}
	fmt.Println("you made it")
}

func checkAce(c Card) bool {
	return c.face == 14
}
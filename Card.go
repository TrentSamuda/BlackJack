package main

import (
	"fmt"
	"strconv"
)

type Card struct {
	face int
	suit string
}

func getFace(c Card) string {
			if c.face > 10 {
			switch c.face {
		case 11:
			return "jack"
		case 12:
			return "queen"
		case 13:
			return "king"
		case 14:
			return "ace"
		}
	}else {
		return strconv.Itoa(c.face)
	}
	fmt.Println("Should never hit")
	return strconv.Itoa(0)
}

func nameCard(c Card) string{
	return getFace(c) + " of " + c.suit
}

func getTrueFace(c Card) int {
	if c.face > 10 {
		switch c.face {
		case 11:
			return 10
		case 12:
			return 10
		case 13:
			return 10
		case 14:
			return 11
		}
	}else {
		return c.face
	}
	fmt.Println("Should never hit")
	return 0
}
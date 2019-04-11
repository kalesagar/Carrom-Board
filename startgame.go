package main

import "fmt"

//StartTheGame ...
func StartTheGame(p1, p2 *Player, c *Coins) {
	fmt.Println("Start The GAME...!!!")
	var flg bool
	for {
		fmt.Println("Select Strike:")
		fmt.Println(`1. Strike
		2. Multistrike
		3. Red strike
		4. Striker strike
		5. Defunct coin
		6. None`)
		var move int
		fmt.Scanf("%d", &move)
		if !flg {
			ExecuteMove(p1, c, move)
			flg = true
		} else {
			ExecuteMove(p2, c, move)
			flg = false
		}
		if CoinsExhausted(c) {
			break
		}
		fmt.Printf("p1 name: %s, points: %d \n", p1.Name, p1.Points)
		fmt.Printf("p2 name: %s, points: %d \n", p2.Name, p2.Points)
		fmt.Printf("Black: %d, Red:%d\n", c.Black, c.Red)
	}
	if (p1.Points-p2.Points) >= 3 && p1.Points > 5 {
		fmt.Printf("%s has Won the Game, Final Score: %d : %d\n", p1.Name, p1.Points, p2.Points)
	} else if (p2.Points-p1.Points) >= 3 && p2.Points > 5 {
		fmt.Printf("%s has Won the Game, Final Score: %d : %d\n", p2.Name, p2.Points, p1.Points)
	} else {
		fmt.Printf("The Game is Drawn.Final Score:(p1.Name:p2.Name) is (%d : %d)\n", p1.Points, p2.Points)
	}
}

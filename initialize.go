package CarromBoardGame

import "fmt"

//InitializeTheGame ...
func InitializeTheGame() (*Player, *Player, *Coins) {
	var player1Name, player2Name string
	var pl1Age, pl2Age int
	var pl1Gen, pl2Gen string
	fmt.Println("Enter the Details of playerOne Name, age, Gender respectively...")
	fmt.Scanf("%s", &player1Name)
	fmt.Scanf("%d", &pl1Age)
	fmt.Scanf("%s", &pl1Gen)
	fmt.Println("Enter the Details of playerTwo Name, age, Gender respectively...")
	fmt.Scanf("%s", &player2Name)
	fmt.Scanf("%d", &pl2Age)
	fmt.Scanf("%s", &pl2Gen)
	p1 := AddPlayer(player1Name, pl1Age, pl1Gen)
	p2 := AddPlayer(player2Name, pl2Age, pl2Gen)
	c := Coins{9, 1}
	return p1, p2, &c
}

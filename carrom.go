package main

//Player ...
type Player struct {
	Name   string
	Age    int
	Gender string
	Points int
	Moves  []int
}

//Coins ...
type Coins struct {
	Black int
	Red   int
}

//AddPlayer ...
func AddPlayer(name string, age int, gender string) *Player {
	player := Player{name, age, gender, 0, []int{}}
	return &player
}

//MoveStrike ...
func (p *Player) MoveStrike(c *Coins) *Player {
	if c.Black >= 1 {
		c.Black = c.Black - 1
		p.Points = p.Points + 1
		p.Moves = append(p.Moves, 1)
	} else {
		p.Moves = append(p.Moves, 0)
	}
	return p
}

//MoveMultiStrike ...
func (p *Player) MoveMultiStrike(c *Coins) *Player {
	if c.Black >= 1 {
		p.Points = p.Points + 2
		p.Moves = append(p.Moves, 2)
	} else {
		p.Moves = append(p.Moves, 0)
	}
	return p
}

//MoveRedCoin ...
func (p *Player) MoveRedCoin(c *Coins) *Player {
	if c.Red == 1 {
		p.Points = p.Points + 3
		p.Moves = append(p.Moves, 3)
		c.Red = 0
	} else {
		p.Moves = append(p.Moves, 0)
	}
	return p
}

//MoveStrikerStrike ...
func (p *Player) MoveStrikerStrike(c *Coins) *Player {
	p.Points = p.Points - 1
	p.Moves = append(p.Moves, -1)
	return p
}

//MoveDefuncCoin ...
func (p *Player) MoveDefuncCoin(c *Coins) *Player {
	if c.Black >= 2 {
		c.Black = c.Black - 2
		p.Points = p.Points - 2
		p.Moves = append(p.Moves, -2)
	} else {
		p.Moves = append(p.Moves, 0)
	}
	return p
}

//CheckLastThreeMoves ...
func (p *Player) CheckLastThreeMoves() *Player {
	if len(p.Moves) < 3 {
		return p
	}
	LastThreeMoves := p.Moves[len(p.Moves)-3:]
	var negvalCheck bool
	for _, val := range LastThreeMoves {
		if val > 0 {
			negvalCheck = true
		}
	}
	if !negvalCheck {
		p.Points = p.Points - 1
	}
	return p
}

//CoinsExhausted ...
func CoinsExhausted(c *Coins) bool {
	if c.Black == 0 && c.Red == 0 {
		return true
	}
	return false
}

//ExecuteMove ...
func ExecuteMove(p1 *Player, c *Coins, move int) {
	switch move {
	case 1:
		p1.MoveStrike(c)
	case 2:
		p1.MoveMultiStrike(c)
	case 3:
		p1.MoveRedCoin(c)
	case 4:
		p1.MoveStrikerStrike(c)
	case 5:
		p1.MoveDefuncCoin(c)
	}
}

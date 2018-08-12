package main

import (
	"fmt"
)

type Follower struct {
	name string
	partnerships []*Partnership
	bestScore int
	bestScorePartnership *Partnership
	lastPartnership *Partnership
}

func (f Follower) String() (string){
	return fmt.Sprintf("%s\nBest Score: %d with %s\nLast Score: %d with %s\nLast competed %s\n",
		f.name,
		f.bestScore,
		f.bestScorePartnership.names,
		f.lastPartnership.highScore,
		f.lastPartnership.names,
		f.lastPartnership.lastCompetition.Format("01/02/2006"),
		)
}
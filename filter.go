package main

import (
	"time"
	"fmt"
)

func FilterFollowers(followers []*Follower, monthsSinceCompeting int, scoreDisparity int, minScore int, maxScore int) ([]*Follower) {
	fmt.Println("Filtering followers")
	matching := make([]*Follower, 0)
	for _, follower := range followers {
		if (follower.lastPartnership.highScore <= follower.bestScore - scoreDisparity ||
			follower.lastPartnership.lastCompetition.AddDate(0, monthsSinceCompeting, 0).Before(time.Now())) &&
			follower.bestScore > minScore && follower.bestScore <= maxScore &&
			follower.lastPartnership.highScore > 0 {
				matching = append(matching, follower)
		}
	}

	return matching
}

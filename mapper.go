package main

import (
	"sort"
	"fmt"
)

func MapFollowers(followerMap map[string][]*Partnership) ([]*Follower){
	fmt.Println("Mapping follower data from partnerships")
	followers := make([]*Follower, 0)
	for k,v := range followerMap {
		followers = append(followers, mapFollowerData(k,v))
	}

	return followers
}

func mapFollowerData(name string, partnerships []*Partnership) (*Follower) {
	bestPartnership := getBestPartnership(partnerships)

	sort.Slice(partnerships, func(i, j int) bool {
		return partnerships[i].lastCompetition.Before(partnerships[j].lastCompetition)
	})

	follower := Follower {
		name: name,
		partnerships: partnerships,
		bestScore: bestPartnership.highScore,
		bestScorePartnership: bestPartnership,
		lastPartnership: partnerships[len(partnerships)-1],
	}

	return &follower
}

func getBestPartnership(partnerships []*Partnership) (*Partnership) {
	bestPartnership := partnerships[0]
	for _, partnership := range partnerships[1:] {
		if partnership.highScore > bestPartnership.highScore {
			bestPartnership = partnership
		}
	}

	return bestPartnership
}

package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	followerMap := LoadPartnershipData()

	followers := MapFollowers(followerMap)

	filteredFollowers := FilterFollowers(followers,5, 75, 1390, 1700)

	exportResults(filteredFollowers, "./results.txt")
}

func exportResults(followers []*Follower, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, follower := range followers {
		fmt.Fprintln(w, follower)
	}
	return w.Flush()
}
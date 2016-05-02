package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Team struct {
	Name   string
	Cost   int
	Points int
}

type Driver struct {
	Name   string
	Team   Team
	Cost   int
	Points int
}

var mer = Team{"Mercedes", 24, 0}
var fer = Team{"Ferrari", 20, 0}
var wf1 = Team{"Williams", 18, 0}
var rbr = Team{"Red Bull", 14, 0}
var fi = Team{"Force India", 13, 0}
var mcl = Team{"McLaren", 8, 0}
var tor = Team{"Torro Rosso", 7, 0}
var ren = Team{"Renault", 6, 0}
var sau = Team{"Sauber", 4, 0}
var man = Team{"Manor", 4, 0}
var haas = Team{"Haas", 3, 0}

var teams = []Team{mer, fer, rbr, wf1, tor, ren, mcl, fi, sau, man, haas}

var drivers = []Driver{
	Driver{"HAM", mer, 23, 0},
	Driver{"ROS", mer, 20, 0},
	Driver{"VET", fer, 18, 0},
	Driver{"RAI", fer, 13, 0},
	Driver{"BOT", wf1, 16, 0},
	Driver{"MAS", wf1, 14, 0},
	Driver{"RIC", rbr, 14, 0},
	Driver{"KVY", rbr, 11, 0},
	Driver{"HUL", fi, 12, 0},
	Driver{"PER", fi, 11, 0},
	Driver{"ALO", mcl, 10, 0},
	Driver{"BUT", mcl, 10, 0},
	Driver{"VER", tor, 8, 0},
	Driver{"SAI", tor, 6, 0},
	Driver{"MAG", ren, 6, 0},
	Driver{"PAL", ren, 4, 0},
	Driver{"NAS", sau, 4, 0},
	Driver{"ERI", sau, 4, 0},
	Driver{"WEH", man, 3, 0},
	Driver{"HAR", man, 3, 0},
	Driver{"GRO", haas, 5, 0},
	Driver{"GUT", haas, 3, 0},
}

type FF1Team struct {
	Teams   [3]Team
	Drivers [3]Driver
}

func main() {

	updateResults()

	bestFF1Team := FF1Team{}

	totalBudget := 75

	bestPoints := 0
	totalCost := 0

	// Loop over the drivers
	for _, driver1 := range drivers {
		for _, driver2 := range drivers {
			for _, driver3 := range drivers {

				// Ignore duplicates
				if driver1.Name == driver2.Name ||
					driver1.Name == driver3.Name ||
					driver2.Name == driver3.Name {
					break
				}

				// Loop over the teams
				for _, team1 := range teams {
					for _, team2 := range teams {
						for _, team3 := range teams {

							// Ignore duplicates
							if team1.Name == team2.Name ||
								team1.Name == team3.Name ||
								team2.Name == team3.Name {
								break
							}

							// Calculate cost and points
							teamCost := team1.Cost + team2.Cost + team3.Cost
							driverCost := driver1.Cost + driver2.Cost + driver3.Cost

							if teamCost+driverCost < totalBudget {
								teamPoints := team1.Points + team2.Points + team3.Points
								driverPoints := driver1.Points + driver2.Points + driver3.Points

								if teamPoints+driverPoints > bestPoints {
									bestPoints = teamPoints + driverPoints
									bestFF1Team.Teams[0] = team1
									bestFF1Team.Teams[1] = team2
									bestFF1Team.Teams[2] = team3
									bestFF1Team.Drivers[0] = driver1
									bestFF1Team.Drivers[1] = driver2
									bestFF1Team.Drivers[2] = driver3
									totalCost = teamCost + driverCost
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Best combination:")
	fmt.Println("Driver 1:", bestFF1Team.Drivers[0].Name)
	fmt.Println("Driver 2:", bestFF1Team.Drivers[1].Name)
	fmt.Println("Driver 3:", bestFF1Team.Drivers[2].Name)
	fmt.Println("Team 1:", bestFF1Team.Teams[0].Name)
	fmt.Println("Team 2:", bestFF1Team.Teams[1].Name)
	fmt.Println("Team 3:", bestFF1Team.Teams[2].Name)
	fmt.Println("Points:", bestPoints)
	fmt.Println("Cost:", totalCost)
}

func updateResults() {
	raw := readResultsFromFile()
	for _, result := range raw {
		qualiPosition, err := strconv.Atoi(result[3])
		if err != nil {
			log.Fatal("Invalid value for qualiPosition")
		}

		finalPosition, err := strconv.Atoi(result[2])
		if err != nil {
			log.Fatal("Invalid value for finalPosition")
		}

		teamMateQualiPosition, err := strconv.Atoi(result[4])
		if err != nil {
			log.Fatal("Invalid value for teamMateQualiPosition")
		}

		for i, driver := range drivers {
			if driver.Name == result[1] {
				drivers[i].Points += calculatePoints(
					qualiPosition,
					finalPosition,
					teamMateQualiPosition)

				for j, team := range teams {
					if driver.Team.Name == team.Name {
						teams[j].Points +=
							positionToPoints(finalPosition)
						break
					}
				}

				break
			}
		}
	}
}

// CSV: race, driver, quali, finish, team mate quali
func readResultsFromFile() [][]string {
	file, err := os.Open("ff1-results.csv")
	if err != nil {
		log.Fatal("Couldn't open results file", err)
	}

	r := csv.NewReader(file)
	results, err := r.ReadAll()
	if err != nil {
		log.Fatal("Couldn't read results from file", err)
	}

	return results
}

func getDriver(name string) *Driver {
	for _, driver := range drivers {
		if driver.Name == name {
			return &driver
		}
	}

	return nil
}

func calculatePoints(qualiPosition, finalPosition, teamMateQualiPosition int) int {
	points := positionToPoints(finalPosition) +
		positionsGainedPoints(qualiPosition, finalPosition) +
		outQualifyTeamMatePoints(qualiPosition, teamMateQualiPosition)

	if points > 0 {
		return points
	} else {
		return 0
	}
}

func outQualifyTeamMatePoints(qualiPosition, teamMateQualiPosition int) int {
	if qualiPosition < teamMateQualiPosition {
		return 5
	}

	return 0
}

func positionsGainedPoints(gridPosition, finishingPosition int) int {
	if finishingPosition < gridPosition {
		return gridPosition - finishingPosition
	}

	return 0
}

func positionToPoints(position int) int {
	switch position {
	case 1:
		return 25
	case 2:
		return 18
	case 3:
		return 15
	case 4:
		return 12
	case 5:
		return 10
	case 6:
		return 8
	case 7:
		return 6
	case 8:
		return 4
	case 9:
		return 2
	case 10:
		return 1
	}

	return 0
}

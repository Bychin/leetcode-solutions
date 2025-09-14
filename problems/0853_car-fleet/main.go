package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	type car struct {
		position int
		speed    int
	}

	cars := make([]car, 0, len(position))
	for i := 0; i < len(position); i++ {
		cars = append(cars, car{
			position: position[i],
			speed:    speed[i],
		})
	}

	// sort cars from closest to the target to furthest,
	// so the next one may be able to catch the previous one
	sort.Slice(cars, func(i int, j int) bool {
		return cars[i].position > cars[j].position
	})

	fleets := 1
	time := float64(target-cars[0].position) / float64(cars[0].speed)
	for i := 1; i < len(cars); i++ {
		currTime := float64(target-cars[i].position) / float64(cars[i].speed)
		// if ideal time for current car is more than fleet's time
		// it means that it will never catch fleet and we have a new (slower) fleet
		if currTime > time {
			fleets++
			time = currTime
		}
	}
	return fleets
}

func main() {
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})) // 3
	fmt.Println(carFleet(10, []int{3}, []int{3}))                          // 1
	fmt.Println(carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))             // 1
}

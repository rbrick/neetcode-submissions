

type car struct {
	position int
	speed    int
}

func (c *car) String() string {
	return fmt.Sprintf("{position: %d, speed: %d}", c.position, c.speed)
}

func carFleet(target int, position []int, speed []int) int {

	cars := make([]*car, len(position))

	for i := 0; i < len(position); i++ {
		cars[i] = &car{
			position: position[i],
			speed:    speed[i],
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	time := func(speed, distance int) float64 {
		return float64(distance) / float64(speed)
	}

	fleets := 0
	lastArrival := 0.0

	for _, c := range cars {
		t := time(c.speed, target-c.position)

		if t > lastArrival {
			fleets++
			lastArrival = t
		}

	}

	return fleets
}
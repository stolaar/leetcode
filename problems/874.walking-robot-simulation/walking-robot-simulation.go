package main

type robot struct {
	Direction int
	X         int
	Y         int
	Obstacles map[int]map[int]bool
	Max       int
}

func (r *robot) move(command int) {
	noobs := len(r.Obstacles) == 0
	if r.Direction == 0 {
		if noobs {
			r.Y += command
			return
		}
		current := r.Y
		s := current
		for s < current+command {
			s += 1
			r.Y = s
			if _, ok := r.Obstacles[r.X][s]; ok {
				if r.Y != current {
					r.Y -= 1
				}
				return
			}
		}
	}

	if r.Direction == 2 {
		if noobs {
			r.Y -= command
			return
		}
		current := r.Y
		s := current
		for s > current-command {
			s -= 1
			r.Y = s
			if _, ok := r.Obstacles[r.X][s]; ok {
				if r.Y != current {
					r.Y += 1
				}
				return
			}
		}
	}

	if r.Direction == 1 {
		if noobs {
			r.X += command
			return
		}
		current := r.X
		s := current

		for s < current+command {
			s += 1
			r.X = s
			if _, ok := r.Obstacles[s][r.Y]; ok {
				if r.X != current {
					r.X -= 1
				}
				return
			}
		}
	}

	if r.Direction == 3 {
		if noobs {
			r.X -= command
			return
		}
		current := r.X
		s := current

		for s > current-command {
			s -= 1
			r.X = s
			if _, ok := r.Obstacles[s][r.Y]; ok {
				if r.X != current {
					r.X += 1
				}
				return
			}
		}
	}
}

func (r *robot) changeDirection(command int) {
	if command == -2 {
		if r.Direction == 0 {
			r.Direction = 3
			return
		}
		r.Direction -= 1
		return
	}
	if r.Direction == 3 {
		r.Direction = 0
		return
	}
	r.Direction += 1
}

func (r *robot) processCommand(command int) {
	if command == -2 || command == -1 {
		r.changeDirection(command)
		return
	}

	r.move(command)

	sqrsum := (r.X * r.X) + (r.Y * r.Y)
	if sqrsum > r.Max {
		r.Max = sqrsum
	}
}

func robotSim(commands []int, obstacles [][]int) int {
	obs := make(map[int]map[int]bool)

	for _, o := range obstacles {
		if _, ok := obs[o[0]]; ok {
			obs[o[0]][o[1]] = true
			continue
		}
		obs[o[0]] = map[int]bool{}
		obs[o[0]][o[1]] = true
	}

	r := robot{
		Direction: 0,
		X:         0,
		Y:         0,
		Obstacles: obs,
		Max:       0,
	}

	for _, command := range commands {
		r.processCommand(command)
	}

	return r.Max
}

func main() {
	println(robotSim([]int{4, -1, 3}, [][]int{}))
	println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
	println(robotSim([]int{4, -2, 4, -1, 4}, [][]int{{2, 4}}))
	println(robotSim([]int{6, -1, -1, 6}, [][]int{}))
	println(robotSim([]int{-2, -1, 8, 9, 6}, [][]int{
		{-1, 3},
		{0, 1},
		{-1, 5},
		{-2, -4},
		{5, 4},
		{-2, -3},
		{5, -1},
		{1, -1},
		{5, 5},
		{5, 2},
	}))
	println(robotSim([]int{-2, -1, -2, 3, 7},
		[][]int{
			{1, -3},
			{2, -3},
			{4, 0},
			{-2, 5},
			{-5, 2},
			{0, 0},
			{4, -4},
			{-2, -5},
			{-1, -2},
			{0, 2},
		}))
}

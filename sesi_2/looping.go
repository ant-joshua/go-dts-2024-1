package main

// firstWayLoop lebih mirip dengan for loop pada umumnya
func firstWayLoop() {

	index := 0

	for i := index; i < 5; i++ {
		println(i)
	}
}

// secondWayLoop lebih mirip dengan while loop
func secondWayLoop() {
	index := 0

	for index < 5 {
		println(index)
		index++
	}
}

func thirdWayLoop() {
	index := 0

	for {
		println(index)
		index++

		if index == 5 {
			break
		}
	}
}

func loopBreakAndContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}

		if i == 8 {
			break
		}

		println(i)
	}
}

func nestedLoop() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			println(i, j)
		}
	}
}

func nestedLoopWithParam(total int) {
	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			println(i, j)
		}
	}
}

func labeledLoop() {
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			println(i, j)

			if i == 2 {
				break outerLoop
			}
		}
	}
}

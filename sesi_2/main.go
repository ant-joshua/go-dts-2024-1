package main

func main() {
	temporaryVariable()
}

func temporaryVariable() {
	currentYear := 2024

	if age := currentYear - 1999; age > 25 {
		println("You are old")
	} else {
		println("You are young")
	}
}

func switchExample() {
	score := 5

	switch score {
	case 1:
		println("You are bad")
	case 2:
		println("You are not good")
	case 3:
		println("You are average")
	case 4:
		println("You are good")
	}
}

func switchRelationExample() {
	score := 5

	switch {
	case score < 3:
		println("You are bad")
	case score == 3:
		println("You are average")
	case score > 3 && score < 5:
		println("You are good")
	default:
		println("You are excellent")
	}
}

func switchFallthroughExample() {
	score := 6

	switch {
	case score == 8:
		println("You are excellent")

	case (score < 8) && (score > 5):
		println("You are average")
		fallthrough
	case score > 3 && score < 5:
		println("You are good")
	}

}

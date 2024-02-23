package main

import ("fmt" "net/http")

// func main() {
// 	r := gin.Default()

// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	r.Run()

// 	// InitDB()
// }

func main() {
	// nama := "Joshua"
	// var nilai int = 10

	// var nilai1, nilai2, nilai3 int = 10, 20, 30

	// nilai1 = 100

	// fmt.Println("nama saya", nama)
	// fmt.Println("nilai saya", nilai)

	// fmt.Printf("tipe data nama = %T\n", nama)
	// fmt.Printf("tipe data int = %T\n", nilai)

	// // fmt.Printf("value dari nama = %d\n", nama)
	// fmt.Printf("value dari nilai = %v\n", nilai)

	// fmt.Printf("value dari nilai1 = %v\n", nilai1)

	// nilai_total, err := strconv.Atoi("nilai 10")

	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }

	// fmt.Printf("value dari nilai_total = %v\n", nilai_total)

	const (
		c1 = iota - 2
		c2
		c3
	)

	fmt.Println(c1, c2, c3)

	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	type Gender int

	const (
		Male Gender = iota
		Female
	)

	fmt.Println("Gender Male = ", Male, "Gender Femaile = ", Female)

	for day := Sunday; day <= Saturday; day++ {
		fmt.Println(day)
	}

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	if 1 == 2 {
		fmt.Println("1 lebih besar dari 2")
	} else {
		fmt.Println("1 lebih kecil dari 2")
	}

	var nilai int = 70

	if nilai >= 70 && nilai <= 80 {
		fmt.Println("nilai anda B-")
	}

}

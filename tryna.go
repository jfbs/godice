package main

import ("fmt"; "math/rand"; "time"; "strconv";)

func main() {
	for k := 0; k < 6; k++ {

		fmt.Println(rolldice())
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func rolldice() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var fivedice string = ""

	for j := 0; j < 5; j++ {
		fivedice += strconv.Itoa(randInt(1,7))		
	}

	return fivedice

}


/* for i, value := range favNums {

} */

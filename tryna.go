package main

import ("fmt"; "math/rand"; "time"; "strconv";)

func main() {

	rollstr := [6]string{}

	for k := 0; k < 6; k++ {
		rollstr[k] += rolldice()
		//fmt.Println(rolldice())
	}
	for i := 0; i < 6; i++ {
		fmt.Println(rollstr[i])
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func rolldice() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var fivedice string = ""

	for j := 0; j < 5; j++ { // build five six-sided dice roll
		fivedice += strconv.Itoa(randInt(1,7))		
	}

	return fivedice

}


/* for i, value := range favNums {

} */

package main

import ("fmt"; "math/rand"; "time"; "strconv";)

func main() {
	getsix() 
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

func getsix() {
	rollstr := [6]string{}

	for k := 0; k < len(rollstr); k++ {
		rollstr[k] += rolldice()
		fmt.Println("roll ",k+1, ": ", rolldice())
	}
}

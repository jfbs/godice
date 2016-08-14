package main

import ("fmt"; "math/rand"; "time"; "strconv";)

func main() {

	fmt.Println(rollStr())
	
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func rolldice() string{
	rand.Seed(time.Now().UTC().UnixNano())
	var fivex string = ""
	rollStr := [1]string{}

	for j := 0; j < 5; j++ {
		fivex += strconv.Itoa(randInt(1,7))
		
	}
	// fmt.Println(fivex)
	rollStr[0] = fivex
	//fmt.Println(rollStr[0])
	return rollStr[0]

}


/* for i, value := range favNums {

} */

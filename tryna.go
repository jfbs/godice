package main

import ("fmt"; "math/rand"; "time"; "strconv";)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var fivex string = ""
	rollStr := [6]string{}

	for j := 0; j < 5; j++ {
		// fmt.Printf("%d", randInt(1,7))
		fivex += strconv.Itoa(randInt(1,7))
		
	}
	// fmt.Println(fivex)
	rollStr[0] = fivex
	fmt.Println(rollStr[0])
 
	
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

/* func rolldice() string{


} */


/* for i, value := range favNums {

} */

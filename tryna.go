package main

import ("fmt"; "math/rand"; "time"; "strconv"; "os"; "bufio"; "strings";)

func main() {
	sixrolls := []string{}
	sixrolls =  getsix()
	/* for v:= range sixrolls {
		fmt.Println(sixrolls[v])  // print each hand
	} */

	file, err := os.Open("diceware.txt")
	if err != nil {
		checkError(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// delimiter := '\t'
	// vartemp := scanner.Text()
	// rightofdel := vartemp.Join(strings.Split(vartemp, delimiter)[1:], delimiter)

	for scanner.Scan() {
		for d := range sixrolls {
			if strings.Contains(scanner.Text(), sixrolls[d]) {
				// varg = sixrolls[d]
				fmt.Printf("Roll %d is: %v\n", d, scanner.Text())
			}
		}
	}

	fmt.Println("first is: ", sixrolls[0]) // check first roll

	if err := scanner.Err(); err != nil {
		checkError(err)
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max - min)
}

func rolldice() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var fivedice string

	for j := 0; j < 5; j++ { // build five six-sided dice roll
		fivedice += strconv.Itoa(randInt(1,7))
	}

	return fivedice
}

func getsix() []string {
	rollstr := []string{}

	for k := 0; k < 6; k++ { // roll five dice six times
		rollstr = append(rollstr, rolldice())
	}

	return rollstr
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

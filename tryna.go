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

	varg := make([]string, 6)
	for scanner.Scan() {
		for d := 0; d < 6; d++ {
			if strings.Contains(scanner.Text(), sixrolls[d]) {
				varg[d] = scanner.Text()
				//varg = append(varg, scanner.Text())
				fmt.Printf("Roll %d is: %v\n", d, scanner.Text())
			}
		}
	}

	fmt.Println("first is: ", sixrolls[0]) // check first roll
	fmt.Println("second is: ", sixrolls[1])
	fmt.Println("third is: ", sixrolls[2])

	fmt.Println("first varg is: ", varg[0])
	fmt.Println("second varg is: ", varg[1])
	fmt.Println("third varg is: ", varg[2])
  for e := range varg {
		fmt.Printf("Roll %d is: %v\n", e, varg[e])
	}
	//fmt.Println(varg)

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

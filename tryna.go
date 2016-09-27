package main

import ("fmt"; "math/rand"; "time"; "strconv"; "os"; "bufio"; "strings";)

const phrasenum int = 8 // change for desired phrase count

func main() {
	xphrase := []string{}
	xphrase =  getrolls()

	file, err := os.Open("diceware.txt")
	if err != nil {
		checkError(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	varg := make([]string, phrasenum) // store matches into slice
	for scanner.Scan() {
		for d := range xphrase {
			if strings.Contains(scanner.Text(), xphrase[d]) {
				varg[d] = scanner.Text()
			}
		}
	}

  for e := range varg {
		fmt.Printf("Roll %d is: %v\n", e+1, varg[e])
	}

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

	for j := 0; j < 5; j++ { // roll a six-sided die five times
		fivedice += strconv.Itoa(randInt(1,7))
	}

	return fivedice
}

func getrolls() []string {
	rollstr := []string{}

	for k := 0; k < phrasenum; k++ { // roll five dice for each phrase
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

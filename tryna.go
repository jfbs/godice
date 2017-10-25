package main

import ("fmt"; "math/rand"; "time"; "strconv"; "os"; "bufio"; "strings";)

const number_of_phrases int = 12 // change for desired phrase count

func main() {
  start := time.Now()
  rolled_hand := []string{}
  rolled_hand =  getrolls()

  file, err := os.Open("diceware.txt")
  if err != nil {
    checkError(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  phrase_id := make([]string, number_of_phrases) // store matches into slice
  for scanner.Scan() {
    for d := range rolled_hand {
      if strings.Contains(scanner.Text(), rolled_hand[d]) {
        phrase_id[d] = scanner.Text() // match
      }
    }
  }

  var whole_phrase string
  for e := range phrase_id {
    input := string(phrase_id[e])
    trimmed := strings.Replace(input, "\t"," ",-1)
    parts := strings.Split(trimmed, " ")
    whole_phrase += parts[1] + " "
  }

  trim_phrase := strings.Trim(whole_phrase, " ")
  fmt.Println("Passphrase: ", trim_phrase)

  if err := scanner.Err(); err != nil {
    checkError(err)
  }

  elapsed := time.Since(start)
  fmt.Println("Runtime: ", elapsed)
}

func randInt(min int, max int) int {
  rand.Seed(time.Now().UTC().UnixNano())
  return min + rand.Intn(max - min)
}

func rolldice() string {
  var fivedice string

  for j := 0; j < 5; j++ { // roll a six-sided die five times
    fivedice += strconv.Itoa(randInt(1,7))
  }

  return fivedice
}

func getrolls() []string {
  rollstr := []string{}

  for k := 0; k < number_of_phrases; k++ { // roll five dice for each phrase
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

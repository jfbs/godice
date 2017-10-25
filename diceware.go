package main

import ("fmt"; "math/rand"; "time"; "strconv"; "os"; "bufio"; "strings";)

const phrase_max int = 12 // change for desired phrase count

func main() {
  rolled_hand := []string{}
  rolled_hand =  getrolls()

  file, err := os.Open("diceware.txt")
  checkError(err)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  phrase_id := make([]string, phrase_max) // store matches into slice
  for scanner.Scan() {
    checkError(err)
    for d := range rolled_hand {
      if strings.Contains(scanner.Text(), rolled_hand[d]) {
        phrase_id[d] = scanner.Text() // match
      }
    }
  }

  // Split dice from match and parse passphrases into a string
  var whole_phrase string
  for e := range phrase_id {
    match := string(phrase_id[e])
    trimmed := strings.Replace(match, "\t"," ",-1)
    parts := strings.Split(trimmed, " ")
    whole_phrase += parts[1] + " "
  }

  trim_phrase := strings.Trim(whole_phrase, " ")
  fmt.Println("Passphrase: ", trim_phrase)
}

func randInt(min int, max int) int {
  rand.Seed(time.Now().UTC().UnixNano())
  return min + rand.Intn(max - min)
}

// Return random values from 1 through 6 five times.
func rolldice() string {
  var fivedice string

  for j := 0; j < 5; j++ {
    fivedice += strconv.Itoa(randInt(1,7))
  }

  return fivedice
}

// Return five dice for each phrase
func getrolls() []string {
  rollstr := []string{}

  for k := 0; k < phrase_max; k++ {
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

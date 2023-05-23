package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	//Declare variables
	i := 1
	var score int

	//Declare flag
	filePtr := flag.String("filename", "problems", "a string")
	timePtr := flag.Int("time", 30, "the time limit for the quiz in seconds")
	shufflePtr := flag.Bool("shuffle", false, "the shuffle function")

	//Parse flag
	flag.Parse()

	//Start the timer
	timer := time.NewTimer(time.Second * time.Duration(*timePtr))

	filepath := "C:\\GolandProjects\\Quiz_game\\Quizes\\" + *filePtr + ".csv"

	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	// Close file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Create new CSV Reader
	reader := csv.NewReader(file)

	//Read all lines
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	//Shuffle the slice
	if *shufflePtr {
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	//Main loop
	for _, record := range records {
		fmt.Printf("%v# task: %v = ", i, record[0])
		answerCh := make(chan string)
		go func() {
			var ans string
			_, err = fmt.Scan(&ans)
			if err != nil {
				panic(err)
			}
			ans = strings.TrimSpace(ans)
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nThe score is: %v out of %v\n", score, len(records))
			return
		case ans := <-answerCh:
			if strings.EqualFold(ans, record[1]) {
				score++
			}
			i++
		}
	}

	fmt.Printf("\nThe score is: %v out of %v\n", score, len(records))
}

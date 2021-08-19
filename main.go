package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	csvFile := flag.String("csv","problems.csv","a csv file in the format of problems")

	timeLimit := flag.Int("limit",30,"the time limit for answering questions")
	flag.Parse()

	file,err := os.Open(*csvFile)
	if err != nil{
		exit(fmt.Sprintf("faild to load %s",*csvFile))		
	}
	r := csv.NewReader(file)
	lines,err := r.ReadAll()
	if err != nil{
		exit(fmt.Sprintf("faild to read %s",*csvFile))		
	}
	problems := parsToSlice(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	for i,p := range problems {
		fmt.Printf("Problem #%d: %s= \n",i+1 , p.question)
		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerCh <- answer
		}()
		select {
		case <- timer.C:
			fmt.Printf("You scored %d out of %d \n",correct,len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d \n",correct,len(problems))
}

func parsToSlice(lins [][]string) []problems{
	ret := make([]problems,len(lins))
	for i,line := range lins{
		ret[i] = problems{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
} 

type problems struct {
	question string
	answer string
}

func exit(msg string){
	fmt.Printf("faild to load %s",msg)
	os.Exit(1)
}
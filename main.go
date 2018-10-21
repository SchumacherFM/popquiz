package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {
	color.White("Welcome to the Go pop quiz!\n")

	quizzes := []quiz{
		{
			`What does this program print?
func main() {
	one, two, three := 0.1, 0.2, 0.3
	fmt.Println(one+two > three)
}`,
			[]string{"true", "false", "0.3"},
			func() (string, string) {
				one, two, three := 0.1, 0.2, 0.3
				return fmt.Sprintf("%t", (one+two) > three),
					fmt.Sprintf("Because: %s > %s", strconv.FormatFloat(one+two, 'f', -1, 64), strconv.FormatFloat(three, 'f', -1, 64))
			},
			"https://twitter.com/davecheney/status/1053490060492890118",
		},
		{
			`What does this program print?
func main() {
	var a []int
	b := a[:]
	fmt.Println(b == nil)
}`,
			[]string{"true", "false", "panic on line 4"},
			func() (string, string) {
				var a []int
				b := a[:]
				return fmt.Sprintf("%t", b == nil), ""
			},
			"https://twitter.com/davecheney/status/1053419185680744448",
		},
		{
			`Does this program compile?
func main() {
    for i := 1
    	i < 10
    	i++ {
    	fmt.Println("Hello!")
    }
}`,
			[]string{"yes", "no", "only on windows"},
			func() (string, string) {
				for i := 1; i < 10; i++ {
					fmt.Println("Hello!")
				}
				return "yes", ""
			},
			"https://twitter.com/davecheney/status/1042396099443453952",
		},
		{
			`What does this program print?
func main() {
	input := []byte("Hello, playground")
	hash := sha1.Sum(input)[:5]
	fmt.Println(hash)
}`,
			[]string{"461ec8144f", "[70 30 200 20 79]", "nothing, doesn't compile"},
			func() (string, string) {
				// input := []byte("Hello, playground")
				// hash := sha1.Sum(input)[:5]
				// fmt.Println(hash)
				return "nothing, doesn't compile", "slice of unaddressable value"
			},
			"https://twitter.com/davecheney/status/1041526653141147654",
		},
		{
			`What does this program print?
func main() {
	fmt.Println(string('7'<<1))
}`,
			[]string{"7", "14", "n", "doesn't compile"},
			func() (string, string) {
				return string('7' << 1), ""
			},
			"https://twitter.com/davecheney/status/1039997464361623552",
		},
	}
	points := 0
	for _, quz := range quizzes {

		prompt := &survey.Select{
			Message: quz.question,
			Options: quz.answers,
		}
		var answer string
		survey.AskOne(prompt, &answer, nil)
		if result, explanation := quz.fn(); answer == result {
			color.Green("Ok\n")
			points++
		} else {
			color.Red("Have: %q\nWant: %q\n%s\n", answer, result, explanation)
		}
		color.Yellow("read more: %s", quz.url)
		fmt.Println("\n")
	}
	fmt.Printf("Your points %d out of possible %d.\n", points, len(quizzes))
}

type quiz struct {
	question string
	answers  []string
	fn       func() (answer string, explanation string)
	url      string
}

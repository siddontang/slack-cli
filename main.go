package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"regexp"
	"strings"
)

var token = flag.String("token", "", "Slack Token")

type Slack struct {
	s *slack.Slack
}

func main() {
	flag.Parse()

	s := &Slack{}
	s.s = slack.New(*token)

	SetCompletionHandler(completionHandler)
	setHistoryCapacity(100)

	reg, _ := regexp.Compile(`([^\s]*"[^"]+"[^\s]*)|([^\s]*'[^']+'[^\s]*)|[^"]?\S+[^"]?|[^']?\S+[^']?`)

	prompt := "slack>"

	for {

		cmd, err := line(prompt)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		cmds := reg.FindAllString(cmd, -1)
		for i, _ := range cmds {
			cmds[i] = strings.TrimSpace(cmds[i])
		}

		if len(cmds) == 0 {
			continue
		} else {
			addHistory(cmd)

			args := cmds[1:]

			cmd := strings.ToLower(cmds[0])
			if cmd == "help" || cmd == "?" {
				printHelp(cmds)
			} else {
				v, err := s.handle(cmd, args)
				if err != nil {
					fmt.Printf("err: %s", err.Error())
				} else if v != nil {
					buf, _ := json.MarshalIndent(v, "", "    ")
					fmt.Printf("%s", buf)
				} else {
					fmt.Printf("ok")
				}

				fmt.Printf("\n")
			}

		}
	}
}

func printGenericHelp() {
	msg :=
		`stack-cli
Type:	"help <command>" for help on <command>
	`
	fmt.Println(msg)
}

func printCommandHelp(arr []string) {
	fmt.Println()
	fmt.Printf("\t%s %s \n", arr[0], arr[1])
	fmt.Printf("\tDescription: %s", arr[2])
	fmt.Println()
}

func printHelp(cmds []string) {
	args := cmds[1:]
	if len(args) == 0 {
		printGenericHelp()
	} else if len(args) > 1 {
		fmt.Println()
	} else {
		cmd := args[0]
		for i := 0; i < len(helpCommands); i++ {
			if strings.EqualFold(helpCommands[i][0], cmd) {
				printCommandHelp(helpCommands[i])
			}
		}
	}
}

func completionHandler(in string) []string {
	var keyWords []string
	for _, i := range helpCommands {
		if strings.HasPrefix(strings.ToUpper(i[0]), strings.ToUpper(in)) {
			keyWords = append(keyWords, i[0])
		}
	}
	return keyWords
}

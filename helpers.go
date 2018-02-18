package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func setLogFiles(path string) {
	if path != "" {
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		} else {
			defer f.Close()
			log.SetOutput(f)
		}
	}
}

func getUsers(users string) []int {
	userStrings := strings.Split(users, ",")
	userInt := make([]int, 0)

	for _, s := range userStrings {
		u, err := strconv.Atoi(s)
		if err == nil {
			userInt = append(userInt, u)
		}
	}

	return userInt
}

func isAllowedID(id int, users []int) bool {
	if len(users) == 0 {
		return true
	}

	for _, v := range users {
		if v == id {
			return true
		}
	}

	return false
}

func isAllowedUsername(name string, names []string) bool {
	if len(names) == 0 {
		return true
	}

	if name == "" {
		return false
	}

	for _, v := range names {
		if v == name {
			return true
		}
	}

	return false
}

func (args ARGs) validate() error {
	helpMessage := "use --help for help"

	if *args.Token == "" {
		return fmt.Errorf("telegram token is required, %s", helpMessage)
	}

	if *args.Lang == "" {
		return fmt.Errorf("lang is required, %s", helpMessage)
	}

	if *args.DialogFlowToken == "" {
		return fmt.Errorf("dialog flow token is required, %s", helpMessage)
	}

	return nil
}

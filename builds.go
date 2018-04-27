package main

import (
	"bufio"
	"log"
	"os"
)

func getBuildList(txtFile string) []string {
	var builds []string

	file, err := os.Open(txtFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builds = append(builds, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return builds
}

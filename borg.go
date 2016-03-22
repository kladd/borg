package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kladd/borg/org"
)

func fname(t time.Time) string {
	return t.Format("20060102.org")
}

func main() {
	t := time.Now()
	y := t.Add(-1 * 24 * time.Hour)

	tfn := fname(t)
	if _, err := os.Stat(tfn); err == nil {
		log.Println(fmt.Sprintf("%s exists already", tfn))
		os.Exit(1)
	}

	output := fmt.Sprintf(
		"#+TITLE: %s\n\n",
		fmt.Sprintf(
			"%s %d",
			strings.ToUpper(t.Month().String())[:3],
			t.Day(),
		),
	)

	yfn := fname(y)
	if _, err := os.Stat(yfn); err == nil {
		log.Println(fmt.Sprintf("migrating %s", yfn))
		output += org.ExtractRemaining(yfn)
	}

	ioutil.WriteFile(tfn, []byte(output), 0644)
}

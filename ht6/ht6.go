package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].birthDay.Equal(p[j].birthDay) {
		return (p[i].firstName + p[i].lastName) < (p[j].firstName + p[j].lastName)
	}
	return p[i].birthDay.Before(p[j].birthDay)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	ivanIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	petrPetrovDate, err := time.Parse("2006-Jan-02", "2004-Aug-11")
	vasyaVasinDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	alekseyAlekseevDate, err := time.Parse("2006-Jan-02", "2004-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate},
		{"Petr", "Petrov", petrPetrovDate},
		{"Vasya", "Vasin", vasyaVasinDate},
		{"Aleksey", "Alekseev", alekseyAlekseevDate},
	}
	sort.Sort(p)
	fmt.Println(p)
}

package main

import "fmt"

type Age int

const (
	MAX_BABY_AGE   Age = 3
	MAX_CHILD_AGE  Age = 10
	MAX_JUNIOR_AGE Age = 16
	MAX_YOUNG_AGE  Age = 35
	MAX_MIDDLE_AGE Age = 60
	MAX_OLD_AGE    Age = 150
)

func (self Age) Group() string {
	self--
	switch {
	case self <= MAX_BABY_AGE:
		fmt.Println("Baby in ", MAX_BABY_AGE)
		return "Baby"
	case self <= MAX_CHILD_AGE:
		fmt.Println("CHILD in ", MAX_CHILD_AGE)
		return "Children"
	case self <= MAX_OLD_AGE:
		fmt.Println("People in ", MAX_OLD_AGE)
		return "People"
	}
	return "unknown"
}

func (self *Age) WhichGroup() string {
	*self--
	switch {
	case *self <= MAX_BABY_AGE:
		fmt.Println("Baby in ", MAX_BABY_AGE)
		return "Baby"
	case *self <= MAX_CHILD_AGE:
		fmt.Println("CHILD in ", MAX_CHILD_AGE)
		return "Children"
	case *self <= MAX_OLD_AGE:
		fmt.Println("People in ", MAX_OLD_AGE)
		return "People"
	}
	return "unknown"
}

func main() {
	var mike Age = 3
	var lucy Age = 19
	fmt.Println("mike is ", mike.Group(), "age ", mike)
	fmt.Println("lucy is ", lucy.WhichGroup(), "age ", lucy)

}

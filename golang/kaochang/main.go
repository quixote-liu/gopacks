package main

import "fmt"

func main() {
	er := Constructor(10)
	fmt.Printf("[1]er.Seat(): %v\n", er.Seat())
	fmt.Printf("[2]er.Seat(): %v\n", er.Seat())
	fmt.Printf("[3]er.Seat(): %v\n", er.Seat())
	fmt.Printf("[3]er.Seat(): %v\n", er.Seat())
	er.Leave(0)
	er.Leave(4)
	fmt.Printf("[4]er.Seat(): %v\n", er.Seat())
}

type ExamRoom struct {
	seats  map[int]bool
	length int
}

func Constructor(n int) ExamRoom {
	seats := make(map[int]bool)
	for i := 0; i < n; i++ {
		seats[i] = false
	}
	return ExamRoom{
		seats:  seats,
		length: n,
	}
}

func (er *ExamRoom) Seat() int {
	emptySeatIndexs := []int{}
	for i, sit := range er.seats {
		if !sit {
			emptySeatIndexs = append(emptySeatIndexs, i)
		}
	}

	if len(emptySeatIndexs) == er.length {
		er.seats[0] = true
		return 0
	}

	var s int
	var leng int
	for _, i := range emptySeatIndexs {
		seatMinLen := 0
		for m := i - 1; m >= 0; m-- {
			if er.seats[m] {
				seatMinLen = i - m
				break
			}
		}

		for n := i + 1; n < er.length; n++ {
			if er.seats[n] {
				rightLen := n - i
				if seatMinLen > rightLen {
					seatMinLen = rightLen
				}
				break
			}
		}

		if seatMinLen > leng {
			s = i
			leng = seatMinLen
		}
	}

	er.seats[s] = true
	return s
}

func (er *ExamRoom) Leave(p int) {
	er.seats[p] = false
}

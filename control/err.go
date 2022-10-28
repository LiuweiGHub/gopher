package main

import (
	"errors"
	"fmt"
)

type myErr struct {
	code int
	msg  string
	err  error
}

var (
	// 哨兵
	ErrResourceFull = errors.New("bufi0: err full")
)

func main() {
	err := errors.New("bufi0: err full")
	if errors.Is(err, ErrResourceFull) {
		fmt.Errorf("hahahahahahahah~~~~~~~~~%s", "hahahahahaha")
		fmt.Println(fmt.Errorf("hahahahahahahah~~~~~~~~~%s", "hahahahahaha").Error())
	}
	fmt.Println(fmt.Errorf("hahahahahahahah~~~~~~~~~%s", "hahahahahaha").Error())
}

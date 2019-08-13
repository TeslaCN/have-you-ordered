package test

import (
	"fmt"
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	fmt.Println("Start func")
	//panic("error")
	var a *int
	*a = 1
	fmt.Println("end func")
}

func TestRecover(t *testing.T) {
	err := recover()
	fmt.Printf("%s\n", err)
	fmt.Println("Start func")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Defering %s\n%T\n", err, err)
		}
	}()
	//panic("error")
	var a *int
	*a = 1
	fmt.Println("end func")
}

func TestCatchException(t *testing.T) {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			func() {
				defer func() {
					if e := recover(); e != nil {
						fmt.Println(e)
					}
				}()
				panic("I'm panic.")
			}()
		}
	}
}

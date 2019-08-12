package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	layout := "20060102"
	for moment, _ := time.Parse(layout, "20180101"); moment.Before(time.Now()); moment = moment.Add(time.Hour * 24) {
		fmt.Println(moment.Format(layout))
	}
}

func TestTimeType(t *testing.T) {
	var t0 time.Time
	fmt.Println(t0.UnixNano())
	t0 = time.Now()
	fmt.Println(t0.UnixNano())
}

func TestTimerChan(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Timer started")
	<-timer.C
	fmt.Println(t)
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(250 * time.Millisecond)
	timer := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-timer.C:
			fmt.Println("done")
			goto out
		}
	}
out:
	fmt.Println("Exit")
}

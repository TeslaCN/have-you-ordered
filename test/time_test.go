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

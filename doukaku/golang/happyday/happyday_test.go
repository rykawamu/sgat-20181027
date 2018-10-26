package main

import (
    "testing"
    "time"
)

func TestCreateStartday(t *testing.T) {
    t1 := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
    t2 := createStartday("2015-01-01")

    if !t1.Equal(t2) {
        t.Fatalf("failed test t1:%v t2:%v", t1, t2)
    }
}

func TestSet(t *testing.T) {
    t1 := time.Date(2015, time.November, 12, 0, 0, 0, 0, time.UTC)
    var ymd StructYmd
    ymd = set(t1)

    if ymd.Year != 2015 || ymd.Month != 11 || ymd.Day != 12 {
         t.Fatalf("failed test t1:%v  : [%v-%v-%v]", t1, ymd.Year, ymd.Month, ymd.Day)
    }
}

func TestCheck(t *testing.T) {
    var ymd StructYmd
    var happy bool

    // Happy Day
    ymd.Year = 2015; ymd.Month = 4; ymd.Day = 22
    happy = ymd.check()
    if !happy {
         t.Fatalf("failed test [%v-%v-%v][%v]", ymd.Year, ymd.Month, ymd.Day, happy)
    }

    ymd.Year = 2015; ymd.Month = 4; ymd.Day = 4
    happy = ymd.check()
    if !happy {
         t.Fatalf("failed test [%v-%v-%v][%v]", ymd.Year, ymd.Month, ymd.Day, happy)
    }

    ymd.Year = 2015; ymd.Month = 10; ymd.Day = 19
    happy = ymd.check()
    if !happy {
         t.Fatalf("failed test [%v-%v-%v][%v]", ymd.Year, ymd.Month, ymd.Day, happy)
    }

    // Not Happy Day
    ymd.Year = 2015; ymd.Month = 12; ymd.Day = 3
    happy = ymd.check()
    if happy {
         t.Fatalf("failed test [%v-%v-%v][%v]", ymd.Year, ymd.Month, ymd.Day, happy)
    }

    ymd.Year = 2015; ymd.Month = 11; ymd.Day = 28
    happy = ymd.check()
    if happy {
         t.Fatalf("failed test [%v-%v-%v][%v]", ymd.Year, ymd.Month, ymd.Day, happy)
    }

}


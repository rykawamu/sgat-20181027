package main

import (
    "testing"
)

func TestCheckMod(t *testing.T) {
    var flag bool
    var a, n int

    a = 3; n = 3
    flag = checkMod(a, n)
    if flag != true {
        t.Fatalf("failes test a:%v n:%v", a, n)
    }

    a = 9; n = 3
    flag = checkMod(a, n)
    if flag != true {
        t.Fatalf("failes test a:%v n:%v", a, n)
    }

    a = 32; n = 3
    flag = checkMod(a, n)
    if flag != false {
        t.Fatalf("failes test a:%v n:%v", a, n)
    }

    a = 53; n = 3
    flag = checkMod(a, n)
    if flag != false {
        t.Fatalf("failes test a:%v n:%v", a, n)
    }

    a = 5; n = 5
    flag = checkMod(a, n)
    if flag != true {
        t.Fatalf("failes test a:%v n:%v", a, n)
    }

    a = 20; n = 5
    flag = checkMod(a, n)
    if flag != true {
        t.Fatalf("failes test a:%v n:%v", a, n)
    }

    a = 52; n = 5
    flag = checkMod(a, n)
    if flag != false {
        t.Fatalf("failes test a:%v n:%v", a, n)
    }

    a = 53; n = 5
    flag = checkMod(a, n)
    if flag != false {
        t.Fatalf("failes test a:%v n:%v", a, n)
    }
}


func TestSameStr(t *testing.T) {
    var flag bool
    var src, dist string 

    src = "1"; dist = "1"
    flag = sameStr(src, dist)
    if flag != true {
        t.Fatalf("failes test src:%v dist:%v", src, dist)
    }

    src = "2"; dist = "1"
    flag = sameStr(src, dist)
    if flag != false {
        t.Fatalf("failes test src:%v dist:%v", src, dist)
    }
}


func TestCheck(t *testing.T) {
    var flag bool
    var i, target int

    i = 3; target = 3
    flag = check(i, target)
    if flag != true {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 9; target = 3
    flag = check(i, target)
    if flag != true {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 32; target = 3
    flag = check(i, target)
    if flag != true {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 53; target = 3
    flag = check(i, target)
    if flag != true {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 47; target = 3
    flag = check(i, target)
    if flag != false {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 5; target = 5
    flag = check(i, target)
    if flag != true {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 20; target = 5
    flag = check(i, target)
    if flag != true {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 52; target = 5
    flag = check(i, target)
    if flag != true {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 53; target = 5
    flag = check(i, target)
    if flag != true {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

    i = 47; target = 5
    flag = check(i, target)
    if flag != false {
        t.Fatalf("failes test i:%v target:%v", i, target)
    }

}


func TestFizzBuzzPlus(t *testing.T) {
    var m = map[int]string {1: "1", 3: "Fizz", 5: "Buzz", 13: "Fizz", 15: "FizzBuzz", 32: "Fizz", 35: "FizzBuzz", 52: "Buzz", 53: "FizzBuzz", 47: "47", 62: "62", 91: "91"} 

    for k, v := range m {
        result := fizzBuzzPlus(k)
        if result != v {
            t.Fatalf("failes test k:%v v:%v  result:%v", k, v, result)
        }
    }
}

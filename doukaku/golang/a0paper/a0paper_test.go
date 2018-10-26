package main

import (
     "testing"
)

func TestCanBuild(t *testing.T) {

    var a0paper A0Paper
    var s string
    var papers []int

    // A0 Paper OK
    papers = []int{0, 3}
    s = a0paper.canBuild(papers)
    if s != "Possible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    papers = []int{0,0,0,0,16}
    s = a0paper.canBuild(papers)
    if s != "Possible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    papers = []int{2,0,0,0,0,0,0,3,2,0,0,5,0,3,0,0,1,0,0,0,5}
    s = a0paper.canBuild(papers)
    if s != "Possible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    papers = []int{0,1,1,3,7}
    s = a0paper.canBuild(papers)
    if s != "Possible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    papers = []int{0,1,1,1,1,1,1,1,1,1,1,2}
    s = a0paper.canBuild(papers)
    if s != "Possible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    papers = []int{0,1,1,1,1,1,1,0,3,1,1,2}
    s = a0paper.canBuild(papers)
    if s != "Possible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    //
    // A0 Paper NG 
    papers = []int{0,0,0,0,15}
    s = a0paper.canBuild(papers)
    if s != "Impossible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    papers = []int{0,1,0,0,0,15}
    s = a0paper.canBuild(papers)
    if s != "Impossible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    papers = []int{0,1,1,1,1,1,1,1,1,1,1,1}
    s = a0paper.canBuild(papers)
    if s != "Impossible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

    papers = []int{0,1,0,1,1,1,1,1,1,1,1,1}
    s = a0paper.canBuild(papers)
    if s != "Impossible" {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, s)
    }

}


func TestCanBuildOneSize(t *testing.T) {
    var a0paper A0Paper
    var size int
    var sheets int
    var flag bool

    //
    // Can build 
    size = 0; sheets = 1;
    flag = a0paper.canBuildOneSize(size, sheets)
    if !flag {
        t.Fatalf("failed test :  size[%v] sheats:[%v] -> [%v]", size, sheets, flag)
    }

    size = 4; sheets = 16;
    flag = a0paper.canBuildOneSize(size, sheets)
    if !flag {
        t.Fatalf("failed test :  size[%v] sheats:[%v] -> [%v]", size, sheets, flag)
    }

    //
    // Can't build 
    size = 1; sheets = 1;
    flag = a0paper.canBuildOneSize(size, sheets)
    if flag {
        t.Fatalf("failed test :  size[%v] sheats:[%v] -> [%v]", size, sheets, flag)
    }

    size = 4; sheets = 15;
    flag = a0paper.canBuildOneSize(size, sheets)
    if flag {
        t.Fatalf("failed test :  size[%v] sheats:[%v] -> [%v]", size, sheets, flag)
    }

}

func TestCanBuildComposite(t *testing.T) {
    var a0paper A0Paper
    var papers []int
    var flag bool

    //
    // Can build 
    papers = []int{0, 1, 2}
    flag = a0paper.canBuildComposite(papers)
    if !flag {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, flag)
    }

    papers = []int{0, 1, 1, 2}
    flag = a0paper.canBuildComposite(papers)
    if !flag {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, flag)
    }

    papers = []int{0, 1, 1, 0, 4}
    flag = a0paper.canBuildComposite(papers)
    if !flag {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, flag)
    }

    papers = []int{0, 1, 1, 0, 0, 0, 0, 32}
    flag = a0paper.canBuildComposite(papers)
    if !flag {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, flag)
    }

    papers = []int{0, 1, 1, 0, 3, 8, 0, 31}
    flag = a0paper.canBuildComposite(papers)
    if !flag {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, flag)
    }


    //
    // Can't build 
    papers = []int{0, 1, 1}
    flag = a0paper.canBuildComposite(papers)
    if flag {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, flag)
    }

    papers = []int{0, 1, 1, 1}
    flag = a0paper.canBuildComposite(papers)
    if flag {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, flag)
    }

    papers = []int{0, 1, 1, 0, 0, 0, 0, 31}
    flag = a0paper.canBuildComposite(papers)
    if flag {
        t.Fatalf("failed test papers:[%v] -> [%v]", papers, flag)
    }

}

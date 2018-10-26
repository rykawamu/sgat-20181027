package main

import (
    "fmt"
    "math"
    "flag"
    "strings"
    "strconv"
)

type A0Paper struct {
    Papers []int
}

const (
    Possible string = "Possible"
    Impossible string = "Impossible"
)

var (
    inputPapers = flag.String("p", "{0,1,2,3}", "Input Papers: A0:0,A1:1,A2:2,A3:3 -> \"{0,1,2,3}\" ")
)

// From the previous two definitions it follows that the A series has the following useful property:
// Whenever you take an A(i) paper and you cut it in half (using a cut that passes through the centers of its longer sides), you will get two pieces of an A(i+1) paper.
// In other words, A1 is one half of A0, A2 is one half of A1, and so on.
// You are given a int[] A. A[i] represents the number of papers of size A(i) you have in stock. For example, A[4] is the number of A4 papers you currently have.
// You are not allowed to cut paper in any way. You can only connect papers (seamlessly and without any waste) by taping them together.
// The papers you connect this way must not overlap. Can you take some of the papers you have and assemble a paper of size A0? Return "Possible" if it can be done and "Impossible" otherwise.
//
// URL: http://community.topcoder.com/stat?c=problem_statement&pm=15005

func main() {

    flag.Parse()
    fmt.Printf("Can build A0 paper ? : Input papers -> %v\n",*inputPapers)

    s := strings.Replace(*inputPapers, " ", "", -1)
    s = strings.Replace(s, "{", "", -1)
    s = strings.Replace(s, "}", "", -1)

    inputlist := strings.Split(s, ",")
    var papers []int
    for _, v := range inputlist {
        sheets, _ := strconv.Atoi(v)
        papers = append(papers, sheets)
    }

    // check
    var a0 A0Paper
    judge := a0.canBuild(papers)
    fmt.Println(judge)
}

// A0用紙を作成可能な場合は「Possible」、不可の場合は「Impossible」を返す
func (a0 *A0Paper) canBuild(papers []int) string {
    a0.Papers = papers
    flag := false

    // check one size build
    for size, sheets := range a0.Papers {
        flag = a0.canBuildOneSize(size, sheets)
        if flag {
            break
        }
    } 
    if flag {
        return Possible 
    }

    // check composite build 
    flag = a0.canBuildComposite(a0.Papers)
    if flag {
        return Possible
    }

    return Impossible 
}

// 任意の用紙単体で対応できるか判定する
func (a0 A0Paper) canBuildOneSize(size, sheets int) bool {
    flag := false
    f64y := float64(size)
    f64Sheets := float64(sheets) 

    c := math.Pow(2, f64y)
    if c <= f64Sheets {
        flag = true
    }
    return flag
}

// 複数の用紙を混合させて対応可能か判定する
func (a0 A0Paper) canBuildComposite(papers []int) bool {

    flag := false
    minPaperSize := len(papers) - 1
    maxPieces := math.Pow(2, float64(minPaperSize))

    pieces := float64(0)
    for size := 1; size <= minPaperSize; size++ {
        sheets := papers[size]
        if 0 < sheets {
            pieces += math.Pow(2, float64(minPaperSize - size)) * float64(sheets)
        }
        if maxPieces <= pieces {
            flag = true
            break
        }
    }
    return flag
}


package main

import (
    "fmt"
    "time"
)

type StructYmd struct {
    Year int
    Month int
    Day int
}

// golangの日付書式作成
const (
    ymdFormat string = "2006-01-02"
)

//今日は4月22日であり、kamipeipaa君の誕生日です。kamipeipaa君はこの日が4=2+2と表せることに気づきました。
//kamipeipaa君は𝑋月𝑌日について𝑌の数字和が𝑋と等しいとき、この日を"HAPPY DAY"であると呼ぶことにしました。ここで数字和とは与えられた整数の各桁の数字が表す数の総和です。
//例えば、4月4日や10月19日は4=0+4、10=1+9となり、HAPPY DAYですが、12月3日や11月28日は12≠0+3、11≠2+8となりHAPPY DAYではありません。
//グレゴリオ暦の2015年においてHAPPY DAYがどれだけあるかkamipeipaa君に教えてあげてください。
func main () {
    target := "2015-01-01"
    startday := createStartday(target)
    endday := startday.AddDate(1, 0, 0)

    var list []StructYmd
    for day := startday; day.Before(endday); day = day.AddDate(0, 0, 1) {
        var ymd StructYmd
        ymd = set(day)
        list = append(list, ymd)
    }

    count := 0
    for _, d := range list {
        happy := d.check()
        if happy {
            count += 1
        }
    }
    fmt.Printf("Happy day count : [%v]\n", count)
}

// 開始日作成
func createStartday(targetday string) time.Time {
    t, _ := time.Parse(ymdFormat, targetday)
    return t
}

// StructYmdへの値設定
func set(day time.Time) StructYmd {
    var ymd StructYmd
    ymd.Year = day.Year()
    ymd.Month = int(day.Month())
    ymd.Day = day.Day()
    return ymd
}

// HappyDayチェック
func (ymd StructYmd) check() bool {
    // ymd.Day -> 27 : x = 2, y = 7
    y := ymd.Day % 10
    x := (ymd.Day - y) / 10
    xy := x + y

    happy := false
    if ymd.Month == xy {
        happy = true
    }

    s := fmt.Sprintf("%v-%v-%v", ymd.Year, ymd.Month, ymd.Day)
    fmt.Printf("check [%v]:[%v] -> [%v][%v] -> [%v]:[%v][%v]\n", s, ymd.Day, x, y, xy, ymd.Month, happy)

    return happy
}

package timeext

import (
    "time"
)

const ISO8601 = "2006-01-02" // date format ISO 8601
const OneDay = 24 * 60 * 60

func ParseDate(input string) (time.Time, error) {
    return time.Parse(ISO8601, input)
}

func FixWeekdaysInterval(start_date time.Time, end_date time.Time) (time.Time, time.Time) {
    // number of days in between
    diff := end_date.Sub(start_date).Seconds() / OneDay
    //fmt.Println(diff)
    
    // fix end_date
    for int(end_date.Weekday()) >= 5 {
        end_date = end_date.AddDate(0,0,-1)
    }
    
    // fix start_date
    start_date = end_date
    for diff > 0 {
        start_date = start_date.AddDate(0,0,-1)
        if int(start_date.Weekday()) < 5 {
            diff -= 1
        }
    }
    
    return start_date, end_date
}

package ystock

// wrapper for Yahoo Finance rest API

import (
    "fmt"
    "strconv"
    "net/url"
    "net/http"
    "encoding/csv"
    "time"
    "ext/timeext"
)

type ClosingValue struct {
    Value float64
    Date time.Time
}

func HistoricalClosingValues(symbol string, end time.Time, days int) (historical []ClosingValue) {
    // HACK: holidays make it hard to know the days when the market is closed,
    // query again with a longer range if not enough days were retrieved the
    // last time
    qdays := days
    historical = []ClosingValue{}
    for len(historical) < days {
        // on average 10 week days adds 4 more weekend days plus 1 occasional holiday
        qdays := qdays+(qdays/2)
        start := end.AddDate(0,0,(-1*qdays))
        //start, end = timeext.FixWeekdaysInterval(start, end)
        
        historical = query(symbol, start, end)
    }
    
    // remove extra days
    historical = historical[0:days]
    return
}

func query(symbol string, start time.Time, end time.Time) (historical []ClosingValue) {
    //fmt.Println(start, end)
    v := url.Values{}
    
    v.Set("s", symbol)
    v.Set("g", "d")
    v.Set("ignore", ".csv")
    
    v.Set("a", strconv.Itoa(int(start.Month()-1)))
    v.Set("b", strconv.Itoa(int(start.Day())))
    v.Set("c", strconv.Itoa(int(start.Year())))
    
    v.Set("d", strconv.Itoa(int(end.Month()-1)))
    v.Set("e", strconv.Itoa(int(end.Day())))
    v.Set("f", strconv.Itoa(int(end.Year())))
    
    query := fmt.Sprintf("http://ichart.yahoo.com/table.csv?%s", v.Encode())
    //fmt.Println(query)
    
    resp, _ := http.Get(query)
    reader := csv.NewReader(resp.Body)
    records, _ := reader.ReadAll()
    
    for _,val := range records[1:] {
        date,_ := timeext.ParseDate(val[0])
        closing,_ := strconv.ParseFloat(val[4], 64)
        
        historical = append(historical, ClosingValue{closing, date})
    }
    
    return
}

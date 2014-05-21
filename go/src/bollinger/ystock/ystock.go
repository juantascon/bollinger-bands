package ystock

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

func HistoricalClosingValues(symbol string, end_date time.Time, days int) (historical []ClosingValue) {
    rdays := days
    historical = []ClosingValue{}
    for len(historical) < days {
        // HACK: holidays make it hard to know the days when the market is closed,
        // requery if not enough days retrieved
        rdays := rdays+(rdays/2)
        start_date := end_date.AddDate(0,0,(-1*rdays))
        //start_date, end_date = timeext.FixWeekdaysInterval(start_date, end_date)
        
        historical = query(symbol, start_date, end_date)
    }
    
    // remove extra days
    historical = historical[0:days]
    return
}

func query(symbol string, start_date time.Time, end_date time.Time) (historical []ClosingValue) {
    //fmt.Println(start_date, end_date)
    v := url.Values{}
    
    v.Set("s", symbol)
    v.Set("g", "d")
    v.Set("ignore", ".csv")
    
    v.Set("a", strconv.Itoa(int(start_date.Month()-1)))
    v.Set("b", strconv.Itoa(int(start_date.Day())))
    v.Set("c", strconv.Itoa(int(start_date.Year())))
    
    v.Set("d", strconv.Itoa(int(end_date.Month()-1)))
    v.Set("e", strconv.Itoa(int(end_date.Day())))
    v.Set("f", strconv.Itoa(int(end_date.Year())))
    
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

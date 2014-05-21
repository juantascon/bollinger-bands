package bands

import (
    "time"
    "math"
    "bollinger/ystock"
    "bollinger/settings"
    //"fmt"
)

type Band struct {
    Date time.Time
    Close float64
    SMA float64
    Up float64
    Down float64
}

func All(symbol string) (bands []Band) {
    historical := ystock.HistoricalClosingValues(symbol, settings.TODAY, settings.SMA_DAYS+settings.GRAPH_DAYS)
    
    start := 0
    end := settings.SMA_DAYS
    
    for i := 0; i < settings.GRAPH_DAYS; i++{
        //fmt.Println(historical[start].Date.String(), historical[end-1].Date.String())
        bands = append(bands, One(historical[start:end]))
        start++
        end++
    }
    
    return
}

func One(historical []ystock.ClosingValue) (result Band) {
    size := len(historical)
    //fmt.Println(size, historical[0].Date.String(), historical[size-1].Date.String())
    
    sum := float64(0)
    for _,h := range (historical) {
        sum += h.Value
    }
    
    // simple moving average
    sma := sum / float64(size)
    
    squares := float64(0)
    for i := 0; i < size; i++ {
        squares += math.Pow((historical[i].Value-sma), 2)
    }
    // standard deviation
    dev := math.Sqrt( squares / float64(size) )
    
    // upper band
    up := sma + (settings.STANDARD_DEVIATIONS * dev)
    
    // lower band
    down := sma - (settings.STANDARD_DEVIATIONS * dev)
    
    return Band{historical[0].Date, historical[0].Value, sma, up, down}
}
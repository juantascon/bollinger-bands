package plot

import (
    "code.google.com/p/plotinum/plot"
    "code.google.com/p/plotinum/plotter"
    "image/color"
    "bollinger/bands"
    "fmt"
)

func PlotBands(symbol string, all []bands.Band) {
    dclose := make(plotter.XYs, len(all))
    dsma := make(plotter.XYs, len(all))
    dup := make(plotter.XYs, len(all))
    ddown := make(plotter.XYs, len(all))
    
    for i,b := range all {
        dclose[i].X = float64(-1*i)
        dclose[i].Y = b.Close
        
        dsma[i].X = float64(-1*i)
        dsma[i].Y = b.SMA
        
        dup[i].X = float64(-1*i)
        dup[i].Y = b.Up
        
        ddown[i].X = float64(-1*i)
        ddown[i].Y = b.Down
    }
    
    p,_ := plot.New()
    
    p.Title.Text = fmt.Sprintf("Bollinger Bands: %s", symbol)
    p.X.Label.Text = "Time (Days)"
    p.Y.Label.Text = "Value"
    
    p.Add(plotter.NewGrid())
    
    lclose,_ := plotter.NewLine(dclose)
    
    lsma,_ := plotter.NewLine(dsma)
    lsma.LineStyle.Color = color.RGBA{B: 255, A: 255}
    
    lup,_ := plotter.NewLine(dup)
    lup.LineStyle.Color = color.RGBA{R: 255, A: 255}
    
    ldown,_ := plotter.NewLine(ddown)
    ldown.LineStyle.Color = color.RGBA{G: 255, A: 255}
    
    p.Add(lclose, lsma, lup, ldown)
    p.Legend.Add("Closing", lclose)
    p.Legend.Add("SMA", lsma)
    p.Legend.Add("Up", lup)
    p.Legend.Add("Down", ldown)
    
    p.Save(16, 9, fmt.Sprintf("%s.png", symbol))
}

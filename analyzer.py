#! /usr/bin/env python3

from datetime import date, timedelta
import argparse

from bollinger import bands, strategies, settings

# Analyzes each strategy efficiency by comparing an hipotetical
# investment DELTA days ago with the selling price TODAY

DELTA=20

strategies_list = [strategies.UpOnce, strategies.DownOnce, strategies.MoreUp, strategies.MoreDown]
symbols_list = [
    "AAPL", "ADBE", "AMD", "AMZN", "FB", "IBM", "INTC",
    "MSFT", "ZNGA", "ADSK", "AKAM", "EA", "EBAY", "SBUX",
    "TACT", "TRIP", "TTWO", "TXN", "VIA", "YHOO", "CSCO"
]

# all = invest on all, best = optimal investments
totals = { "all": 0, "best": 0 }
for s in strategies_list:
    totals[s.name] = 0

# change start date in order to calculate the bands
settings.TODAY = settings.TODAY - timedelta(days=DELTA)

for symbol in symbols_list:
    
    print ("[ %s ] " % (symbol), end="")
    
    b = bands.Bands(symbol)
    b.fetch()
    result = b.all()
    
    gain = (result[DELTA-1].get("close") - result[0].get("close"))
    print ("gain: %.4f" %(gain), end="")
    
    totals["all"] += gain
    if gain > 0:
        totals["best"] += gain
    
    for s in strategies_list:
        suggestion = s(b).invest()
        print (" %s: %s" % (s.name, "YES" if suggestion else "NO"), end="")
        if suggestion:
            totals[s.name] += gain
    print("")
    
print("totals: %s" %(totals))
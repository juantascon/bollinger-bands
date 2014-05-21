##AUTHOR: Juan Diego Tascón
##EMAIL: juantascon@horlux.org
##CREATED: 2014-05-05

# 1. Charts:

The charts are very simple, please read the code comments for more details.

One thing is missing though, the example images on the description had a
weird X axis, it is more useful[1][2]  if the x axis represents
time, in this implementation it represents the number of past days.

* [1] [stockcharts.com](http://stockcharts.com/help/doku.php?id=chart_school:technical_indicators:bollinger_bands)
* [2] [biz.yahoo.com](http://biz.yahoo.com/charts/guide13.html)

# 2. Investment Strategies:

I came up with 4 simple independet invesment strategies:

* uponce: invest if on the last 20 days close price passes at least once over upper band

* dowonce: invest if on the last 20 days close price passes at least once under lower band

* moreup: invest if on the last 20 days the number of times close price passes
over the uppper band is greater than the number of times close price passes
under the lower band
    
* moredown: invest if on the last 20 days the number of times close price passes
under the lower band is greater than the number of times close price passes
over the uppper band

I wasn't sure which strategy would be best so I went further and
developed a helper script (analyzer.py) that analyzes each strategy
efficiency by comparing an hipotetical investment made DELTA days
ago with the selling price TODAY (not really today but the last
registered closing price). This approach will allow me to come up
with better strategies and test their performance instantly against
other candidates.

I took a sample of 21 stocks, all from the technology sector and I used
DELTA=20 to compare the price 20 days before TODAY with the price of
the same stock TODAY.

The results seems to indicate that "moredown" and "downonce" are the best
candidates. My own personal conclusion is that in the short term the more
times a stock lower its price below the lower bollinger band the more
likely is that the price increases on the short term as well. It is "weird"
that a price goes below the lower bollinger band, when it happends the price
might tend to normalize itself by going back up. Becareful though because
these was only tested on tech stocks, it might differ on other stocks.

In the end to be honest I wouldn't use any of these for real investments,
the stock market is too caotic, with way too many variables, this makes
it hard to define completely safe invesment strategies, there will always
be risk involved.

# 3 Python:

* bolly.py:
    * generates (outputs to $PWD/$SYMBOL.png) bollinger bands, ex: bolly.py plot AMZN FB
    * prints wether or not you should invest given a strategy, ex: bolly.py suggest AMZN FB -s moredown

* analizer.py:
    * analyzes which strategies are better

# 4 GO:

* main.go: 
    * generates (outputs to $PWD/$SYMBOL.png) bollinger bands, ex: go run main.go -p AMZN FB
    * prints wether or not you should invest given a strategy, ex: go run main.go -s moredown AMZN FB
    * each symbol is processed using goroutines so that blocking syscalls won't stop other symbol's processing
    
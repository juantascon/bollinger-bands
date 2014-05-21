from . import bands, settings

import matplotlib.pyplot as plt

class Plot:
    def __init__(self, bands):
        self.bands = bands
    
    def save(self):
        result = self.bands.all()
        
        # indexes are negative because they
        # represent the number of days in the past
        index = [ (i.get("index")+1)*-1 for i in result ]
        close = [ i.get("close") for i in result ]
        up = [ i.get("up") for i in result ]
        middle = [ i.get("middle") for i in result ]
        down = [ i.get("down") for i in result ]
        
        plt.plot(index, up, label="upper band")
        plt.plot(index, down, label="lower band")
        plt.plot(index, middle, label="middle band")
        plt.plot(index, close, label="close price")
        
        plt.xlabel("Past days (0 = today)")
        plt.ylabel("Value (USD$)")
        plt.title("%s bollinger bands" % (self.bands.symbol))
        
        # enables the grid for every single decimal value
        plt.minorticks_on()
        plt.grid(True, which="both")
        
        legend = plt.legend(fancybox=True, loc="best")
        legend.get_frame().set_alpha(0.5)
        
        plt.savefig(self.bands.symbol+".png")
        
        # the plot must be closed for otherwhise matplotlib
        # will paint over previous plots
        plt.close()

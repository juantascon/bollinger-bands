#! /usr/bin/env python3

from bollinger import bands, plot, strategies
import argparse

parser = argparse.ArgumentParser(description="plots bollinger bands or suggests investments", epilog="example: bolly.py plot AMZN FB")

parser.add_argument("action", metavar="ACTION", choices=["plot", "suggest"], help="either plot or suggest")
parser.add_argument("symbols", metavar="SYMBOL", nargs="+", help="stock symbols")
parser.add_argument("-s", "--strategy", choices=["uponce", "downonce", "moreup", "moredown"], default="moredown", help="selects invesment strategy")

args = parser.parse_args()

if args.action == "plot":
    for symbol in args.symbols:
        print("plot [ %s ]: " %(symbol), end="")
        b = bands.Bands(symbol)
        b.fetch()
        try:
            p = plot.Plot(b)
            p.save()
            print("OK")
        except Exception as ex:
            print("FAIL: (%s)"%(ex))
            
if args.action == "suggest":
    for symbol in args.symbols:
        print("suggest [ %s ]: " %(symbol), end="")
        b = bands.Bands(symbol)
        b.fetch()
        try:
            if args.strategy == "uponce": s = strategies.UpOnce(b)
            elif args.strategy == "downonce": s = strategies.DownOnce(b)
            elif args.strategy == "moreup": s = strategies.MoreUp(b)
            elif args.strategy == "moredown": s = strategies.MoreDown(b)
            
            print("YES" if s.invest() else "NO")
        except Exception as ex:
            print("FAIL: (%s)"%(ex))

#!/usr/bin/env python

import math

def aperture(ev, st):
    if 1/st < 1:
        return math.sqrt((2**ev)*1/st)
    elif 1/st == 1:
        return math.sqrt((2**ev))

def main():
    ev = float(input("Enter EV: "))
    shutter_time = float(input("Enter shutter time (1/t): "))

    a = aperture(ev, shutter_time)
    print(a)

if __name__ == "__main__":
    main()


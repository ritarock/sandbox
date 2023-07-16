from cgitb import reset
import re
from unittest import result


n = int(input())
h = list(map(int, input().split()))


for i,v in enumerate(h, 1):
    if i == 1:
        hh = v
        result = i
    else:
        if hh < v:
            hh = v
            result = i

print(result)

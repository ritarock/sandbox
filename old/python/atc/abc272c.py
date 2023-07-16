import re


n = int(input())
a = list(map(int, input().split()))
a.sort(reverse=True)

odds = []
evens = []

for i in range(n):
    if a[i] % 2 == 0:
        evens.append(a[i])
    else:
        odds.append(a[i])

if len(evens) < 2 and len(odds) < 2:
    print(-1)
elif len(odds)<2:
    print(evens[0]+evens[1])
elif len(evens)<2:
    print(odds[0]+odds[1])
else:
    print(max(evens[0]+evens[1], odds[0]+odds[1]))

s = input()

s = s[::-1]

dr = "dreamer"[::-1]
er = "eraser"[::-1]
d = "dream"[::-1]
e = "erase"[::-1]

while len(s)>=5:
    if len(s) >= 7 and s[:7] == dr:
        s = s[7:]
        continue
    if len(s) >= 6 and s[:6] == er:
        s = s[6:]
        continue
    if len(s) >= 5 and (s[:5] == d or s[:5] == e):
        s = s[5:]
        continue
    else:
        break

if len(s) == 0:
    print("YES")
else:
    print("NO")

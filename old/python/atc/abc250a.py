h, w = map(int, input().split())
r, c = map(int, input().split())

result = 4

if r + 1 > h:
    result -=1

if r - 1 < 1:
    result -=1

if c + 1 > w:
    result -=1

if c - 1 < 1:
    result -=1

print(result)

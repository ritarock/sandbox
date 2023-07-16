n = int(input())
a = list(map(int, input().split()))

result = 0

a.sort(reverse=True)
for i, j in enumerate(a):
    if (i %2 == 0):
        result += j
    else:
        result -= j

print(result)

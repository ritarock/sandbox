n = int(input())
a = list(map(int, input().split()))

for i in range(0, 2001):
    if i not in a:
        print(i)
        exit()

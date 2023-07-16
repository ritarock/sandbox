n, m = map(int, input().split())

l = []
for _ in range(m):
    i = list(map(int, input().split()))
    i.pop(0)
    l.append(i)

for i in range(1, n+1):
    for j in range(i+1, n+1):
        count = 0
        for k in range(m):
            if i in l[k] and j in l[k]:
                count+=1
                break
        if count == 0:
            print("No")
            exit()

print("Yes")

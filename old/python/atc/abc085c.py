N, Y = map(int, input().split())

for x in range(N+1):
    for y in range(N-x+1):
        z = N-x-y
        if (Y == 10000*x + 5000*y + 1000*(N-x-y)):
            print(x,y,z)
            exit()
        continue

print(-1,-1,-1)

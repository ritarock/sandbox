n = int(input())
t, x, y = 0, 0, 0

for i in range (n):
    next_t, next_x, next_y = map(int, input().split())
    d = abs(next_x-x) + abs(next_y-y)
    dt = next_t-t
    if d > dt or (d - dt)%2 == 1:
        print("No")
        exit()
    else:
        t, x, y = next_t, next_x, next_y

print("Yes")

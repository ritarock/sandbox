x, y, z = map(int, input().split())

if abs(x) < abs(y) or x * y < 0:
    print(abs(x))

elif y * z > 0 and abs(z) < abs(y):
    print(abs(x))

elif y * z < 0:
    print(abs(x) + 2*abs(z))

else:
    print(-1)

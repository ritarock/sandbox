from hashlib import new


h, w = map(int, input().split())

c = [list(input()) for _ in range(h)]

new_c = []

for i in range(len(c[0])):
    tmp = []
    for v in c:
        tmp.append(v[i])
    new_c.append(''.join(tmp))

result = []
for i in new_c:
    result.append(str(i.count('#')))

print(' '.join(result))

n, q = map(int, input().split())

l = []
for _ in range(n):
    ll = list(map(int, input().split()))
    l.append(ll)


for _ in range(q):
    s, t = map(int, input().split())
    print(l[s-1][t])

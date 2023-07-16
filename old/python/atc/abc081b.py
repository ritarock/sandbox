n = int(input())
a = list(map(int, input().split()))
cnt = 0

while (all([item%2==0 for item in a])):
    a = [item / 2 for item in a]
    cnt+=1

print(cnt)

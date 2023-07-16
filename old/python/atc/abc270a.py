a, b = map(int, input().split())

# 1 2 4
def q(i):
    if i == 0:
        return [0]
    elif i == 1:
        return [1]
    elif i == 2:
        return [2]
    elif i == 4:
        return [4]
    elif i == 3:
        return [1, 2]
    elif i == 5:
        return [1, 4]
    elif i == 6:
        return [2, 4]
    elif i == 7:
        return [1,2,4]

_a = q(a)
_b = q(b)

print(sum(list(set(_a+_b))))

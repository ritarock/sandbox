ax, ay = map(int, input().split())
bx, by = map(int, input().split())
cx, cy = map(int, input().split())
dx, dy = map(int, input().split())

def vector(ax, ay, bx, by):
    return (ax-bx, ay-by)

def inner_traiangle(ax, ay, bx, by, cx, cy, point_x, point_y):
    ab = vector(bx, by, ax, ay)
    bp = vector(point_x, point_y, bx, by)

    bc = vector(cx, cy, bx, by)
    cp = vector(point_x, point_y, cx, cy)

    ca = vector(ax, ay, cx, cy)
    ap = vector(point_x, point_y, ax, ay)

    opter_product_a = ab[0]*bp[1] - ab[1]*bp[0]
    opter_product_b = bc[0]*cp[1] - bc[1]*cp[0]
    opter_product_c = ca[0]*ap[1] - ca[1]*ap[0]

    if (opter_product_a >0 and opter_product_b >0 and opter_product_c >0) or (opter_product_a <0 and opter_product_b <0 and opter_product_c <0):
        return True


a = inner_traiangle(ax, ay, bx, by, cx, cy, dx, dy)
b = inner_traiangle(bx, by, cx, cy, dx, dy, ax, ay)
c = inner_traiangle(cx, cy, dx, dy, ax, ay, bx, by)
d = inner_traiangle(dx, dy, ax, ay, bx, by, cx, cy)

if (a == True) or (b == True) or (c == True) or (d == True):
    print("No")
else:
    print("Yes")

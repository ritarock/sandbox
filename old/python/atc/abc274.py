from decimal import Decimal, ROUND_HALF_UP

a, b = map(int, input().split())

s = Decimal(b/a).quantize(Decimal('0.001'), rounding=ROUND_HALF_UP)

print(s)

n = int(input())

h = str(hex(n))

print(('00' + h.replace('0x', '').upper())[-2:])

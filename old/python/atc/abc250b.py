n, a, b = map(int, input().split())


def create_line(n, b):
    line = ''
    for i in range(1, n+1):
        if i % 2 != 0:
            line += '.' * b
        else:
            line += '#' * b
    return line

def create_block(a, line):
    block = ''
    for _ in range(a):
        block += line + '\n'
    return block

def create_result(n, block: str):
    result = ''
    inversion_block = block.replace('.', '-').replace('#', '.').replace('-', '#')
    for i in range(1, n+1):
        if i %2 != 0:
            result += block
        else:
            result += inversion_block
    return result

line = create_line(n, b)
block = create_block(a, line)
print(create_result(n, block))

def main():
    a = 0
    b = 1
    cnt = 0
    while cnt < 100:
        print(b)
        a, b = b, a + b
        cnt = cnt + 1


if __name__ == '__main__':
    main()

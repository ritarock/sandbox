s = input()

if 'a' in s:
    reverse = ''.join(reversed(s))
    i = reverse.find('a')
    s[i]
    print(len(s)-i)
else:
    print(-1)

def getId(code):
    f = 0
    b = 127
    l = 0
    r = 7
    for c in code:
        if c == 'F':
            b = (f+b)//2
        elif c == 'B':
            f = (f+b)//2 + 1
        elif c == 'L':
            r = (l+r)//2
        elif c == 'R':
            l = (l+r)//2 + 1
    return f * 8 + l


with open('input.txt') as f:
    ids = []
    for line in f:
        ids.append(getId(line))
    ids.sort()
    all_ids = set(range(ids[0], ids[len(ids)-1]+1))
    print(all_ids.difference(set(ids)).pop())

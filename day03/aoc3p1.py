with open('input.txt') as f:
    m = []
    for line in f:
        m.append(line)
    x = 0
    y = 0
    count = 0
    while (y < len(m)):
        count += 1 if m[y][x] == '#' else 0
        x = (x+3) % (len(m[0])-1)
        y += 1
    print(count)


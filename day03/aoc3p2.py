def solve(m, xs, ys):
    x = 0
    y = 0
    count = 0
    while (y < len(m)):
        count += 1 if m[y][x] == '#' else 0
        x = (x+xs) % (len(m[0])-1)
        y += ys
    return count

with open('input.txt') as f:
    m = []
    for line in f:
        m.append(line)
    product = 1
    slopes = [(1,1), (3,1), (5,1), (7,1), (1,2)]
    for slope in slopes:
        product *= solve(m, slope[0], slope[1])
    print(product)


with open('input.txt') as f:
    count = 0
    for line in f:
        tok = line.split(' ')
        ran = list(map(int, tok[0].split('-')))
        if tok[2].count(tok[1][0]) in range(ran[0], ran[1]+1):
            count += 1
    print(count)

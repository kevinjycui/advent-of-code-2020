with open('input.txt') as f:
    count = 0
    for line in f:
        tok = line.split(' ')
        pos = list(map(int, tok[0].split('-')))
        if (tok[2][pos[0]-1] == tok[1][0]) != (tok[2][pos[1]-1] == tok[1][0]):
            count += 1
    print(count)

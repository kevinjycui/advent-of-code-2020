with open('input.txt') as f:
    q = f.read().split('\n\n')
    count = 0;
    for l in q:
        count += len(set(l.replace('\n', '')))
    print(count)

with open('input.txt') as f:
    q = f.read().split('\n\n')
    count = 0;
    for l in q:
        sol = set(l.split('\n')[0])
        for k in l.split('\n'):
            if k != '':
                sol = sol.intersection(set(k))
        count += len(sol)
    print(count)

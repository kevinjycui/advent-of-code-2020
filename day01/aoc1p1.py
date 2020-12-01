with open('input.txt') as f:
    a = []
    for line in f:
        a.append(int(line))
    for year in a:
        if 2020-year in a:
            print(year*(2020-year))
            break

with open('input.txt') as f:
    a = []
    for line in f:
        a.append(int(line))
    for year in a:
        for year2 in a:
            if 2020-year-year2 in a:
                print(year*year2*(2020-year-year2))
                a = []

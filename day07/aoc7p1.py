def hasshiny(chain, bag):
    if bag == 'shiny gold bags':
        return 1
    elif bag not in chain:
        return 0
    for bag in chain[bag]:
        if hasshiny(chain, bag) == 1:
            return 1
    return 0

with open('input.txt') as f:
    chain = {}
    for line in f:
        p = line.split(' contain ')
        colour = p[0]
        bags = p[1].replace('\n', '').replace('.', '').split(', ')
        chain[p[0]] = chain.get(p[0], [])
        for bag in bags:
            name = ' '.join(bag.split(' ')[1:])
            if bag.split(' ')[0] == '1':
                name += 's'
            chain[p[0]].append(name)
    count = 0
    for bag in chain.keys():
        if bag != 'shiny gold bags':
            count += hasshiny(chain, bag)
    print(count)

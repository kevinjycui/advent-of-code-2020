def inshiny(chain, bag):
    if bag not in chain:
        return 1
    count = 1
    for num, bag in chain[bag]:
        count += num * inshiny(chain, bag)
    return count

with open('input.txt') as f:
    chain = {}
    for line in f:
        p = line.split(' contain ')
        colour = p[0]
        bags = p[1].replace('\n', '').replace('.', '').split(', ')
        chain[p[0]] = chain.get(p[0], [])
        for bag in bags:
            if bag.split(' ')[0] == 'no':
                continue
            num = int(bag.split(' ')[0])
            name = ' '.join(bag.split(' ')[1:])
            if bag.split(' ')[0] == '1':
                name += 's'
            chain[p[0]].append((num, name))
    print(inshiny(chain, 'shiny gold bags') - 1)

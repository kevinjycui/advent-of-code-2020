with open('input.txt') as f:
    inp = f.read()
    ppts = inp.split('\n\n')
    count = 0
    valid = {'byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'}
    for p in ppts:
        fields = p.replace('\n', ' ').split(' ')
        a = []
        for field in fields:
            a.append(field.split(':')[0])
        if set(a).intersection(valid) == valid:
            count += 1
    print(count)

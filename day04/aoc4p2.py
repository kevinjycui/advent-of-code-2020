with open('input.txt') as f:
    inp = f.read()
    ppts = inp.split('\n\n')
    count = 0
    valid = {'byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'}
    for p in ppts:
        fields = p.replace('\n', ' ').split(' ')
        a = {}
        for field in fields:
            if len(field.split(':')) == 2:
                a[field.split(':')[0]] = field.split(':')[1]
        if set(a.keys()).intersection(valid) == valid and int(a['byr']) in range(1920, 2003) and int(a['iyr']) in range(2010, 2021) and int(a['eyr']) in range(2020, 2031) and ((a['hgt'][-2:] == 'cm' and int(a['hgt'][:-2]) in range(150, 194)) or (a['hgt'][-2:] == 'in' and int(a['hgt'][:-2]) in range(59, 77))) and (a['hcl'][0] == '#' and all(c in '0123456789abcdef' for c in a['hcl'][1:])) and a['ecl'] in ('amb','blu','brn','gry','grn','hzl','oth') and a['pid'].isdigit() and len(a['pid']) == 9:
            count += 1
    print(count)

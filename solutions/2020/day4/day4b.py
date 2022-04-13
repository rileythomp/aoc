def isValid(curpassport):
    fields = ['ecl', 'pid', 'eyr', 'hcl', 'byr', 'iyr', 'cid', 'hgt']
    for field in curpassport:
        key = field.split(':')[0]
        val = field.split(':')[1]
        if key == 'ecl':
            if val not in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
                return False
        elif key == 'pid':
            if len(val) != 9 or len([x for x in val if x not in '0123456789']) > 0: return False
        elif key == 'eyr':
            if len(val) != 4 or val < '2020' or val > '2030': return False
        elif key == 'hcl':
            if val[0] != '#' or len(val) != 7 or len([x for x in val[1:] if x not in '0123456789abcdef']) > 0: return False
        elif key == 'byr':
            if len(val) != 4 or val < '1920' or val > '2002': return False
        elif key == 'iyr':
            if len(val) != 4 or val < '2010' or val > '2020': return False
        elif key == 'hgt':
            if val[-2:] == 'cm':
                if len(val) != 5 or val[0:3] < '150' or val[0:3] > '193': return False
            elif val[-2:] == 'in':
                if len(val) != 4 or val[0:2] < '59' or val[0:2] > '76': return False
            else: return False
        if key in fields:
            fields.remove(key)
        
    return fields == [] or fields == ['cid']

lines = [x for x in open('4.in')]
curpassport = []
valid = 0
for line in lines:
    if line == '\n':
        if isValid(curpassport): valid += 1
        curpassport = []
        continue
    curpassport.extend(line.strip('\n').split(' '))
if isValid(curpassport): valid += 1
print(valid)

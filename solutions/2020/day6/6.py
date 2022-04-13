lines = [x for x in open('6.txt')]
total = 0
str = ''
for line in lines:
    line = line.strip('\n')
    if line == '':
        total += len(str)
        str = ''
        continue
    for char in line:
        if char not in str:
            str += char

total += len(str)
print(total)

total = 0
common = -1
for line in lines:
    line = line.strip('\n')
    if line == '':
        total += len(common)
        common = -1
        continue
    if common == -1:
        common = set(line)
        continue
    common = common.intersection(set(line))

total += len(common)
print(total)

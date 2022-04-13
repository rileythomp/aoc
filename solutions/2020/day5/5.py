ans = 0
plane = [[0 for y in range(8)] for x in range(128)]
for line in open('5.txt'):
    line = line.strip('\n')
    rowL, rowU = 0, 127
    colL, colU = 0, 7
    for val in line:
        if val == 'F':
            rowU = rowU - (rowU+1-rowL)/2
        elif val == 'B':
            rowL = rowL + (rowU+1-rowL)/2
        elif val == 'L':
            colU = colU - (colU+1-colL)/2
        elif val == 'R':
            colL = colL + (colU+1-colL)/2
    plane[int(rowL)][int(colL)] = 1
    seatId = 8*rowL + colL
    if seatId > ans:
        ans = seatId
print(int(ans))

for i in range(len(plane)):
    for j in range(len(plane[i])):
        if plane[i][j] == 0:
            if (j > 0 and plane[i][j-1] == 1) or \
            (j == 0 and i > 0 and plane[i-1][j+7] == 1):
                print(8*i + j) 
                exit()
    

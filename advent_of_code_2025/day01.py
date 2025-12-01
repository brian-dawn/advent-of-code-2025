import sys

spot = 50
part1 = 0

for line in sys.stdin:
    line = line.strip()

    num = int(line[1:])

    if line[0] == "L":
        spot -= num
    elif line[0] == "R":
        spot += num
    else:
        raise Exception("Unknown direction")
    spot = spot % 100

    if spot == 0:
        part1 += 1

print(part1)





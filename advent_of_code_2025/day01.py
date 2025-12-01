import sys

spot = 50
part1 = 0
part2 = 0

for line in sys.stdin:
    line = line.strip()

    num = int(line[1:])


    new_spot = spot

    if line[0] == "L":
        new_spot -= num
    elif line[0] == "R":
        new_spot += num
    else:
        raise Exception("Unknown direction")

    if new_spot < 0 or new_spot >= 100:

        if new_spot < 0 and spot != 0:
            part2 += 1
        part2 += int(abs(new_spot) / 100)

    if new_spot == 0 and spot != 0:
        part2 += 1
    

    new_spot = new_spot % 100
    if new_spot == 0:
        part1 += 1

    spot = new_spot
    assert spot < 100 and spot >= 0


print(part1)
print(part2)





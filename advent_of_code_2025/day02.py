import sys
import re


input = sys.stdin.read()


pairings = re.findall(r"(\d+)-(\d+)", input)


def is_invalid(n: int) -> bool:
    s = str(n)
    split = len(s) // 2

    return s[:split] == s[split:]


def is_invalid_advanced(n: int) -> bool:
    """
    Returns true if is made up of some sequence of digits at least twice
    """

    # Sequence is the first time we find a character that's the same as the current one.
    s = str(n)

    for sequence_length in range(1, len(s) // 2 + 1):
        # Check if match.
        seq = s[:sequence_length]

        # Occurences * length should match the o.g. string length otherwise there are other nums breaking
        # the sequence.
        if s.count(seq) * len(seq) == len(s):
            return True

    return False


assert is_invalid_advanced(999)
assert is_invalid_advanced(1188511885)

part1 = 0
part2 = 0
for start_str, end_str in pairings:
    start = int(start_str)
    end = int(end_str)

    for i in range(start, end + 1):
        if is_invalid(i):
            part1 += i
            part2 += i

        elif is_invalid_advanced(i):
            part2 += i


print(f"part 1: {part1}")
print(f"part 2: {part2}")

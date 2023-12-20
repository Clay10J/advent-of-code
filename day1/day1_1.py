def day1():
    total = 0
    lines = []
    with open("./day1_1_input.txt") as f:
        for line in f:
            total += find_value(line)
    print(total)


def find_value(line):
    first, last = "", ""
    for i in range(len(line)):
        if line[i].isdigit():
            first = line[i]
            break

    for i in range(len(line) - 1, -1, -1):
        if line[i].isdigit():
            last = line[i]
            break

    return int(first + last)


day1()

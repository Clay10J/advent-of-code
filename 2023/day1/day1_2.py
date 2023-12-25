numStrs = [
    ("zerone", "01"),
    ("oneight", "18"),
    ("twone", "21"),
    ("threeight", "38"),
    ("fiveight", "58"),
    ("sevenine", "79"),
    ("eightwo", "82"),
    ("eighthree", "83"),
    ("nineight", "98"),
    ("zero", "0"),
    ("one", "1"),
    ("two", "2"),
    ("three", "3"),
    ("four", "4"),
    ("five", "5"),
    ("six", "6"),
    ("seven", "7"),
    ("eight", "8"),
    ("nine", "9")
]


def day1():
    total = 0
    with open("./day1_1_input.txt") as f:
        for line in f:
            total += find_value(line)
    print(total)

def find_value(line):
    for old, new in numStrs:
        line = line.replace(old, new)

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

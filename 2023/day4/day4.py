class Day4:
    def __init__(self) -> None:
        self.solution1 = 0
        self.solution2 = 0

    def day4(self):
        with open("./day4_input.txt") as f:
            for line in f:
                cards = line.split(":")[1]
                winningNums, haveNums = cards.split("|")
                self.solution1 += self.pointsForCard(winningNums.split(), haveNums.split())

    def pointsForCard(self, winningNums, haveNums) -> int:
        numMatches = 0
        winning = {int(x) for x in winningNums}
        for num in haveNums:
            if int(num) in winning:
                numMatches += 1
        points = pow(2, numMatches - 1) if numMatches > 0 else 0
        return points


soln = Day4()
soln.day4()
print(soln.solution1)

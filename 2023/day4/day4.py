from typing import Dict


class Day4:
    def __init__(self) -> None:
        self.solution1 = 0
        self.solution2 = 0
        self.cardCounts: Dict[int, int] = dict()

    def day4(self):
        with open("./in.txt") as f:
            for line in f:
                card_id, card = line.split(":")
                card_id = int(card_id.split()[1])
                winningNums, haveNums = card.split("|")
                self.cardCounts[card_id] = self.cardCounts.get(card_id, 0) + 1

                numMatches = self.pointsForCard(winningNums.split(), haveNums.split())
                points = pow(2, numMatches - 1) if numMatches > 0 else 0
                self.solution1 += points

                print(numMatches)
                for i in range(1, numMatches + 1):
                    self.cardCounts[card_id + i] = 1 * self.cardCounts.get(
                        card_id + i, 1
                    )
        print(self.cardCounts)

        self.solution2 = sum(self.cardCounts.values())

    def pointsForCard(self, winningNums, haveNums) -> int:
        numMatches = 0
        winning = {int(x) for x in winningNums}
        for num in haveNums:
            if int(num) in winning:
                numMatches += 1
        return numMatches


soln = Day4()
soln.day4()
print(soln.solution1)
print(soln.solution2)

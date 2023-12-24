class Day2:
    MAX_RED = 12
    MAX_GREEN = 13
    MAX_BLUE = 14

    def __init__(self) -> None:
        self.solution1 = 0
        self.solution2 = 0

    def day2(self):
        with open("./day2_input.txt") as f:
            for line in f:
                game, sets = line.split(":")
                game_id = game.split()[1]
                if self.isPossible(sets):
                    self.solution1 += int(game_id)
                self.solution2 += self.getPower(sets)

    def isPossible(self, sets) -> bool:
        highest_red, highest_green, highest_blue = 0, 0, 0
        rounds = sets.split(";")
        for round in rounds:
            for r in round.split(","):
                rd = r.split()
                count, color = rd[0], rd[1]
                if color == "red":
                    highest_red = max(highest_red, int(count))
                elif color == "green":
                    highest_green = max(highest_green, int(count))
                else:
                    highest_blue = max(highest_blue, int(count))

                if (
                    highest_red > self.MAX_RED
                    or highest_green > self.MAX_GREEN
                    or highest_blue > self.MAX_BLUE
                ):
                    return False
        return True

    def getPower(self, sets) -> int:
        highest_red, highest_green, highest_blue = 0, 0, 0
        rounds = sets.split(";")
        for round in rounds:
            for r in round.split(","):
                rd = r.split()
                count, color = rd[0], rd[1]
                if color == "red":
                    highest_red = max(highest_red, int(count))
                elif color == "green":
                    highest_green = max(highest_green, int(count))
                else:
                    highest_blue = max(highest_blue, int(count))

        return highest_red * highest_green * highest_blue


soln = Day2()
soln.day2()
print(f"Solution1: {soln.solution1}")
print(f"Solution2: {soln.solution2}")

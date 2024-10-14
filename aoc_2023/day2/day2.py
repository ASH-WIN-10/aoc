import re

with open('input.txt', 'r') as file:
    lines = file.readlines()

def part_1():
    possible_games = []
    for line in lines:
        possible = True

        sets = line.split(':')[1].split(';')
        for set in sets:
            red, blue, green = 0, 0, 0
            set_of_cubes = set.split(",")
            for cube in set_of_cubes:
                match = "".join(re.findall("[0-9]", cube))
                number_of_cubes = int(match)
                if cube.find("red") != -1:
                    red = red + number_of_cubes
                if cube.find("blue") != -1:
                    blue = blue + number_of_cubes
                if cube.find("green") != -1:
                    green = green + number_of_cubes

            if red > 12 or blue > 14 or green > 13:
                possible = False
                break

        if possible:
            match = re.findall("[0-9]", line.split(':')[0])
            game_id = int("".join(match))
            possible_games.append(game_id)

    print(sum(possible_games))


def day_2():
    power_of_cubes = []
    for line in lines:
        sets = line.split(':')[1].split(';')
        max_in_cubes = [0, 0, 0]
        for set in sets:
            all_cubes = { "red": 0, "blue": 0, "green": 0 }

            set_of_cubes = set.split(",")
            for cube in set_of_cubes:
                match = re.findall("[0-9]", cube)
                number_of_cubes = int("".join(match))
                for color_cube in all_cubes:
                    if cube.find(color_cube) != -1:
                        all_cubes[color_cube] = all_cubes[color_cube] + number_of_cubes

                if max_in_cubes[0] < all_cubes["red"]:
                    max_in_cubes[0] = all_cubes["red"]
                if max_in_cubes[1] < all_cubes["blue"]:
                    max_in_cubes[1] = all_cubes["blue"]
                if max_in_cubes[2] < all_cubes["green"]:
                    max_in_cubes[2] = all_cubes["green"]

        power_of_cubes.append(max_in_cubes[0] * max_in_cubes[1] * max_in_cubes[2])

    print(sum(power_of_cubes))

day_2()

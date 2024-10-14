import re

file = open("input.txt", "r")
lines = file.readlines()

numbers_in_words = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
numbers_in_numbers = ["1", "2", "3", "4", "5", "6", "7", "8" ,"9"]

sum = 0
pattern = "(?=(one|two|three|four|five|six|seven|eight|nine|[0-9]))"
for line in lines:
    match = re.findall(pattern, line)
    match1 = match[0]
    match2 = match[len(match) - 1]

    if match1 in numbers_in_words:
        idx = numbers_in_words.index(match1)
        match1 = numbers_in_numbers[idx]
    if match2 in numbers_in_words:
        idx = numbers_in_words.index(match2)
        match2 = numbers_in_numbers[idx]

    num = int(match1 + match2)
    sum = sum + num
print(sum)

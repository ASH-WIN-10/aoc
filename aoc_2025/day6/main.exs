defmodule Solution do
  def part1(nums, operators) do
    len = length(operators)

    Enum.reduce(0..(len - 1), 0, fn i, acc ->
      column = Enum.map(nums, fn row -> Enum.at(row, i) end)
      operator = Enum.at(operators, i)

      acc +
        if operator == "+",
          do: Enum.sum(column),
          else: Enum.product(column)
    end)
  end

  def extract_nums(contents) do
    len = String.length(hd(contents))

    Enum.reduce((len - 1)..0//-1, "", fn i, str ->
      column =
        Enum.map(contents, fn row -> String.at(row, i) end)
        |> Enum.join()
        |> String.trim()

      if column == "", do: column <> "  " <> str, else: column <> " " <> str
    end)
    |> String.split("  ")
    |> Enum.map(fn x ->
      String.split(x) |> Enum.map(&String.to_integer/1)
    end)
  end

  def part2(contents) do
    {nums_str, operators} = Enum.split(contents, -1)
    [operators] = Enum.map(operators, &String.split/1)
    len = length(operators)

    nums = extract_nums(nums_str)

    Enum.reduce(0..(len - 1), 0, fn i, acc ->
      column = Enum.at(nums, i)
      operator = Enum.at(operators, i)

      acc +
        if operator == "+",
          do: Enum.sum(column),
          else: Enum.product(column)
    end)
  end
end

contents =
  File.read!("day6/input.txt")
  |> String.split("\n", trim: true)

{nums, [operators]} =
  contents
  |> Enum.map(&String.split/1)
  |> Enum.split(-1)
  |> then(fn {first, last} ->
    {Enum.map(first, &Enum.map(&1, fn x -> String.to_integer(x) end)), last}
  end)

IO.puts("Part 1: #{Solution.part1(nums, operators)}")
IO.puts("Part 2: #{Solution.part2(contents)}")

defmodule Solution do
  def calculate(nums, operators) do
    Enum.zip(nums, operators)
    |> Enum.reduce(0, fn {group, op}, acc ->
      acc +
        case op do
          "+" -> Enum.sum(group)
          _ -> Enum.product(group)
        end
    end)
  end

  def part1(contents) do
    {rows, [operators]} =
      contents
      |> Enum.map(&String.split/1)
      |> Enum.split(-1)

    nums =
      Enum.map(rows, &Enum.map(&1, fn x -> String.to_integer(x) end))
      |> Enum.zip()
      |> Enum.map(&Tuple.to_list/1)

    calculate(nums, operators)
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
    |> Enum.map(fn x -> String.split(x) |> Enum.map(&String.to_integer/1) end)
  end

  def part2(contents) do
    {nums_str, [operators]} = Enum.split(contents, -1)
    operators = String.split(operators)
    nums = extract_nums(nums_str)

    calculate(nums, operators)
  end
end

contents =
  File.read!("day6/input.txt")
  |> String.split("\n", trim: true)

IO.puts("Part 1: #{Solution.part1(contents)}")
IO.puts("Part 2: #{Solution.part2(contents)}")

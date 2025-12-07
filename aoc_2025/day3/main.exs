defmodule Solution do
  defp pick_largest(batteries, drop) do
    batteries
    |> Enum.drop(-drop)
    |> Enum.max()
  end

  defp pick_largest_after(batteries, battery) do
    {_, [_ | rest]} = Enum.split_while(batteries, &(&1 != battery))
    pick_largest(rest, 0)
  end

  def part1(contents) do
    contents
    |> Enum.reduce(0, fn line, totalJoltage ->
      batteries = String.codepoints(line) |> Enum.map(&String.to_integer/1)
      first = pick_largest(batteries, 1)
      second = pick_largest_after(batteries, first)
      totalJoltage + (first * 10 + second)
    end)
  end

  defp pick_batteries(picked, _batteries, _drop, 12), do: picked

  defp pick_batteries(picked, batteries, drop, count) do
    battery = pick_largest(batteries, drop)
    {_, [_ | remaining]} = Enum.split_while(batteries, &(&1 != battery))
    pick_batteries([battery | picked], remaining, drop - 1, count + 1)
  end

  def part2(contents) do
    contents
    |> Enum.reduce(0, fn line, totalJoltage ->
      batteries = String.codepoints(line) |> Enum.map(&String.to_integer/1)

      picked =
        pick_batteries([], batteries, 11, 0)
        |> Enum.reverse()
        |> Integer.undigits()

      totalJoltage + picked
    end)
  end
end

# contents = File.read!("day3/input_test.txt") |> String.split("\n", trim: true)
contents = File.read!("day3/input.txt") |> String.split("\n", trim: true)

IO.puts("Part 1: #{Solution.part1(contents)}")
IO.puts("Part 2: #{Solution.part2(contents)}")

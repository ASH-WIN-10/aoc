defmodule Solution do
  def part1(contents) do
    Enum.reduce(contents, {0, 50}, fn <<dir::binary-1, dist_str::binary>>, {count, dial} ->
      dist = String.to_integer(dist_str)

      offset = if dir == "R", do: dist, else: -dist
      new_dial = rem(dial + offset + 100, 100)

      {count + if(new_dial == 0, do: 1, else: 0), new_dial}
    end)
    |> elem(0)
  end

  def part2(contents) do
    Enum.reduce(contents, {0, 50}, fn <<dir::binary-1, dist_str::binary>>, {count, dial} ->
      dist = String.to_integer(dist_str)

      offset = rem(if(dir == "R", do: dist, else: -dist), 100)
      new_dial = rem(dial + offset + 100, 100)

      wraps = div(dist, 100)

      extra_crossings =
        case dir do
          "R" when dial != 0 and dial + offset >= 100 -> 1
          "L" when dial != 0 and dial + offset <= 0 -> 1
          _ -> 0
        end

      {count + wraps + extra_crossings, new_dial}
    end)
    |> elem(0)
  end
end

# contents = File.read!("day1/input_test.txt") |> String.split("\n", trim: true)
contents = File.read!("day1/input.txt") |> String.split("\n", trim: true)
IO.puts("Part 1: #{Solution.part1(contents)}")
IO.puts("Part 2: #{Solution.part2(contents)}")

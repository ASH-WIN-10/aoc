defmodule Solution do
  defp check_id(id) do
    id_str = to_string(id)
    half = div(String.length(id_str), 2)
    String.duplicate(String.slice(id_str, 0, half), 2) == id_str
  end

  defp check_id_2(id) do
    id_str = to_string(id)
    doubled = id_str <> id_str
    substring = String.slice(doubled, 1..-2//1)
    String.contains?(substring, id_str)
  end

  defp find_invalid_ids(range, fun) do
    [first_id, last_id] =
      String.split(range, "-")
      |> then(fn [a, b] -> [String.to_integer(a), String.to_integer(b)] end)

    first_id..last_id
    |> Stream.filter(fun)
    |> Enum.sum()
  end

  def part1(contents) do
    contents
    |> Enum.map(&find_invalid_ids(&1, fn id -> check_id(id) end))
    |> Enum.sum()
  end

  def part2(contents) do
    contents
    |> Enum.map(&find_invalid_ids(&1, fn id -> check_id(id) or check_id_2(id) end))
    |> Enum.sum()
  end
end

# contents =
#   File.read!("day2/input_test.txt")
#   |> String.split([",", "\n"], trim: true)

contents =
  File.read!("day2/input.txt")
  |> String.split([",", "\n"], trim: true)

IO.puts("Part 1: #{Solution.part1(contents)}")
IO.puts("Part 2: #{Solution.part2(contents)}")

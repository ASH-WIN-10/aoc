defmodule Solution do
  def part1(ranges, ids) do
    Enum.count(ids, fn id ->
      Enum.any?(ranges, fn [start, finish] ->
        id >= start and id <= finish
      end)
    end)
  end

  def part2(ranges) do
    ranges = Enum.sort(ranges)

    Enum.reduce(tl(ranges), [hd(ranges)], fn curr, res ->
      [last_start, last_finish] = hd(res)
      [start, finish] = curr

      if start <= last_finish do
        [[last_start, max(last_finish, finish)] | tl(res)]
      else
        [curr | res]
      end
    end)
    |> Enum.reduce(0, fn [start, finish], acc -> acc + (finish - start + 1) end)
  end
end

[ranges, ids] =
  File.read!("day5/input.txt")
  |> String.split("\n\n", trim: true)
  |> Enum.map(&String.split(&1, "\n", trim: true))
  |> then(fn [ranges, ids] ->
    [
      Enum.map(ranges, fn range ->
        String.split(range, "-") |> Enum.map(&String.to_integer/1)
      end),
      Enum.map(ids, &String.to_integer/1)
    ]
  end)

IO.puts("Part 1: #{Solution.part1(ranges, ids)}")
IO.puts("Part 2: #{Solution.part2(ranges)}")

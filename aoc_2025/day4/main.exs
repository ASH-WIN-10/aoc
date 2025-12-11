defmodule Solution do
  @adjacent_pos [
    {-1, -1},
    {0, -1},
    {1, -1},
    {-1, 0},
    {1, 0},
    {-1, 1},
    {0, 1},
    {1, 1}
  ]

  defp get_next_pos(%{x: x, y: y}, width) when x < width - 1, do: %{x: x + 1, y: y}
  defp get_next_pos(%{y: y}, _width), do: %{x: 0, y: y + 1}

  defp get_pos_val(grid, %{x: x, y: y}), do: String.at(elem(grid, y), x)

  defp is_out_of_bounds(%{x: x, y: y}, width, height),
    do:
      x < 0 or x >= width or
        y < 0 or y >= height

  defp get_adjacent_rolls(grid, curr_pos, removed_rolls, width, height) do
    Enum.reduce(@adjacent_pos, 0, fn {dx, dy}, acc ->
      pos = %{x: curr_pos.x + dx, y: curr_pos.y + dy}

      if is_out_of_bounds(pos, width, height) do
        acc
      else
        case get_pos_val(grid, pos) do
          "@" -> if MapSet.member?(removed_rolls, pos), do: acc, else: acc + 1
          _ -> acc
        end
      end
    end)
  end

  defp walk(grid, curr_pos, accessible_rolls, removed_rolls, width, height) do
    cond do
      is_out_of_bounds(curr_pos, width, height) ->
        accessible_rolls

      get_pos_val(grid, curr_pos) == "@" and
        not MapSet.member?(removed_rolls, curr_pos) and
          get_adjacent_rolls(grid, curr_pos, removed_rolls, width, height) < 4 ->
        walk(
          grid,
          get_next_pos(curr_pos, width),
          MapSet.put(accessible_rolls, curr_pos),
          removed_rolls,
          width,
          height
        )

      true ->
        walk(grid, get_next_pos(curr_pos, width), accessible_rolls, removed_rolls, width, height)
    end
  end

  def process_part2(grid, removed_rolls, width, height) do
    accessible_rolls = walk(grid, %{x: 0, y: 0}, MapSet.new(), removed_rolls, width, height)

    if MapSet.size(accessible_rolls) == 0 do
      MapSet.size(removed_rolls)
    else
      removed_rolls = MapSet.union(removed_rolls, accessible_rolls)
      process_part2(grid, removed_rolls, width, height)
    end
  end

  def part1(contents) do
    grid = List.to_tuple(contents)
    width = String.length(elem(grid, 0))
    height = tuple_size(grid)

    walk(grid, %{x: 0, y: 0}, MapSet.new(), MapSet.new(), width, height)
    |> MapSet.size()
  end

  def part2(contents) do
    grid = List.to_tuple(contents)
    width = String.length(elem(grid, 0))
    height = tuple_size(grid)
    removed_rolls = MapSet.new()

    process_part2(grid, removed_rolls, width, height)
  end
end

# contents = File.read!("day4/input_test.txt") |> String.split("\n", trim: true)
contents = File.read!("day4/input.txt") |> String.split("\n", trim: true)

IO.puts("Part 1: #{Solution.part1(contents)}")
IO.puts("Part 2: #{Solution.part2(contents)}")

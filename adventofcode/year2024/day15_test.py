import pytest

from adventofcode.library import testit
from adventofcode.year2024 import day15

example = """
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
"""
answer = """
##################################################
##OO............#..........O.O..O.O.....O..O..O..#
#O...#.O.O.O....#OO.#O....#.........O.#..........#
#O....O..O.OO.OOO.#.....O#.....#..#O.OO..........#
#.O#OOO...O............O##......#..O..O......O..##
#......#....#...OO....O.#OO..O.OO.OO.O#........#.#
#...O...##.....OO...O..#...O....O....O.....O...O.#
##.O.#...##....O..............#....O....O.O..O#.O#
#.OO#...O........O......#.#O...OO.OO...O.O.#....##
#O....O...O.O.O.#.#OO.O#.O....#..O..O...#....OO..#
#O.O..OO..##.....#.#O..O.O.OOO.....O..#.O..#.....#
#.....OO..O...O.OO..#......O...O...O..O.#........#
#..O..O#.O.O.O..#.O.O..OO....#....O#.....O.OO..O.#
#..#O#..O..O.....#.#O...O..OO#.OO..OO...........O#
#..............O#O..#OO..OO...OOO.....O#..#...O.##
#....OO..#...O..#..O.#..OOO........#OO.O.O..O....#
#O#.#.O.O.OO..O.....OO..##....O..O........#.O....#
#.O.O.O.O#O..#.....O#.OO.#O.....O.O......O..O....#
#O.....O#O...O#....OO.OO.......O.O...O.......O...#
#....OO....O.O..O.#.......OOOO...O.O..OO....OO...#
##.....O....O..O.O..OOOO.....#...OO.....O#...#..O#
#O.#...O...O.O.O.OO.O...#....O..O.OO..O...O#..O.O#
#....O.#O.OOOOO..OO...OOO.#OO.O...O...O...O..O...#
##O.....#O.....O....O..O.O...O...O.O..OO.O.O...O.#
#......#.O#.O...#....OO.@#.....##..#.OO#....#....#
#OOO...O..OO.........O.#.O....O#...O.OO....#O...O#
#...OO..O..O#...O..O.#O.O.O#............O......O.#
#OO.O.OO..OO.#....O....O.....OO....O..O..O.OO...O#
#..O.O.O...OO.#..OO..#O....O...#...OO...O.O.O#OO.#
#O.OO#..O..........OO....#...#..#.#..O........O..#
#..##..OOO...#.#.....##.O..O..#..#O.......O#.O.OO#
#O.O.#.O.O..#.OO.#O..O...O..#..OO.OO..O.O..OOOOO.#
#.#O..OOO#..#.....O.O...O...O.OO....OO.OO#O#O#O#O#
#...O.O.#...#O....O.OOO...O.#.....O..O.O.......OO#
#..OO.......O.O..OO#.....OOOOO.....O.OO..#O...OO##
#.....O..#.O.......O...O....OO.O.........O.......#
#..O.O.........O....O.#O.....#....#.O....#O......#
#..##....O.O...O..O.O..##O...#....O#O.......OO#O##
##OOOO.O.O.O.OO.........O..O........O#O..O....#OO#
#..OO#...O...OO.....#.OOO....OO.OOO..O.O..OOO.O..#
#O..OO..O..OO............O.#..#O..O...O..##OOOO#O#
#...O..O#..O.O....O..O...O....#.#O.OO.O...#O.OOO.#
#......O.....O...OO..##.O.O...O....O...O....O....#
#O.O.#O.#O..O..OO.#O#..O..O..O.#.O...O.O...OO.O..#
#.OO....#OO.O...O.O.#...#.O.O..O#......O........O#
#....O.#O.OOO....O.O#..O.O.#..O.....O.OOO.O.#.##.#
#..#O..#.....OO....O.O.O.O.O..#..O.O#..OO.....O..#
#...O.............OO..#....#O......O..O.OO.O#.O.O#
#...#..O..O...O.#O.O.O.....OOO..OO#......O..OOO..#
##################################################

>v^>^^^^^>>^^vvvv<v<<v^^^<>v^^v^<^>>^v^^v>>v<^<>v>v^v^v>>>^<v>v<vv^<^^>><^^v<<<^>>>><<<>>^^><^>v<>v>><vv>^v<^vv<vv^<>><v<^vv<v^>^v^^<>^>v^vv<v<^v^<<v<>^v<<vv^^<vvv>^v<>^<^v^>>v<>>>>v<<<>v<<^^<<v>v^<^><>v>^vv><<^^><<<<v^v>>^v^^vvv^v>v>^<>^<<<v<<<<>vvv>v<>><vv<^vv^vvv<^^>>vv><v^^><<^>>v<v^^>v^^<<^^>^><<>>^v<>>^>vv>^v<v<<^>><vv<<^^v^^<v>vv><vvv<>v<v<^<<<vvvv^vv^>><^<<^><><^v>>v>>>><<v^<^v<^>v<^v>>>^>v>^^v^><^^>>^>v>v<v<<>v>^<<v^>^>^><v^<<v><v^><^<>>>>^>vv^^vvv>v^><<<v>^>v><^>v>>><v^>^>>v^>>^v^^v>^vv^^^v>>^<>>vv^<<<<<^v<>>vvv><><<v><>>v<<>>v^>>^<^v<v^v^>v>>>vv<>v>v^v><v>vvv<v<<><v^<<^>^^^<>^^v^vv>v<^>v<<^<^<<>v^<<><v<v^v^v^>^<<v^<<v^><v^><v>>v^vvvv^^vvv>v^^^><>vv<<<>>><>v<<>>vv>vv^><<<v^>>^vv>v^v>^v<<<vvvv<>v^<>^^>^>vv<^v>^v>>>>v^v<^^>^<><>v<<>>^>v><<><<><<^>>v><>v>^<v<v<v><>>v<>^<v<>>v<<>>v>>><^><>>vv<v^<^>vv^v><vvv>^vvv<><><<^^>v<>>>><^>^<v^v<v>>^>><><<vv>>>>><>>>>^<vv^>^<<<v>vvv>v>>v^^v^^^>^<vv>v<<<<<<<^>v^^vv<<<v^<vvv><<<<<v>>>vv<v<^>>^^<<^<<^>>^v^>^v<<^>^>vv>>v>><<>>v<^^<>v>v>v<><^><>
>vv<<<v^v<>^v^>^vv^<^v<<v>^v<^><>^^>>vv^<vvv<^>^>>>>>v<<^<>vvvv><<>^^>><><v>v^vv^<>^<><vv>^^<><v<v<<>v^>>>><vv>^<><^<<^>^>^^^><>>vv<^>v<<^v><^^>^v>v<v<>v^v>^v^>><vv<>^>^^<>^><v><^v><><<^<v<<>^<v^^<^><>>^>v^v^^^v<^v^><>>^v^><vv^<^<^<vv>>v^vvvv^vv<<<v<<^>^vv<^<v<<<^<^v^><><^>^^v<<<>><<v<<^<>><<vv><vvvv><^v<^>v<<v>v<^>v>v<<^^v<^^>v^v^^^v<<>><<>v<><>v^^v^v>^><vv>>>>v><><><<^^v<<v<v<vv>^<>v<>v^>vvv>>v^>>>^vvv^<^v<>>>>>^<^v>vv>><vv^v^<v^<<>>^<vv<v<v<><<<v><^>v^>^v<>^^v><v><^^^>>^<^v>^^^<><v<>>^<<<^<>v^^><>v^<^^vv<v^^^<^^v>v><<vv<^v<<>^v<v>v<v>^v>>>>v>^v>^v<>^v><^vv^v<vv>^^^^>^^v<vvv^<>>^^>>>^^^>v<v<<><^<<v><<>^><<^vv<^>vv<^><>^^>>^<^^>>>><<^v^<v<<<^^v<^v>>^>>v^>^>><<v>^>^<^^v>v^<v^^<<<>vv<^>^<^><>^<<<v<<^vvv^^v<>vv><vv^>>^^<>^vv<<^><^^^<>^^>>^><><><<><^<>^>><vv<v>^>vv>v><v><>v><vv^v^>v<><vv>>>vvv<>v^<><vvv^v>>v^><><^>^^v^<^^>v^<v<^>^<^v^v>^>>>><<v><<v^<^>^^>vv>>^^^<>^v<>><><^vv^>v<<^v<v^^^>>>>^>v<>^<^vv<<>^v>^^v>v><^<vv><<>>^<vv><><^v^v>^^<>^<>^vv^^<v^v>^^<v<<^vv><v>v^^v>v^<>^>>^v<v>v><>^v>>
>^<v^<<<^vv>>><^<^^<>v>>vvvv<v^>>^<<^>>>v<vvv^<<<<<<v>v^><<v<<>v>><vv^^^>^><vvv<v>>v>v<>>>^<^vv^>>^>v>^v^vv^^^v<<>v<<<^>v>v^>>v^vvv<^<v^>^<<<v>^v<>>vvvv^>v^<<v^>^^<>>^<^v^<v><^^<>>>v<<<^vv<^>>><<^^^<>vv><>^^v^v<<><v<v>v<^^^v>>^v<^v>vv>^^^>^v<<^<vv^^<>v<^>^v>v<>vvv<^<>^<<vvvvv^<><<vvv>>vv>>>>><^vv^v>>v<>><v^>vvv<><v^>>v^>v^v^^^vv>^v<<>vv>v>^v^<<v>v^<^v><>>>v^<<>v<>>vv^<><><<v>v><v<<<vv<^^v<<<<^>>^<>^v>>>>>>vv^v>>>v<>^v>v^>^v^>vv^^>>>>^v<v^v<^><v<v^<>v<v<v^^^^><^>>vv^<<>^<^>v>vv<>^>>>><<<>^vv>vvv>v^^>><v^<v<v<vv<^^v^>^>v<>vv<>><v<v^>v^<<^v<<v>v<>v^v^>^^^<<<^^v^vv<<^<>^>v>vv<v^^^v<>v><>>><>v<<v><>vv^><>^v<><^v^<<>v^>^>v<<<v<v>v>^<>>^^^v<><>v^v^v<>>v<><><<><v<^v<v^>>v<^^v><>>v^vv<<^vv>>><>v^>v<<<v>>^v^<><<><^>>><>^><>><^>^v<v<^>><>^<<^>v>^^^>v<v^^<^>^^<<>^vv<^v>>>v<vvv>^<v^^^v^>^<<><<v><^<^^<^v>vv><<^<^<<<v^^^<<^v><^^v<<>^<<<^vv^v<v^<vv><<v^>v<>v^^^^^^<>^>^<v^vv>>v^><v<<>v^^>v>^v>>^><>>^>v><^<v<>^v^>v>>>v^vv<>>^v^>>vv^^v^<<<>^>v<>v<^<^>^v>v>^>><vv^>v<>^^<^v>v^><v<>^^<><<^^vv>v><<^v>>><^^^v
>><v^v^v>vv<>^>>>^vvv<^^^<><<><^vv<^^^<v>^><>v^<v>>>^>^vv>>^^v><>^<>v<>^^>vvvvv<v>>v<>v<<v>^^>>>^v^v>>>>><<^<^vv^v>><v>^<^v^^^^v>^>v<v^v<>^v^<^v>^<^v<>^><><<^>v>v<><v^<^^^vvv^v^><<vv><v>><<<v><^^v>>><<<vv<v^^><^>>>v<v<^v<>><^^^>vv^<^>>v<>>>v^><v^v^^^<v>^<^>^v><v<v<^<^v>^vv<v>^^^vv^><>v<^^><>><^>><>^v^>^v>v>vv^v>^^^>>v>^>^v><<v><^><v^vv>^^^v^<><v^<^^<<<vv^<><><v>v<^v<v<<v<<vvvv<^>^<><<v>^v<vv^^<v^>v^^<<^<^<<<>^^<v<<v^<><v>^^>v<><<^>v>><^>v<^>^><<<v<v^v<<<<^v^^^v>v<v>v><<^><v^^<<^v<v^<^<><>^>>v><<<><v^>v<>^^<^>^<>v>^v>v>vv>>>vv>^^^>vvv<^^vvvv^v>^v^<^<^<<^^>>v><>^>>>v>>>v>>^v^><v>^v>v^v^>^>v^^<^><v>>>v^v>^<><>><^<vv><<^<^^vv>v^><v>>^^^<>v>v^^v<<^>>v>^v^vv><v^^vvv>^^<>>v<v^vvv^<>^vv^>^>>^^<^v^^^^v>vv^^vv>v<v^^v<<>^>^v^<v^>>><v>^>>^><>>vv^>>>^v>v^^<>>>v>v>^>v^v>^v<<<^v<v><><>^vv>v<<v^>>^><<v<<^>^^^v^^v^>^<>v<>^vv^v^vvv^vv>^^^<>^<<^^^v<<>^>vv^^^<^><<<<<^^>^<>^^>^><>^<<>v>>>v^vv^<v<<<^><vv<>^>v><<^>vv^<^<^v<>^<<<>v>>v>>vv<v^<v>^v^vvvv^v^>^v>^v><^<>><^vvv^<>^>>v<>vv<^<^<v>vv^<^><<v^<^vv<^v<<v<
<^v^^v^<<>v>>>vv<^>v>>vv>>^v^<^vv^^>>>^>>vv^>>vvv^<^vv>v><<<<v^<vv^v^>v^<^>^<<<v<^<v<><^^vvv>>>^<>vv>v>^vv>>v^^v<v^^v<^><^^<^<^v<>v^^v^v><><^<v^^^<^<>v<v<<vv^>v>>><^<v>^<^v<v<<v<^<v>^v><^v^<><^v^><v<>>v><<>>^><vvv^v<vv<vv>>v>>><v<^<<v^v>v^<^<v>>v<v<>vvv^><>>v<^<^^^^^>>v<v<^^vvvvv^<v<vv<<v<^^<^^vv^^v<<^vv>^>^vv^>vv>>v^<<^><^v^>v>>>>v<<v>v^^<^><v<^v^<^v<^<><^<><v<^^>^<^^^>^<>v>^^v^>>^v><<v>>v<<>><>v<^^v><><<^v^^v^>vv^><v<<>^<<^><<>^^<^^vv<><<<><^v<<<><><^<vvv>^^<>vv>v^^>^v^<><><<><^^>v<^v<>v<^<v^v><>>>^<>>>v^v^v<<^<<<<v<<<>v<><<v>^^^vv<>>^^v<<^>v<<^<vv<v<<^v><^v<vv^>>^><^^v<^<^^>^>^v>^<<vv^^>>v>v<<<><>>^vv<vv<<v<v>>^>>vv^>>v^v<<<v>vv^<v^^>^>>v>>>>>><v<><<^>v^>^^<>vv>><vv>^><><>v>^<v<><<<>>>>^v<^>^>v<><vvvv<^<vv>>^>>v<v>^^vvvv><<<^^>^<v^^><<^>^^>><vv>>^>^>>>vv^v<<><><>>vv>>><>vv<^><>>>>^>>><<<><vv<v<<^<^<vv><>v>v><vvv^v^v><<^^vv^^v><>^^<^>^>^^vv^v>vv<<^<^v^>^<vv^^<v>v<><<><<v^<<^>>vv<<^v^v^^vv^vv<v>>^<^^<<<>vvvv^><vv<v<v>vvv<v^v>v<>>>^<^<<^^^>>v>v><^>>^^vv^<v<vv>vvv<^><^^v^>>^>v^^^>v^<>>>
vv^>>^>><^<^^><<^>><>^>><>>vvvv<^>>^^v^><v>v<^>><>^<^v>^^v<>^^^><<>^>>^<>>^^<<<v<<><><<v>^^<vv>>v^><v<<<v>^<<<vv<<v<vv>><^><><^>>>>^>>vv^><>>><<>^><>>v<v<v<>>^^^vv<^>^v<<>^<^^<^v<v^><<v>>v<>v^^^vv^v^v>^>>vv>v>>^vv^>>v^>^^^>>>>v>^v^>v>v^v<>^<v<<><v<vv<v>^^><^v<<>^<^vv>v^^<v^<>v<<vvv><v<^^vvv>vvv<^vv^^<>><^^<>vv<vv>>>^v>^^^^>>vv<v>^>^>v>^^<^^^^^<<<>^<<v^<^><>vvv>>^v>^v<>>v<v<>v><^>>^<<v^<<^^^>v>vv><<^<^v>>><v^^^>>v><>^v^v><>>vv^<<v>>>v^<><>vv><^>><^><^<<>v>>><>>v><<^<v>vvv>v>^vvvv^<^^<v><>v<^^<>>^<<<v^v^>vv<<^^v>>v<<v<v>^<v>><^vv<^^v>vv^><^<v<v<vv>^>v^vv<^v<v<v^vvv><<^>^>^^v^<^^><v^^<^v^v>v>>^<<><^>v>><<>^^^>^^vv>><^<>v<^<vv<<>>^<>vvv^>>^v<^>^^<^v<<>^<><>v><>^vv^^^<>v>>^v>^>vvvv><v^<><^><v<v^vvv>v>^<<v^v>>v<>v>>v^><^>v<<<^^^<v<^>v>^^<<>^v^>^>v^><>>vvvv>>v><^>v^v<><>v<v><<^^>v^<^^^^v<^v>vvvv^>^>^<v>v^v^^>>>^<<^<v>>><vvvv<><vv<<^v<<<v^>>^>^v^>^v<<>v<<^>^>>v^^<^<v^<v<^v<>>^>>>>vv<<^>^<v>><<<v>><^<^>v><vv^v<^^>^v<v^>v<^><<>^>^vvv>v^<vv<>v<>>^<>>^v<v^^^<<>>^^>v<^>v<<^<>>>>vvvv^v<v<<<v<^>vv^>v
^>^v>>vv>>^<<^><<^^v^<<<^^v^>><v<^<>vv<<>^>vv<v>^v<><>^v>><<v>^>>^>v>vvv>>^>^^<^<^v>^^<v<<^<vv><^>>v>v<>^<v>v^<<<<^>v<<^<<^>>v>vv<^<^v>vv>v^>^<^<v<>vvv^v>>^v^v>><>vv<><^<<>vv>>vv>^v^vvvvv>v><>^^^v<^<^^<<v<v<^v<><^v<vv^vv>^<<<vvv><^<^v^>^<vvv<^>>v>>><<v<^v^<>^<><^v^<<<<>^v>>^<v><><<^><<>^^>^^v>vv>^<vv<^<>vv^<<^<^<><v>>v^<^^^<>^v^^>^>v^v>^>>v<^<v^vv<<^v<<<<vv>vv^>^^<>vv<v>^<>^^^<<^<^><<<><<^><>>>^vv<v>v^<vv<<^v>>>>^>v<>^^<><v<>>^<vv^>v>>>^<>>^vvv^vvv>^<^v>>^>vv<>>v<<>^><<v><<><<<v<>^v>^^>vv<><v^>v^>v^>><v^<v^^>v<^<vv><^<<v>^><v<<<<<v>^v<<<^v<>>><>>^<>v<<>^>>v<v^<><^v<<>v<<^^^^^^>^v^><<^<vv<>^>>>^<v^v<^^>^v^>>^v>^<v<v^<><>^v^vv^<>^vv^^>^<v^^<^>vv^<>^v^<vv<v<v<>^v><<<<<<><v<<^<^v^v<^v><>vv>v^<<v^^>^>v^>^^>^<^vvv^>>^^^<^v<<>v><v^^v<v>^>v>><^vvv^>^vv<v^>^>v<<^^>vv>^<v^<v<vv^^<^<<>vv<^<vv^^^^v^^>vv<>>^^><^^^>v<>^<<>><v<<<v^^<v><>^><>^>v^<<<v<v<^<<<<^v<vv>vvv><vv>^<^>>^<v>^<v<vv^<^v><v^v>^v<vv<<<v>>>vv^>>v<>v^v<<vv>^^v<>^^>v<<<>v>v<^v^vv<<>>><<v<v>v<^<^^^>^^<<>>^v^>^^>v<<<><^^<><<vvvv>>^v>>>vv
><^<v<^^^<v<>^><<^^v^<<>>><^>v<<<<<<<v<><>><v^>v<>v>>>v<>>v<>^<><v^<>^>>v<^><^>^<^<<><><><<^<^>>><v^v^^v^>>^<^<^v<v>v^^<<>^>^v<>>v^v^v<^^^v><^><^vv^>>^vvv>^^<<<<<^<>^^><^<^^v><^vvvv<>>^v^v^<^>>>>^<<<vv<<^^<^>vv>>^<>^^<^><<^^v<vv^>^^<v<^^>v^^^>>>>^>v^v>^^^vvvv^<v<^<<>>^^^<v^vv^vvv^>>>>^v><^^><><>^v<^>^^>>>^vv><v>>v<v<^<>^^^^<<vv><^vv^vv>^>>^<<<^^vv><>vv^>>^<^^^>><<^<>v><v>v^<^>vv<<v^>^>^v<<<vv>v<>vv<^<v><<v>v>vv>>^>v^>v>v<^vv<^vv<^<^^vv>v^^<^^<^^vv<^><<^vvv<<v^^^v<<^v>v>^v<>><v^><v><>>vv^<>^<>^<><<>>^^v<^>v<><<<>><vv^v^<v<>^v<v>^^<<v^<>><v<<^^<v^^v<>>^^^v>>>v><<^^^<<v<<>>>^>v<<>vv>^^^v^>v^v^^>>^^<^^<>>^^vv<<v<>>v^^vvv>><>^<^>^^>^^>vvvv<<<<<vv<^><<v^v><>v<^>>v>><>^^v>>^>^>>vv^^>>^v<v^v>^>^^v<v>v>>^><<v><>>>^<>>^v^<^^v<<><>^>^><v<^^^v^><><v>^>>^^>v^>^<<<>^<<><^^<v<v^><<^>^<>><vvv^>v<>>>^v<^^<vv^vv<><<^<>v^>^>^>^<^<><v>^^<><^v^<v><^v>vv^^^^>v><<>^<vv^>>><^^v^v<v>>v<>>v><>v>^<v<v>v<v<>v^<<>^^^vv>>v<>^<>v^v>^><v^>^^^v^v^^^^vv<v<<<<>v^>><<<>^>^v^^<^<>^^^v^^<v^^vvv^^^v<^v<>>vv>>vv^><v^>^>v<<>^
v><^^<<>^^vv>^><><^<^>^vv<<><^^v>^v^>v><<vvvvvv>^v>v<>^<^v<<>vv<v>v<v^>^<^>>^<<<>><^v>^<<<<^vv><v>v^>>><v^<><^v^<v^^>vvvv<^><<^<^>>>^^vv>v<<v>>vvvv<^><v<><>>v<>v><<>>v<><<^vv<<vv<^>>>>^>><>>><<><v^<>^vv><>^^<v^^<>^v<<^v>^<^>^<>>^<v<v>^>><>vv>^><vv^><^<^v<^^vv^>>^>^^v><^<v<>><^<^<v<^vv^^>v^vv>^<v<<>v<>v^v>><v<<><>^<><><>>v^v><>^<^<>^><v<>^>^><<<^v<v>v^>^>^<><<<<><vv<v>^v^v<<>><<v<^v^<><<><v<vv>^^v<><v>^v<v<^^v^<<>>>^vv^>>>vv>vv><^^v><>>>>v<v><v<v^v<^<>v<^v>><<<^v<><<>vv^v<>>v<v<><>^vv<<>v^><v<<<<><>v<<v<<v<v^v<^^v^^v<^v^^>v<<vv<<^v<<><<vvv><<v^<<v<<<><v>>^<>v^<^v>^v>v>^v>v<<v<v<^<v<^>v<>^<>>^vv>vv<^<<<<^^<^v^v<^>v>^<<v^><>^<>^>>>>>v^v>^^^v^^^v>vv^<<^^^>^^<v><><>v>>><vv<>vvvv>vvv<^v><><v^>v<>^v^^^v<^<>vvv>^^>^<>vv^^<<^>>^<v^^vv^^v<^>^^^^>vvvv<><^>>v>vv>>vvv^^<>vv>>v^v^v><v>vv<<<<><><vvv^><><<^^<<>^v^v<>v^^^<v^>^>vv^v<<<vvv>v>vv<<v<><><^v>^>^<<v>^^v>v>^<^>^>^^<>>^vv>v^>v<>^>^>v^^^v<>^<v>vv<^<^^<<<vv>v<^^v<>>^<>><>vvv>^<>><><><>^vvv>><><v<v^^^<v^v^>^^><^<vvv^v^><>^^<<>^>>v^^>^^>><v^<<v<<v>
^v<<<>^>^^<^v>v>^v>vv<>v<^>^^^><^^><>^>v<>v><<<^><^<^<<<<v^>^<^v>>>^>^^<vv<>>vvvv<<<<v<<v>v<><<v^v<<>^^<v<^v>^^><<vvv><>^vv^>>^>v>>v<<<vvv<^><^><^v>v^v>v^v>^vv>v<>><vvvvv^>^><<>>><^v>v>>v^^^^^>v<v>v>>^^><v^v<>v^><><^<>><v^>^<<vv><><^>>^><^>vv^>><^<<vv>v>^<vv>v<>^v>^^<^v>vv>>^v^^^>vvvv^v<<>v><>>^>v<<^<><v^>^>><<^v^<^v<>><<>^^^v<^^>>^^^v<>v<><>^>>^>^<^<<^^v><^<v^vv<^^>vvv<>>>>^<vv<>v<>>^v<^v>v^^><<v<><v>><>vv^v^<>^<<><v<^v<vv^<>>>^<<^^v^<<v<vv<<v<v>vvvv<><><^>>vv>v^^>>v>><><v^^>><^v^v<^v><>>^<v<^<<>^^^>^vvv^<^<<vv^^v^v>^>>^><vv<><>>^<v<vv^^v><vv^v>^vv^^<v>>><v<>^<^>v<v^vvvvv<>>><><^<><<vv>v<^>>>v<<<<vv<^^<^>>v>>^<v^>^v><v<<><^><v>^>v>>vv^><^>v^>vv>^<v<v>v^v>^^v<>^>>^^v>>^><^<v<>^v><>v<<<<^^vv^><>><>^><^<v<<<>><vv<<<^<^vvv^vv^^v><<v>vv<v<>v<>>v><^>v^v<<>^<^^>vv^<><>^^^v^v<<<>^>>v^>>^<^^<v>><><^^vv<<v<v<^>^>v^^<><<^<><v>^^<v<v>><^<<^^<^>>v<^>^v>v>^^^^^>>^>v^><v^v>>v>>><>^^>^>>^><v^^<vv^v^<<<v^v<v>><>v<v^>>>>v^v>^<>v^v^v^v^>^>v<v^>^^v>v><v<<<v<>>>^^>vv<<<v^^v>v<v<v>^^<v^^^^^^>>v<<v>^^^<<<>>
<v^<<>^vv<^^<^v^>>^^^vv>vv^<^>><<v^<<v<v^v>^><^>^^^>>vv^v^>v>>v<v<^><^>><>^v^^<v^<>^^vv^^><<<><>^>^>^>^>>^>v>>><v^>^<<>^><<vv^>>v<vv<<<^^^^>^>^^>^vvv<<vv<<><^^^>v<<>^v<<^v>^><^<<>^^<><<v^<^<^v<>vv><>^^>>><><<<^^<^^^>^>><<>^>>><>^><^>^^^<v^^>^><^>v^^>^>^^v>^^>vv<^<^<^vv<<^><^<>^v^v<>>^^v>>v>v<><vv<<v^>>>^v^<v<^<vv^^v<<^>^^v<<<>^v^^v<>v>>vv><<<<<v^v<>>v>^><^vv>^v><<<v>^>>>^>vv>><^<<<<^><><>>>v<^v<v>vv>vv^v^^<v^<<vv<<<<>v>^^<^><vv^><v>v^vv>>><<v>>^v<>>><^<vv>v>>><v^v>vv^<>v^^v<>><v^v>>>^^^^^^<<^>>>^vvvv<<<><v^v^>^^>v>v^^v><><^v<<>^<<<>>>vvv^<>>vv^>>><^>^<><<<>v<v<<>><<^v<v<><>^vv<v>^^^v^v^v>^<<<>>><v<>^v<v^<>^<<v<><<v^^v^>^<<>^^v>v^^^><v>>>>>>>><>v<<>>v<v><<^>v^^^<<>>v^^^><<^>^>><vv^^vv<vv<^><>v<<^v>vv^v<v<<<^><>^^^^v>>>^>><<<>>>>>v>v^><^v>>^<v<<<<<v<>v<v^vv^<>>^^^><<v>^<^<v>^^^v><vv<<^vv>v^^><vv<>^><>><>v>>v>^vv><v<>^<v<><v<><^<^><^>v>>^<>^<<^>^>v<>><^^^v<<>^vv>><<v>^<^<>>>v><v^<v>v<>>^v<vv<>><^v^^v^<<^v<v^<v^^^vvv>>^v^><^><^>><<>vv<v><<>^<<^>v<v<>><^^^>^^vv><<>>><^<<<^>v<<<^v^^vv<<^v>vv
<>^^v<^>><>>>>^vv>^^v^v><v<<<<vv<><><^^<v<>>>>><><<>^<^>v><<<<<v^>v^v<>^^>vvvv^^v^<^v>v^<^^>^^^>>^v<^<<^><^>v>v<>v<v><v<<<v<^>v><>^^v<><v>v<<v>>v<<<^>v^><>>>^<vv>v><>^^<>^>v<v>>>v<^>^<^><>v>><v><>^v<^vv><><<vv>>^^^<>^v>v><>^v<v<<v>^<<>^v>^>>vv<<^v^v><v^^^<v<^>>>v<^^>v^<>>><^<<v><^^^>>^<^<<>>>v><<<^>v<^^>>v>v<v<>^^>>>^v^<>><<v<vvv^v>vv^v>^>v>^<v><>>v>^^<vv<><v^><>^^^<>>>>>v<<>vvv>>><^^>^^<v>v^>v<>v>vv<^v^vv<v^vvv>>^^^v>v>vv><^<>^^vv>^>>^<<v<><^<vv^^^>>v>vv>vv<^vvvv^<^>>v>>>v<<><^v>v<vvv^vv><>^v^v^^v^^>><>>^vv<<^v^v>v^vv><>v><^vv>>vv<<^v>^vv^^^v<vv><<v^vv^v<vvvv<<v^<vvvvv>^<<<^>v<><v<>>^v<<^>v>>vv^<>^>v<vvv<><><><<>^v>v^><^vv>v>>>vv^<>v^<><>>^vv<<v>^<^>><v<<v<v^v<^v>>vv<<<<v>^^v>>><>v<<v^vv^<^<^>v>^<<<<v<^^<<<<^>^<><<^v>^v<>>v<v^<<>v<>^<<>v<>v<v>^v^^>^v^>^>v<>^>^<<<<>>v<^>^<<v<>^^>^<v>>vvv^^<^v<vv>^<vv<v^v<v^^<v<<>><vv<^<<>^<<<v><>>vv^^^^<>>v>vvv^>>>v<<vvv<v<>>><v<^>vv>^v><^><>><<<<vv^>^><^><vv>vv>v^^>>><>><vv<<<^<^><^>^>><^<vv>v<><v><^v><<<^^vv<v^v<v^v^v<<^v>^^<<^v^>v^><v<>>><^<^^<v<>>v
<>v^>vv<^>>><v^^^v>^vvv>^v^^v>><v<<^<v^<>vvv<<^^^>^>v^><^>>^v><v^^^^>^v^>^vv<>v^^^^v<<v>v^<<^<>>v>^v>>^v<<v>>v^<^vvv^^v>^><<<<^><^><>v><v<^>>^vv^vv^v><^<><v^^<>v>^v<>v^v>>v>^><<>^<v<<>><<^v<^^^^<>^^^v^><^<^^<<>v<<vvvv>>v>^vv^v><^^>v>><v<v<v<^<vvv>><vvv>^^^v^<v<vv<^vv><v^<<v^><>^v<<<>^<v^v<^v<v<v>^^>vv>v>v<vv>^^^^>vv<^v<<vvv>^>v<<^>>^^^<v<<>>^^vv<>>vv>^v<>v<v<v>^^<v^^v<v^v^vv<^^^><v<v^^<^^v<v>^<<^v<<^<<><<^<^>v<>>v<<><>^vvvvv^v<v>>v<<^>^^>^^<>^v>v^v<^>v<>v<<<<^>><><>v<v>v<><<<>^>^>>^^>^<v^<^v>^><<v^^<vv^>><>>><<<^^v^>^^<<><>v<>><v<<><>vvv>>v>v>^><><^vv^^>^^^v>v<^><><v^^><>^<^^^vvvv<v>^>>>v<<<^<^^vv><v<v>v<^<^>v<v>>>>^<^><v><>>v><vvv>^v^>^<<<^><vv>^v>v>>v<v><<v>><>v<>^<<v<<>>><^>^<>^^<v<<^v>>^>^^>v><><^>v<^<^<<^^<>^v>>v^>v<^v>>v<<^^<vv<v^<>>>vv^^v>><^<v<>><<<v^v>^v^>v<<<v<<<^><v^>><<>^<^v>>v<>>vvv^<v^><^<>>v<>vv>>>>v<<v^>>v^^^>><^><vv<v<v>vv<>v^<^^^<^><^>>vv<<vvv>v^>>^>>^<^v>>>v^<v^^<>vv<v>v<^v<><vv<>^><<^>>>v^^<<>>^>><vv^v^^v>^>v^^<^<v^vv>^><<>><<^><^>v^<v^<^^^<<^>>>>^^v^<vv><<^>v^v>v>>
<^>v<v^<v^<>vv<v<^^<>^^vv><<>v^^>vvv>^v<^^^>v^^><v^<^^vv><v^>>v^>v<><vv>v<v>^^^>v<v><><>>v^^<^<><^<><^<^<>>v<^>^>^v<<v<<^v^^<v^v>^^vv^v<>v^^<^<v<>v><^v>>>^>^>>v^v><v<vvvv><v>><<>v<^v>^v>>v<<>>v^^<><^^<<>^>>vv<v<^^^<v<>^v<v^^<<^>v>v<><<^<<<v^^><^v<^v<^>^^^<^<<v^<>><^v^>^<v>v<>v^^>^v>^><>>>>vvv<>>^<<^v<v<<^<<>v<>v>^<<^v^<v<>v>v<<v^vv^v>v>^v<>>vv^vv<><><vv^<<>>^^><^<^vv^v>v<v^<^<v^v<vv<>v^<^>v<^><><<<^<>>^v^<><v^^<vv>><<v>^^>v>>vv><v><<>><>v>>><v^>>><^^v^^<>vvv^^>^><^<>^>^^vvv<vv^^^<>v<<>>>v>^<>>^v<^<<v>>^<v^v^v^^<vv<v<<^>>>>^>^^>>^^v>^>^^<^<<v<<<v^^<v<<<>^>v<^v^v>vv<v^v^><^^v>v<>^<v>v^<<^v<>>^^>v<^vvv><<^<<v^>v>><>>v>^^^^>v>v^^<>^^<^<v<<>v>><v>^v>>v^^vv<<>^^><v^v><><^>v<<^><v<>v<>v<v^<>^>^<<<^v<^><v<<<<<v>><<<<>>^^v><^<vvv^<<><^^><^>v<v<>^^v<^v>>><^<<v>^<>>v<<^^>^<<>>^v<^>^v^<<v^v<v<^^>^^^><<v<>>^^>><>v<^^>>><><v^^><>^^<^v>v<v<<^v<^v<^v><<v<vvv^>><<<vv<<^^<>^<>^<<<<<<^^^^>^>>>^>>^>^<v<<^vv><<^vv^><>>v><v<^v^>>vv>v><<><<><^<v>v<vvv><<><<v^v^v<^^<>v>vv^v>v>vv>>>vv><^<v>^^<vv^><<>>vv^<v^v^^
><<^vv<^<v><^v^vv>>>><v>^<<^>^>><vv<<>^v><<>^v<<vv^<<<><v<<>><^^<<><v<vv>^v^^^<^^^<<>v^><<v><<v^<^><^^vvv><v>^^>>^v<<>>^<>v^vv>^v^^<v>>^>>v^vv<>vv<>v^>^v^<v^<<<v^^>><>>v<vv^><<>vv<v<>v^<><vv>^^>>v^<<><<^>>><v<^>v^v>>v>^><^v^><>vv<v<<>^>>^v<<>>vv><vv^^^>^v>><>v^^^^v>v<vv^vv^<^<<<>>vv><>^<>v^><^<>^<<^v>^<^v^>^^v^vv^^^v>v>^v<>>v^v>>v>>>^>^<v<><<^<^>vv>>>^vv^^v<^v<<v<>^^<^v<<v>^>v>^>>^vvv>v>^v^<^<^<>v^<v<>^><><vv^^<<v^>^<vv>vv^v>^><>v<<><^<<<><vv><^<<v^v<v<<><>>>v^^^>^<<>>>><v>^<<^^>>v^^<>>^>^<v<<<^vv<v^^^<v<^v<<<^><^v>vv<>vv>>>^><>v>>>>^<^<^^v<>^>^v><><^>v>><<v>^><<v>vv^<^<>^v^vv^^>vv><^<^v^^<^^^^v>v^v<^^vv<v><><^^><><vv<^^>>v>^vvvv^<>v<v<^^^>^><v<v<<^v^v^^><<><>v^^^<>>>>^^v>v<<vv<^^vv<<><<vv<v^^^<><^v>vvv^^^^v>><<<v^<v^^v>>>^<v<^<^vv^><>v>v<v^^v<>>^>vvv<v<>^v^><v<<^^>^^>>vv<<vv>^v^>v>^<>^v>^><<^<v>><>v<vv<v<v^^v^<<<v<v^^v<<^vv^v<^v><^>^vvv>v<>>>^>v<<>>^>>>^>^v<^<<><^^v<<>v>v<<<>^v<><^<^>>>v><<>^^^>^v<^<>>v^><<<>^>vv><<^<<v^>vv^^>v<<v><v<>vv^^<<>^v<<^<v>^^><v>^v>^v^^<v<^^<<>^<<<<^>^>v<^><
>><>^>v<^vv<^^>v^^><<v^v>>><<^><<v>^^^<v^^v<v<<^^^v<v^vv<^^<<^>>><^>>>v><v><v><<^<v><><>v>><>>vv>><><<><<^v><<>v^vv^>v>^^^^^<v<><v^>v^><><^>vv<>v^<^>v><^v^vv>vv>>>^^^<<^>v>v^<v<v<v^>^<<^><<^<<>v^>v^^<><^>v>^><vv<v^<v^^>^v>^^^v^>^>v<v><^>><v^^<>>v>v^>^<<^>v<<<vvv^<^v<<v<^v<><v<<>v<vv>vv>^<^vv^>v^>>^>vvv^^^>^><<^^v<^^vv^>v<<v<^^<^v<^<<^vv^><<<>^v^^^>>vv><^^<vv^<<>v>^><>^^<>>>vv>^<<^<^v<v<^<<vv^^><<>^^<<>vvv^>^v<v>^>v<>vv>>v^^><v<v^v<<v<<<v^>^^>v>v^v<<<v>v<><<^>vv^v^^^^v^<^><>vv^v<><<v>>v<>^^><<^^^^<<^>v<>v<<><^^<v<<^>^^^^vv><<<v^<><<v^^>v<^>vv><>^vv<>>v>vvv^vv<^v^^v<v<^v^<^^^>>v<v^vv>v^<^v^v^^^<v>^<vv^^<<<v^>><^^<>v^v><<>>>^>^>><^v>>v^>^^>>^<><>><v<>^^v<vvv><>^^v<>^^>^<<<^>^^><<<<^vv>>^v<>^^v>v<^^<>v>^<^<v<>^^<>>vvvv<v<<<<>^v><<>v<^v<^<>vv^^<<^<^<v>^^^v<<v>v<<><v<><><v^^>><><><^v<v<vv<<vv>>^<^><<vv>>^v<><<^<vv^v><^^vv^>^<v^<^<<>>^<><<<>^<<v^v<<><v>>^v<v>^<<<>v<v<>v^<<^^>vv^>^v<^v^^vv<^<<<<>v<vv>v>^<<^><^vv^^^>><>vvv^<>^v<>v<^^<^^v>v^^^<^<^^><><><^^^<^^v>v><v<v^<v<^^<<v>^^v><v<<vvv<^>>vvv
<v>>^v<<^>vvvv<>v>^><v^<v<^><>^v>^^><><vv^v>^vvv^><<^<^v<^<<<v>>vv^^^^^<^>v<><>v><^>v>^v^>v<><>^v<>^>^^>v>^><v>^vv<^<^<^^<vv<<>>>^><<>^v^^^^vv^<<v<><<^vv<^v><<<^v<v>vvv><<><><^><vv^^v<<<>v^^^^>^>>^<<vv<^^<vv^v^^v^<vvv<^>><^v>^<^>v<v>>v>>vv><v>>^>>v^vv>>vv<>><^<<^^^^v<<v^v<>v<<^>v^^^vv^<>^<<vv<><v><^v^v<^^^<>><^v^>^<<v<v^<<<v>>vv<<vvvv^><<^^v<^v^<>^^^^v^^v<>>vvvv^v<v<v>>^>v>^^<v<<>^><<^<v^v>^>^^>^<vvv<vv>vv><v<^>^^^<><v>><^><^v>><vvv<v^<>>>^<<v^^<<<><><<^^><^^v><v<^^^v^^v<v^<vv>><^>v<<v<<>>v^>v^>>^>vvv<^>vvv<>^^>^vv<v<v><>v^<^v>^>vv^v^^vv^>^v^v>^<^>><^^<vvvv>>^^<^v>^^^>><vv<v>v<><><<^>^>^<><^<^<<v^<<^vv<<>^v>^>vvv<vv^vv<vvv>><^^v>>><v><>vv><><vv><v^^^<>v<>v^^v>>^<<><>^v><<<^<v<v><vvv>v<<<^>v>^>^>v>>v^<<<vvv>>v><^v>>^<^><^<<v^^<vv^><>>^<^v<^vv<>>>>^^v^<v>>^^<<>^><>>^v>^v^><v<>>^^>^><v<^<^^>>^<^<v^<>v<v^v>vv^>^^<<v<>^v^<<^<^<^^><v>v>><<vv>><<<><^v^^>^v>>vv^v><v>vvvv^vv^<<>v>>^v^^<v>>v>v<<^^<>v<^v^^vv^^^^v><v><^^>vv<>><v^vv>>^><><<vv<^<vv>><^v<v>>^v^^v^^v><>^v>v^<v<vv>vvvv<>v^^v^^><>vv^^>v
^<v><^>v>v><<vvv^^v^><>>^<v<>>>v^>^<v^>v>v^^>v<><^>><v^><^><><v^^v^v^^^v<<^vv^v^>vv<>v>v>vv<^><^<>><>v^>vv^v>vvv<v<^vv<>>vvvv>v<<>^><^vv><<v^^^<v>vvvv^vv^<<<<><vv>^>>vv>><^<<^>><<><<v^^><^^^v^vv>v<^v>^vv>^v<<v^v<><>>vv^^>^v^>^v<>>v>^>><>^vvvv><>^<^>>>^>>^^v<vv^>>><><<>^<^^>vvvv>v>>v^^vvv><vv^v^><^^^>v<v^^<vv<v<><^<><v>^>^<><v<v^^><<^v><^><v>>v<<^vvv^vv>>^<v<^v>v<v>>v<vv<^><>^>>>^^<v<>^<^v<><><vvv<v<^<>v><<>^>><<^<^>v^>>v>^v>v^<>v<<>^v<^>v<<^^>>>v<^<^^^v<>v^<v^^^<<v<v<<^<<>v<^<<^^v><v^>v^><v^^<>^v><vv<<v<<>vv^^><v^>>v><^>v<>v<^vvv<>v>v^vvvv<<v<^^<<^<^vv<<^v^<><>v<>>^<<><>^^<v<><v><v<^>><v>v<<v^vv^^<<^^><<vv><v^^<><^<><^vv<^>^<<><^^^<<^v^<>^>v>vv>v<<v<<^>>^v^<^^v<^><vv<><<^<v<<<^^>^^vvv>v<>^^vvv<v><v<v<>>>>>vv<^><^><^<vv><>v<v>>v<<<><vvvv><^>>^<><^<>v<v>^<<^v<>v^^<><^<<^^<>>>>^<<v><>v<^<>v>>>><vv<<^v^v<v^v^><^<>^^<><v><v<<><<v<<<v<v><v^v><<v^<v>v^v>v^<<<^^<<vvv<^>v>v<>^><<>>^>>^v<v<^^^v<v<vv><^v<<v<>^><v^^<>v><v>>>v<>>><^>^>^^<^vvv<>>vv<^<^>><><v^^^v^v<^>^<<^<><v>>v<vv^^>>^v<vv<<<^^v^^>>
v>v>>>v>v>v>vv^<v>^^v^<vv<>^^<>>>v^v>><^^vv<v<^<^vv^<<<^<^<^^<v^<<<v<v<<v><<<vv^>><>>^<v<^>>v>v^v^vv^^^^>^^^^v<v>^><>^<^>^>>v<>v^<<v<vv>>v<<v^<<v>>><^>>vv>><>>^v><<<>>>^v^<>v>^v^^^^<<v^>^^^v<^v>^vvvv<>vv<>v>^v^v>v<>v<v^v<<><^v^^>>v>v<vv<v^>>vv<^v<><vvv>^>><<<^>><<>^<v^vv<vv^<^vv^v^<v<^><>^<vv>>><<><<^v^^<^^<>^vv>><^<<^>^<vv><^>>^vv<>vvv>^<v><>^<v>^v<<<^^<v><<^>^^v^<><<>^v>^^>^^<>v^>^<v>v<>>>vv<^v>vvv<>v>>>v<<>v^<<<>^>v<>v>><v<>^^^<<>^v>vv>^><^<v<vv<><v>><>><^>^>^v^v^>vvv>v>^^^^v^v^<<^>^><<v^v>^v^v<>v><^<^>>^^>^v^^>>^><^^>^><<>^><^v<><vv<>v><>^v<^vv<v>^<v<<^v<<>^vvvv>v><<^vv^<>><^>>>v^v^v>v^^v>v<><vv^<>><<<^^v^v>>vv>>v^v<^>v^v>>v<>^^v<^<^^vvv>>><>v>>^v^>^vvv>^^^^^vv>v^>v>v><<^^^<>^vv<vv<>vv^<vv^>><<v<<<<v^>v>^^<<>>>^v><v>^^^<^>>v^^^^v<^^<<v<>>vv<v>^<^^<>><v^^<<>^<^><<<>^<^>v>>^><^<>v>v<>v^v>vvv<<<^>^>^^v^><<^vv<v>>>^^<<><>^vv>^<<^^v><vv>^v<>v^<v<^<<>v^<v<>v^<^v>><<<>^<^>^>v<v<v^>v^v^>v^v^v<vvvvvv<v>v>^<>>^<>^^^<>^>>^<^>>^>>^vv><<>>^v>^^><vv>><><v^^v<^^>>>>v^v>>^><^<v<v<^>^^v<<vv^^v<<^>>
<>^<v^^<><<><^^vv><v<^<^v>>>vvv<v<><<v<>>>^><<^^>^<<<<v<v^^>^^>v^vvv><vv<><><^^>>v<<v<v<v>><<^><^v<>vv^v<vv^^<vvv<^<<<^>v^>^<>^>v^><<<<<^<v^>v^vv<<><>v^^^^^v<vv>>v<^vv>v<>^<v<^v<>^>^>>><^<v^<vv^>^<<^<v<v>v<<^^<^<>^<<<^^<v^v^v^>>v><v>v^>v><vv^>vvv>v>^vv>>v<>^<<>^<v>^>^^<^v<v>^^><>^>>^<^>>vv>^>>v^^<v><^^><^^<^v>>>^<<v<<^<><vv>>v^^>^<><<><>><>><^>v>v^<^v><v<<<v>v<<<<<v<><v>>v>>^^>^^^<vv^^vv><^v>^>v<<<<>vv^>>><^^>>><vv<^^<<<>^<^>v>>^v><>^><v<>>v<v>v<^vv<v<<>^<v^v>^>vvv<^^<>>>^>>v^^v<>v>^>v^<vv>>><v<<<v><v^<v><v^>>>v>^<^>v<>^^>>>>v>vv^^v<v^>vv<v<^vv<><^<>^v^<vv<><>^<><>v>^>>v<^<><^^<>v>vv>^v<v^<<v>vv<>><<<>>^>>vv^^^<^<v<>vv<>>>>>v^><vvv>^>vv>v<^><<>^>v<<v><^>v>^^<^<^v^^v<>>><<^^<>^><v^>v^><>^v<><^>^v<<^^><^^^<>>>^v<<>>v^<>v<>><vvv>v><v>^<^^<^<vv^<>^v<<v^v^^^<>>^vv<^<vv<^v<<>^<<^vv>vv>^<<<>>v^>v<v>>v<>vv^><v<^<^vvvv>>^><><v<^>v>>>v<^^v<>^<><v>v<<<>>>v^>^v<<v><^^^<>^v>^v>^^>vv<<^^v^<v^<<>>v<v^v>^<<<>>v<<^^<<^v><><v>>v<v^<v^v<^^vv<^^<^^<>>^^<v<^v><^^v<<^v<^><>^<<>^^v^>><^<^vvv>v^>^<v><^^>v><v^
"""


class TestDay:
    @pytest.mark.parametrize(testit.names, [(example, 10092), (answer, 1568399)], ids=testit.ids())
    def test_first_star(self, indata: str, want: int) -> None:
        testit.first_star(day15.Day(indata), want)

    @pytest.mark.parametrize(testit.names, [(example, 9021), (answer, 1575877)], ids=testit.ids())
    def test_second_star(self, indata: str, want: int) -> None:
        testit.second_star(day15.Day(indata), want)
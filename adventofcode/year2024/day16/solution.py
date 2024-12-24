import collections

# def second_star(maze: list[list[str]], row: int, col: int) -> int:
#     return _travel(maze, row, col)
import heapq

#
# def _travel(maze: list[list[str]], row: int, col: int) -> int:
#     deq = collections.deque()
#     deq.append(((row, col), '>', {(row, col)}, 0))
#     min_score = 1000000000
#     rows = len(maze)
#     cols = len(maze[0])
#
#     while deq:
#         point, curr, visited, score = deq.pop()
#
#         if score > min_score:
#             continue
#
#         if maze[point[0]][point[1]] == 'E':
#             min_score = min(min_score, score)
#             continue
#
#         for direct, coords in _directions.items():
#             if any((
#                 curr == '>' and direct == '<',
#                 curr == '<' and direct == '>',
#                 curr == '^' and direct == 'v',
#                 curr == 'v' and direct == '^',
#             )):
#                 continue
#
#             new_row = point[0] + coords[0]
#             new_col = point[1] + coords[1]
#
#             if 0 <= new_row < rows and 0 <= new_col < cols:
#                 diff = 1 + (0 if curr == direct else 1000)
#
#                 if maze[new_row][new_col] == '#' or (new_row, new_col) in visited:
#                     continue
#
#                 deq.append(((new_row, new_col), direct, visited | {(new_row, new_col)}, score + diff))
#
#     return min_score


# def first_star(maze: list[list[str]], row: int, col: int) -> int:
#     return _travel(maze, row, col)


_directions = {
    '^': (-1, 0),
    'v': (1, 0),
    '<': (0, -1),
    '>': (0, 1),
}


def _travel(maze: list[list[str]], start_row: int, start_col: int) -> int:
    import heapq

    # Dimensions of the maze
    rows, cols = len(maze), len(maze[0])

    # Priority queue: stores (score, row, col, direction)
    pq = []
    heapq.heappush(pq, (0, start_row, start_col, '>'))  # None for no initial direction

    # Visited dictionary with the format {(row, col, direction): score}
    visited = {}

    # Directions with their respective deltas
    _directions = {
        '>': (0, 1),  # Right
        '<': (0, -1),  # Left
        '^': (-1, 0),  # Up
        'v': (1, 0),  # Down
    }

    while pq:
        score, row, col, curr_dir = heapq.heappop(pq)

        # Avoid revisiting already processed states with equal or better score
        if (row, col, curr_dir) in visited and visited[(row, col, curr_dir)] <= score:
            continue

        # Mark the current cell and direction as visited
        visited[(row, col, curr_dir)] = score

        # Check if we reached the end cell
        if maze[row][col] == 'E':
            return score
            # return score

        # Iterate over all possible directions
        for direction, (dr, dc) in _directions.items():
            # Compute the new cell's coordinates
            new_row, new_col = row + dr, col + dc

            # Check boundaries and obstacles
            if 0 <= new_row < rows and 0 <= new_col < cols and maze[new_row][new_col] != '#':
                # Calculate the cost for this move
                move_cost = 0 if curr_dir == direction or curr_dir is None else 1000
                new_score = score + 1 + move_cost
                # print(heapq.nsmallest(2, pq))
                # Add the new state to the priority queue
                heapq.heappush(pq, (new_score, new_row, new_col, direction))

    # Return -1 if no path to 'E' is found
    return -1


def _travel2(maze: list[list[str]], start_row: int, start_col: int) -> tuple[int, int]:
    import heapq
    from collections import defaultdict

    # Dimensions of the maze
    rows, cols = len(maze), len(maze[0])

    # Priority queue: stores (score, row, col, direction)
    pq = []
    heapq.heappush(pq, (0, start_row, start_col, '>'))  # None for no initial direction

    # Visited dictionary to store minimal score and path count for each state
    visited = defaultdict(lambda: (float('inf'), 0))  # (score, path_count)

    # Directions with their respective deltas
    _directions = {
        '>': (0, 1),  # Right
        '<': (0, -1),  # Left
        '^': (-1, 0),  # Up
        'v': (1, 0),  # Down
    }

    # Initialize the starting position
    visited[(start_row, start_col, '>')] = (0, 1)

    while pq:
        score, row, col, curr_dir = heapq.heappop(pq)

        # Skip if this state is no longer optimal
        if score > visited[(row, col, curr_dir)][0]:
            continue

        # Check if we reached the end cell
        if maze[row][col] == 'E':
            # Sum up all ways to reach the end cell with the minimal score
            for key, (sc, count) in visited.items():
                if sc == score:
                    print(key, count)

            return score, 1

        # Iterate over all possible directions
        for direction, (dr, dc) in _directions.items():
            # Compute the new cell's coordinates
            new_row, new_col = row + dr, col + dc

            # Check boundaries and obstacles
            if 0 <= new_row < rows and 0 <= new_col < cols and maze[new_row][new_col] != '#':
                # Calculate the cost for this move
                move_cost = 0 if curr_dir == direction or curr_dir is None else 1000
                new_score = score + 1 + move_cost

                # If this path improves or matches the minimal score for the state
                if new_score < visited[(new_row, new_col, direction)][0]:
                    visited[(new_row, new_col, direction)] = (
                        new_score,
                        visited[(row, col, curr_dir)][1],
                    )
                    heapq.heappush(pq, (new_score, new_row, new_col, direction))
                elif new_score == visited[(new_row, new_col, direction)][0]:
                    # Accumulate path counts for states with equal scores
                    visited[(new_row, new_col, direction)] = (
                        new_score,
                        visited[(new_row, new_col, direction)][1] + visited[(row, col, curr_dir)][1],
                    )

    # Return -1 if no path to 'E' is found
    return -1, 0


def _travel3(maze: list[list[str]], start_row: int, start_col: int) -> tuple[int, int, int]:
    import heapq
    from collections import defaultdict

    # Dimensions of the maze
    rows, cols = len(maze), len(maze[0])

    # Priority queue: stores (score, row, col, direction)
    pq = []
    heapq.heappush(pq, (0, start_row, start_col, '>'))  # None for no initial direction

    # Visited dictionary to store minimal score, path count, and point count for each state
    visited = defaultdict(lambda: (float('inf'), 0, 0))  # (min_score, path_count, point_count)

    # Directions with their respective deltas
    _directions = {
        '>': (0, 1),  # Right
        '<': (0, -1),  # Left
        '^': (-1, 0),  # Up
        'v': (1, 0),  # Down
    }

    # Initialize the starting position
    visited[(start_row, start_col, '>')] = (0, 1, 1)  # Starting state

    while pq:
        score, row, col, curr_dir = heapq.heappop(pq)

        # Skip if this state is no longer optimal
        if score > visited[(row, col, curr_dir)][0]:
            continue

        # Check if we reached the end cell
        if maze[row][col] == 'E':
            # Sum up all ways to reach the end cell with the minimal score
            total_paths = sum(
                count for (r, c, d), (s, count, _) in visited.items() if (r, c) == (row, col) and s == score
            )
            total_points = sum(
                points for (r, c, d), (s, _, points) in visited.items() if (r, c) == (row, col) and s == score
            )
            return score, total_paths, total_points

        # Iterate over all possible directions
        for direction, (dr, dc) in _directions.items():
            # Compute the new cell's coordinates
            new_row, new_col = row + dr, col + dc

            # Check boundaries and obstacles
            if 0 <= new_row < rows and 0 <= new_col < cols and maze[new_row][new_col] != '#':
                # Calculate the cost for this move
                move_cost = 0 if curr_dir == direction or curr_dir is None else 1000
                new_score = score + 1 + move_cost

                # If this path improves or matches the minimal score for the state
                current_paths, current_points = visited[(row, col, curr_dir)][1:]
                if new_score < visited[(new_row, new_col, direction)][0]:
                    visited[(new_row, new_col, direction)] = (
                        new_score,
                        current_paths,
                        current_points + 1,
                    )
                    heapq.heappush(pq, (new_score, new_row, new_col, direction))
                elif new_score == visited[(new_row, new_col, direction)][0]:
                    # Accumulate path counts and point counts for states with equal scores
                    visited[(new_row, new_col, direction)] = (
                        new_score,
                        visited[(new_row, new_col, direction)][1] + current_paths,
                        visited[(new_row, new_col, direction)][2] + current_points + 1,
                    )

    # Return -1 if no path to 'E' is found
    return -1, 0, 0


import heapq
from collections import defaultdict

# Directions for movement
_directions = {
    '>': (0, 1),  # Right
    '<': (0, -1),  # Left
    '^': (-1, 0),  # Up
    'v': (1, 0),  # Down
}


def _travel4(maze: list[list[str]], start_row: int, start_col: int) -> tuple[int, int, int]:
    rows, cols = len(maze), len(maze[0])

    # Priority queue for state (score, row, col, direction)
    pq = []
    heapq.heappush(pq, (0, start_row, start_col, None))  # Start with no initial direction

    # Dictionary to store minimal score, path count, and visited cells for each state (row, col, direction)
    visited = defaultdict(lambda: (float('inf'), 0, set()))  # (min_score, path_count, visited_cells)

    # Initialize the starting point
    visited[(start_row, start_col, None)] = (0, 1, {(start_row, start_col)})

    while pq:
        score, row, col, curr_dir = heapq.heappop(pq)

        # Skip if this state is no longer optimal
        if score > visited[(row, col, curr_dir)][0]:
            continue

        # Check if we reached the end
        if maze[row][col] == 'E':
            # Sum up all paths with minimal score and count distinct cells
            total_paths = sum(
                count for (r, c, d), (s, count, _) in visited.items() if (r, c) == (row, col) and s == score
            )
            # Count the unique cells covered in the paths with minimal score
            all_visited_cells = set()
            for (r, c, d), (s, _, cells) in visited.items():
                if (r, c) == (row, col) and s == score:
                    all_visited_cells.update(cells)
            return score, total_paths, len(all_visited_cells)

        # Explore neighboring cells in all directions
        for direction, (dr, dc) in _directions.items():
            # Compute the new cell coordinates
            new_row, new_col = row + dr, col + dc

            # Check if the new cell is within bounds and not a wall
            if 0 <= new_row < rows and 0 <= new_col < cols and maze[new_row][new_col] != '#':
                # Calculate move cost (0 if same direction, 1000 if different direction)
                move_cost = 0 if curr_dir == direction or curr_dir is None else 1000
                new_score = score + 1 + move_cost

                # Update visited state if this path is better or equal
                new_visited_cells = visited[(row, col, curr_dir)][2].copy()
                new_visited_cells.add((new_row, new_col))

                # If this path improves or matches the minimal score for this state
                if new_score < visited[(new_row, new_col, direction)][0]:
                    visited[(new_row, new_col, direction)] = (
                        new_score,
                        visited[(row, col, curr_dir)][1],
                        new_visited_cells,
                    )
                    heapq.heappush(pq, (new_score, new_row, new_col, direction))
                elif new_score == visited[(new_row, new_col, direction)][0]:
                    # Accumulate path counts and merge visited cells
                    visited[(new_row, new_col, direction)] = (
                        new_score,
                        visited[(new_row, new_col, direction)][1] + visited[(row, col, curr_dir)][1],
                        visited[(new_row, new_col, direction)][2].union(new_visited_cells),
                    )

    # Return -1 if no path to 'E' is found
    return -1, 0, 0


def first_star(maze: list[list[str]], row: int, col: int) -> int:
    return _travel(maze, row, col)


def second_star(maze: list[list[str]], row: int, col: int) -> int:
    sdda = _travel4(maze, row, col)
    print(sdda)

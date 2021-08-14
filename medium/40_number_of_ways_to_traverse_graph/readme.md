# Number Of Ways To Traverse Graph

You're given two positive integers representing the width and height of a grid-shaped, rectangular graph. Write a function that returns the number of ways to reach the bottom right corner of the graph when starting at the top left corner. Each move you take must either go down or right. In other words, you can never move up or left in the graph.

For example, given the graph illustrated below, with width = 2 and height = 3, there are three ways to reach the bottom right corner when starting at the top left corner:

```
 _ _
|_|_|
|_|_|
|_|_|
```

```
Down, Down, Right
Right, Down, Down
Down, Right, Down
```

Note: you may assume that width * height >= 2. In other words, the graph will never be a 1x1 grid.

## Sample Input

```
width = 4
height = 3
```

## Sample Output

```
10
```

### Hints

Hint 1
> Think recursively. How many positions in the graph can access the bottom right corner of the graph? In other words, what positions do you need to reach before you can reach the bottom right corner?

Hint 2
> The number of ways to reach any position in the graph is equal to the number of ways to reach the position directly above it plus the number of ways to reach the position directly to its left. This is because you can only travel down and right.

Hint 3
> Using the information in Hints #1 and #2, can you come up with an efficient way to solve this problem that doesn't repeatedly perform the same work? What does a dynamic-programming implementation look like?

Hint 4
> To efficiency solve this problem, simply loop through the entire graph, column by column, row by row, and calculate the number of ways to reach each position. If you're on the top or left edge of the graph, there's only one way to reach your position. If you're anywhere else in the graph, the number of ways to reach your position is the number of ways to reach the position directly above it plus the number of ways to reach the position directly to its left (which you've already calculated and should be storing). Every time you calculate the number of ways to reach a position, store the answer so that you can use it later in the calculation of other positions.

```
Optimal Space & Time Complexity
O(n + m) time | O(1) space - where n is the width of the graph and m is the height
```

![solution](answer.png)

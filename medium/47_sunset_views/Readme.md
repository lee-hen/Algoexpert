# Sunset Views

Given an array of buildings and a direction that all of the buildings face, return an array of the indices of the buildings that can see the sunset.

A building can see the sunset if it's strictly taller than all of the buildings that come after it in the direction that it faces.

The input array named buildings contains positive, non-zero integers representing the heights of the buildings. A building at index i thus has a height denoted by buildings[i]. All of the buildings face the same direction, and this direction is either east or west, denoted by the input string named direction, which will always be equal to either "EAST" or "WEST". In relation to the input array, you can interpret these directions as right for east and left for west.

Important note: the indices in the ouput array should be sorted in ascending order.

## Sample Input #1

```
buildings = [3, 5, 4, 4, 3, 1, 3, 2]
direction = "EAST"
```

## Sample Output #1

```
[1, 3, 6, 7]

// Below is a visual representation of the sample input.
//    _
//   | |_ _
//  _| | | |_   _
// | | | | | | | |_
// | | | | | |_| | |
// |_|_|_|_|_|_|_|_|
```

## Sample Input #2

```
buildings = [3, 5, 4, 4, 3, 1, 3, 2]
direction = "WEST"
```
## Sample Output #2

```
[0, 1]
// The buildings are the same as in the first sample
// input, but their direction is reversed.
```

### Hints

Hint 1
> Is there a way to solve this problem in one loop?

Hint 2
> How does your solution change based on the direction that the buildings are facing? You can use the same approach for each direction by simply changing the direction in which you traverse the array of buildings.

Hint 3
> There are multiple ways to solve this problem, but one is to maintain a running maximum of building heights. Loop in the opposite direction that the buildings are facing, and keep track of the maximum building height that you've seen. At each iteration, compare the height of the current building to the running maximum; if the current building is taller, then it can see the sunset; otherwise, it can't. Finally, at each iteration, update the running maximum.

Hint 4
> Another way to solve this problem is to use a stack. Loop in the direction that the buildings are facing, and add the index of the current building to the stack at the end of each iteration. Before adding elements to the stack, compare the current building height to buildings at the top of the stack. Pop off the top of the stack until the current building height is shorter than the height of the building at the top of the stack. This will remove all buildings that are blocked from seeing the sunset by the current building. At the end of the algorithm, the stack will only contain elements that can see the sunset.

```
Optimal Space & Time Complexity
O(n) time | O(n) space - where n is the length of the input array
```

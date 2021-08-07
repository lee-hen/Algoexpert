# Summary

This question is equivalent to the Spiral Traverse question on AlgoExpert; we recommend watching the video explanation of that question to better understand this question's solution.

The main differences here are that:

   - 1) We're dealing only with square-shaped matrices, which simplifies the problem; we don't need to handle edge cases where we have a lone row or column in the middle of the matrix.
   - 2) We're overwriting values in the input matrix, which means that we need to store a reference to one value in each spiral or ring that we spin--the top right value in our solutions--so as not to lose it forever in the process of spinning values.
   
Otherwise, once again, this problem is equivalent to Spiral Traverse and can be implemented either recursively or iteratively; the iterative version is naturally better from a space-complexity point of view.

## Complexity Analysis

This question is self-evidently solved in linear time (with respect to the total number of value in the input matrix) and with constant space (for the iterative solution).

## Closing Thoughts

As mentioned a couple of times above, this question is a variant of the Spiral Traverse question on AlgoExpert. It's important to walk through an example with a large-enough matrix (a 5x5 matrix, for example) and to avoid off-by-one errors as we spin each ring with our for loops.

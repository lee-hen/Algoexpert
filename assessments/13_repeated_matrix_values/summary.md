# Summary

This question can be solved optimally and cleanly with two passes through the entire input matrix, as well with an additional preliminary step.

## Preliminary Step

First, we have to realize that, as obvious as it may seem, for an integer found in the matrix to be one of the integers that we have to return, it has to be found in every row and every column of the matrix. After all, this is the type of integer that the problem asks us to return.

This means that the integers in the first row (or in any row for that matter) of the matrix are the only integers that we might return, since any integer that's not in the first row of the matrix is obviously not in every row of the matrix.

Thus, the preliminary step of our solution is to store every integer in the first row of the matrix in a hash table (we pick the first row for simplicity, but we could pick any row); this hash table will keep track of the counts of all the potential integers to return. We initialize these counts to 0.

Note that we can perform a slight space-complexity optimization here by picking the integers in the first column instead of those in the first row if the height of the matrix is smaller than its width; this will minimize the size of our hash table.

## Passes Through Matrix

The goal behind our algorithm is to iterate through every value in the matrix row by row and then to reiterate through every value in the matrix column by column, counting every unique integer found in each row and column, in the hopes of having some integers with counts of matrixWidth + matrixHeight after both passes through the matrix.

During each pass, at each integer that we visit, we first check if the integer is stored in our hash table. If it isn't, that means that it wasn't found in our first row or column, so we can skip it. Not storing counts for these integers saves us space. If the integer is in our hash table, we check that it has the correct "count so far." In other words, we check that the integer has been found at least once in each previous row and / or column. During our first pass, this correct "count so far" will simply be the index of the row or column that we're at. If the integer in question doesn't have the correct count, this means that we're either dealing with an integer that was missing in a previous row and / or column, or that we're dealing with an integer that was already visited in the given row or column; in both cases, we skip the integer. Otherwise, we're dealing with a potentially "valid" integer, and we increment its count by 1.

Once we're done with both passes through the matrix, we return all of the integers in the hash table whose counts are equal to matrixWidth + matrixHeight.

## Complexity Analysis

This question is self-evidently solved in linear time with respect to the total number of elements in the matrix, since we only do two passes through the entire matrix (plus the preliminary pass through a single row or column). Our solution's space complexity is linear with respect to the size of the row or column that we select in the preliminary step.

## Closing Thoughts

This is a fairly straightforward matrix problem that simply requires a few edge cases to be handled, like when potentially "valid" integers appear twice in a given row or column.

This is also a good example of a problem that is much more cleanly and easily solved with two separate passes through the input matrix. You might try to solve this problem with a single pass through the matrix, but you'll likely complicate your code and logic for no space-time-complexity gain.

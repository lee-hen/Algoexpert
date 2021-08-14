# Class Photos

It's photo day at the local school, and you're the photographer assigned to take class photos. The class that you'll be photographing has an even number of students, and all these students are wearing red or blue shirts. In fact, exactly half of the class is wearing red shirts, and the other half is wearing blue shirts. You're responsible for arranging the students in two rows before taking the photo. Each row should contain the same number of the students and should adhere to the following guidelines:

   - All students wearing red shirts must be in the same row.
   - All students wearing blue shirts must be in the same row.
   - Each student in the back row must be strictly taller than the student directly in front of them in the front row.

You're given two input arrays: one containing the heights of all the students with red shirts and another one containing the heights of all the students with blue shirts. These arrays will always have the same length, and each height will be a positive integer. Write a function that returns whether or not a class photo that follows the stated guidelines can be taken.

Note: you can assume that each class has at least 2 students.

## Sample Input

```
redShirtHeights = [5, 8, 1, 3, 4]
blueShirtHeights = [6, 9, 2, 4, 5]
```

## Sample Output

```
true // Place all students with blue shirts in the back row.
```

### Hints

```
Hint 1
Start by determining which row will have the students wearing blue shirts and which row will have the students wearing red shirts. Once you know this, how can you determine if it's possible to take the photo?
```

```
Hint 2
The shirt color of the tallest student will determine which students need to be placed in the back row. The tallest student can't be placed in the front row because there's no student taller than them who can be placed behind them.
```

```
Hint 3
Once you know which students should be placed in each row, you can simply check if each student in the back row can be paired with a student in the front row who is shorter than them. If you can't find a satisfactory pairing for every student in the back row, then you can't take the photo.
```

```
Hint 4
Sort each input array in descending order, then determine which students will be in the front and back rows following Hint #2. After this, simply loop through your sorted input arrays, and check if the current tallest student in the back row is taller than the current tallest student in the front row. If you find that the current tallest student (one that has yet to be placed) in the back row isn't taller than the current tallest student in the front row, then the photo can't be taken.
```

```
Optimal Space & Time Complexity
O(nlog(n)) time | O(1) space - where n is the number of students
```

## Common matrices operations

1. **Addition**: Add corresponding elements of two matrices.
2. **Subtraction**: Subtract corresponding elements of two matrices.
3. **Multiplication**: Multiply two matrices by the dot product of rows and columns.
4. **Scalar Multiplication**: Multiply each element of a matrix by a scalar value.
5. **Transpose**: Swap the matrix's rows with its columns.
6. **Determinant Calculation** (for square matrices): Calculate the determinant of a matrix.
7. **Inverse** (for square matrices): Find the matrix that, when multiplied by the original matrix, results in the identity matrix.

## Common interesting 2D problems

Here are some examples of 2D array problems that are commonly encountered:  
1. **Matrix Rotation**: Rotate a given NxN matrix by 90 degrees clockwise or counterclockwise.  
2. **Flood Fill Algorithm**: Given a 2D screen, a point in the screen and a new color, fill the surrounding area until the color changes from the original color.  
3. **Island Counting** (Number of Islands): Given a 2D grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.  
4. **Sudoku Solver**: Given a partially filled 9x9 2D array, write a program to fill the grid using Sudoku rules.  
5. **Knight's Tour Problem**: On a chessboard, find a path for the knight to visit every square once.  
6. **Maximum Sum Rectangle** in a 2D Matrix: Given a 2D array, find the rectangle (sub-matrix) with the maximum sum of its elements.  
7. **Spiral Matrix**: Given a matrix, return all elements of the matrix in spiral order.  
8. **Set Matrix Zeroes**: Given an m x n matrix, if an element is 0, set its entire row and column to 0.  
9. **Word Search**: Given a 2D board and a word, find if the word exists in the grid. The word can be constructed from letters of sequentially adjacent cells, where "adjacent" cells are horizontally or vertically neighboring.  
10. **Longest Increasing Path in a Matrix**: Given an m x n integers matrix, return the length of the longest increasing path in the matrix. From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary.  

These problems cover a range of topics including graph theory, dynamic programming, backtracking, and search algorithms, providing a good mix of challenges for practicing 2D array manipulations and algorithmic strategies.


Here are some additional common problems involving 2D arrays:

1. **Diagonal Traverse**: Given a matrix, return all elements of the matrix in a diagonal order starting from the top left corner.

2. **Pacific Atlantic Water Flow**: Given an `m x n` matrix of non-negative integers representing the height of each unit cell in a continent, the "Pacific ocean" touches the left and top edges of the matrix and the "Atlantic ocean" touches the right and bottom edges. Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.

3. **Minimum Path Sum**: Given a `m x n` grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path. You can only move either down or right at any point in time.

4. **Rotate Image**: You are given an `n x n` 2D matrix representing an image, rotate the image by 90 degrees (clockwise). You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

5. **Zero Matrix**: Write an algorithm such that if an element in an `MxN` matrix is 0, its entire row and column are set to 0.

6. **Magic Squares In Grid**: A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum. Given an `grid` of integers, how many 3 x 3 "magic square" subgrids are there? (Each subgrid is contiguous).

7. **Search a 2D Matrix II**: Write an efficient algorithm that searches for a value in an `m x n` matrix. This matrix has the following properties: Integers in each row are sorted in ascending from left to right. Integers in each column are sorted in ascending from top to bottom.

8. **Kth Smallest Element in a Sorted Matrix**: Given an `n x n` matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

9. **Cherry Pickup**: Given an `n x n` grid representing a field of cherries, each cell is one of three possible integers. The goal is to collect the maximum number of cherries possible by following the rules of moving from the top-left corner to the bottom-right corner and back to the top-left corner.

10. **Game of Life**: According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970." Given a board with `m` by `n` cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules. Write a function to compute the next state (after one update) of the board given its current state.
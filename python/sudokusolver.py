import sudoku

kodus = [8,0,0, 0,0,0, 0,0,0,
          0,0,3, 6,0,0, 0,0,0,
          0,7,0, 0,9,0, 2,0,0,
          0,5,0, 0,0,7, 0,0,0,
          0,0,0, 0,4,5, 7,0,0,
          0,0,0, 1,0,0, 0,3,0,
          0,0,1, 0,0,0 ,0,6,8,
          0,0,8, 5,0,0, 0,1,0,
          0,9,0, 0,0,0, 4,0,0]
try:
	solver=sudoku.Sudoku(kodus)
	print solver
	solver.solve()
	print solver
except sudoku.SudokuError:
	print "This is no sudoku!"

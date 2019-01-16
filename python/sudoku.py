#!/usr/bin/python

class Sudoku(object):
	sudoku = []
	rows = {}
	columns = {}
	cells = {}
	count=0
	candidates = []
	def __init__(self,sudoku):
		self.sudoku=sudoku
		for x in xrange(9):
			self.columns[x]=0
			for y in xrange(9):
				if self.sudoku[x+y*9] < 0 or self.sudoku[x+y*9] > 9:
					raise SudokuError
				if (self.columns[x] & 1<<self.sudoku[x+9*y])>0:
					raise SudokuError
				if self.sudoku[x+y*9]>0:
					self.columns[x]|=1<<self.sudoku[x+9*y]
				else:
					self.candidates.append(x+y*9)
		for y in xrange(9):
			self.rows[y]=0
			for x in xrange(9):
				if (self.rows[y] & 1<<self.sudoku[x+9*y])>0:
					raise SudokuError
				if self.sudoku[x+y*9]>0:
					self.rows[y]|=1<<self.sudoku[x+9*y]
		for cell in xrange(9):
			self.cells[cell]=0
			for xx in xrange (3):
				x=xx+int((cell%3)*3)
				for yy in xrange (3):
					y=yy+int(cell/3)*3
					if (self.cells[cell] & 1<<self.sudoku[x+9*y])>0:
						raise SudokuError
					if self.sudoku[x+y*9]>0:
						self.cells[cell]|=1<<self.sudoku[x+9*y]

	def __repr__(self):
		answer=""
		for y in xrange(9):
			for x in xrange (9):
				answer+=str(self.sudoku[y*9+x])+","
			answer+="\n"
		answer+="tries "+str(self.count)+"\n"
		return answer
	def get_cell(self,x,y):
		return 3*int(y/3)+int(x/3)
	def candidate(self):
		mp=10
		pos=81
	 	for i in self.candidates:
			if self.sudoku[i]>0:
				continue
			else:
				possible=0
				x=i%9
				y=int(i/9)
				cell=self.get_cell(x,y)
				for j in xrange(9):
					if ((self.cells[cell] | self.columns[x] | self.rows[y]) & 1<<j==0):
						possible+=1
				if possible==1:
					return i
				if possible==0:
					return 82
				if possible<mp:
					pos=i
					mp=possible
		return pos
	def available(self,n):
		answer=[]
		x=n%9
		y=int(n/9)
		cell=self.get_cell(x,y)
		for i in xrange(9):
			if ((self.cells[cell] | self.columns[x] | self.rows[y]) & 1<<i+1==0):
				answer.append(i+1)
		return answer

	def solve(self):
		trynext=self.candidate()
		if trynext == 82:
			return 0
		if trynext == 81:
			return 1
		for i in self.available(trynext):
			self.count+=1
			self.sudoku[trynext]=i
			x=trynext%9
			y=int(trynext/9)
			cell=self.get_cell(x,y)

			self.columns[x]|=1<<i
			self.rows[y]|=1<<i
			self.cells[cell]|=1<<i
			self.candidates.remove(trynext)
			if self.solve()==1:
				return 1
			self.cells[cell]^=1<<i
			self.rows[y]^=1<<i
			self.columns[x]^=1<<i
			self.sudoku[trynext]=0
			self.candidates.append(trynext)
		return 0


class SudokuError(Exception):
	pass

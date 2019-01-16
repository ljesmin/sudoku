import os
import tornado.ioloop
import tornado.web
from sudoku import Sudoku, SudokuError

class MainHandler(tornado.web.RequestHandler):
    def get(self):
    	sudokulahendus = ""
	sudokutabel=[]
	for i in range (81):
		sudokutabel.append(0)
        self.render("sudoku.html",sudokulahendus=sudokulahendus,sudokutabel=sudokutabel)
    def post(self):
    	sudokulahendus = ""
	sudokutabel = []
	teine = []

	for i in range (81):
		numbers=self.get_arguments("cell-"+str(i))
		if (len(numbers) == 1):
			sudokutabel.append(int(numbers[0]))
			teine.append(int(numbers[0]))
		else:
			sudokutabel.append(0)
			teine.append(0)
	try:
	        katsetus=Sudoku(sudokutabel)
		katsetus.solve()
		sudokulahendus=str(katsetus)
	except SudokuError:
	        sudokulahendus="See ei ole sudoku!"

        self.render("sudoku.html",sudokulahendus=sudokulahendus,sudokutabel=teine)

handlers = [
    (r"/", MainHandler),
    (r"/css/(.*)",tornado.web.StaticFileHandler, {"path": "./css"},),
    ]

settings = dict(
        template_path=os.path.join(os.path.dirname(__file__), "templates"),
	)               

application = tornado.web.Application(handlers, **settings)

if __name__ == "__main__":
    application.listen(8888)
    tornado.ioloop.IOLoop.instance().start()

#include <stdio.h>
#include <stdlib.h>

static char sudoku[9][9] = {
		{8,0,0, 0,0,0, 0,0,0},
		{0,0,3, 6,0,0, 0,0,0},
		{0,7,0, 0,9,0, 2,0,0},

		{0,5,0, 0,0,7, 0,0,0},
		{0,0,0, 0,4,5, 7,0,0},
		{0,0,0, 1,0,0, 0,3,0},

		{0,0,1, 0,0,0 ,0,6,8},
		{0,0,8, 5,0,0, 0,1,0},
		{0,9,0, 0,0,0, 4,0,0}


};
static int mitu = 0;

int kontrolli_kas_on_sudoku(int rida, int veerg);
int lahenda();
int kontrolli_sektor(int i);
int kontrolli_rida(int rida);
int kontrolli_veerg(int veerg);
void print_bitmap(int i);


int main(int argc,char **argv) {
	int i,j,x;

	if (kontrolli_kas_on_sudoku(10,10)) {
		i=lahenda();
		if (i<10) {
			printf ("Lahendasin sudoku, kordi %d\n",mitu);
			for (j=0;j<9;j++) {
				for (x=0;x<9;x++) {
					printf ("%hhd,",sudoku[j][x]);
				}
				printf("\n");
			}

		}
	} else {
		printf("Vigane sudoku!\n");
	}

	return 0;
}

int kontrolli_kas_on_sudoku(int rida, int veerg) {
// 0 - pole sudoku
// 1 - on sudoku 

	int i;
	
	if (rida<9) {
		if (!kontrolli_rida(rida)) {
			return 0;
		}
	} else {
		for (i=0;i<9;i++) {
			if (!kontrolli_rida(i)) {
				return 0;
			}	
		}
	}
	if (veerg<9) {
		if (!kontrolli_veerg(veerg)) {
			return 0;
		}
	} else {
		for (i=0;i<9;i++) {
			if (!kontrolli_veerg(i)) {
				return 0;
			}	
		}
	}
	if (rida<9 && veerg<9) {
		if (!kontrolli_sektor((int)(rida/3)+(3*(int)(veerg/3)))){
			return 0;
		}
	} else  {
		for (i=0;i<9;i++) {
			if (!kontrolli_sektor(i)) {
				return 0;
			}
		}
	}
	return 1;
} 

int lahenda(){
	mitu++;
	int x,y,nr;
	for (x=0;x<9;x++) {
		for (y=0;y<9;y++) {
			if (!sudoku[x][y]) {
				nr=1;
				while (nr<10) {
					sudoku[x][y]=nr;
					if (kontrolli_kas_on_sudoku(x,y)) {
						if  (lahenda()) {
							return 2;
						}
					}
					nr++;
				}
				sudoku[x][y]=0;
				return 0;
			}
		}
	}
	return 2;
}

int kontrolli_sektor(int i) {
	int read,veerud;
	int x,y,j,z;
	
	read=i%3;
	veerud=i/3;
	
	//printf("Sektor %d\n",i);
	z=0;
	for (x=0;x<3;x++) {
		for (y=0;y<3;y++) {
			//printf("%hhd ",sudoku[veerud*3+x][read*3+y]);
			j=1<<(sudoku[veerud*3+x][read*3+y]);
			if (sudoku[veerud*3+x][read*3+y]>0 && (z & j)>1) 
			{
				//printf ("paha sektor %d\n",i);
				return 0;
			} else {
				z=z | 1<<(sudoku[veerud*3+x][read*3+y]);
			}
		}
		//printf("\n");
	}
	return 1;
}
int kontrolli_rida(int rida) {
	int i,j,x;

	if (rida<0 || rida>8) {
		//printf("paha rida\n");
		exit(1);
	}

	x=0;
	for (i=0;i<9;i++) {
		j=1<<(sudoku[rida][i]);
	//	printf ("sudoku: %d %d %hhd %d\n",rida,i,sudoku[rida][i],j);
		if (sudoku[rida][i]>0 && (x & j)>1) 
		{
			//printf ("paha rida %d\n",rida);
			return 0;
		} else {
			x=x | 1<<(sudoku[rida][i]);
		}
	}
	//print_bitmap(x);
	return 1;
}
int kontrolli_veerg(int veerg) {
	int i,j,x;

	if (veerg<0 || veerg>8) {
		//printf("paha veerg\n");
		exit(3);
	}

	x=0;
	for (i=0;i<9;i++) {
		j=1<<(sudoku[i][veerg]);
		//printf ("sudoku: %d %d %hhd %d\n",i,veerg,sudoku[i][veerg],j);
		if (sudoku[i][veerg]>0 && (x & j)>1) 
		{
			//printf ("paha veerg %d\n",veerg);
			return 0;
		} else {
			x=x | 1<<(sudoku[i][veerg]);
		}
	}
	//print_bitmap(x);
	return 1;
}

void print_bitmap(int i) {
	int j;
	for (j=10;j>0;j--) {
		if (1<<j & i ) {
			printf("1");
		} else {
			printf("0");
		}
	}
	printf ("\n");
}

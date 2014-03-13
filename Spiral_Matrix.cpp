#include <vector>
#include <iostream>

#include <iterator>


#define M 3
#define N 4


class Solution
{
public:
  	static std::vector<int> spiralOrder(std::vector<std::vector<int> > &matrix) 
	{      
		std::vector<int> temp; 
		if(matrix.size() == 0)		    
			return temp;
		if(matrix.size() == 1 && matrix[0].size() == 1)
		{
			temp.push_back(matrix[0][0]);
			return temp;
		}
		int row = M, col = N;
		int k = 0, j = 0;
		
		for(int i = 0; i < row; i++)
		{
			for(j = i; j < col; j++)				
			{
				std::cout<< "A:" << i << " " << j << std::endl;
				temp.push_back(matrix[i][j]);	
			}
			for(k = i+1; k < row && col-1 > 0; k++)				
			{
				temp.push_back(matrix[k][col-1]);	
				std::cout<< "B:"  << k << " " << col-1 << std::endl;
			}
			for(j = col - 2; j > i && row-1 != i; j--)				
			{
				temp.push_back(matrix[row-1][j]);	
				std::cout<< "C:"  << row-1 << " " << j << std::endl;
			}
			for(k = row - 1; k > i && col -1 != i && col-1 > 0; k--)				
			{
				temp.push_back(matrix[k][i]);
				std::cout<< "D:"  << k << " " << i << std::endl;
			}
			row--;
			col--;
		}
		return temp;
	}

 std::vector<int> spiralOrder2(std::vector<std::vector<int> > &matrix) 
	{      
		std::vector<int> temp;
		SO_2(matrix, temp,  0, 0, matrix.size(), matrix[0].size());
		return temp;
	}

	void SO_2(std::vector<std::vector<int> > &matrix, std::vector<int> &temp, int sr, int sc, int row, int col)
	{
		int k = 0, j = 0;
		std::cout<<  "R:" << row << " " << col << " " << sr  << " " << sc << "\n"; 
		if(row <= 0 || col <= 0) return;
		if(row == 1 && col == 1)
		{
			temp.push_back(matrix[sr][sc]);
			std::cout<< "E:" << sr << " " << sc << " " << matrix[sr][sc] << std::endl;
			return;
		}
		else if(row == 1)
		{
		    for(int i = sc; i < sc+col; i++)
		    {
		    	std::cout<< "F:" << sr << " " << i << " " << matrix[sr][i] << std::endl;
		    	temp.push_back(matrix[sr][i]);
		    }
		}
		else if(col == 1)
		{
		    for(int i = sr; i < sr+row; i++)
		    {
		    	temp.push_back(matrix[i][sc]);
				std::cout<< "G:" << i << " " <<sc << " " << matrix[i][sc] << std::endl;
		    }
		}
		else
		{
			for(j = sc; j < sc+col; j++)				
			{
				std::cout<< "A:" << sc << " " << j << " " << matrix[sr][j] << std::endl;
				temp.push_back(matrix[sr][j]);	
			}
			for(k = sr + 1; k < sr+row; k++)				
			{
				temp.push_back(matrix[k][sr+col-1]);	
				std::cout<< "B:"  << k << " " << col-1<< " " << matrix[k][sr+col-1] << std::endl;
			}
			for(j = sc+col - 2; j >= sr; j--)				
			{
				temp.push_back(matrix[sr+row-1][j]);	
				std::cout<< "C:"  << row-1 << " " << j << " " << matrix[sr+row-1][j] << std::endl;
			}
			for(k = sr+row - 2; k > sc; k--)				
			{
				temp.push_back(matrix[k][sc]);
				std::cout<< "D:"  << k << " " << sc  << " " << matrix[k][sc] << std::endl;
			}
			std::cout<<  "M:" << row << " " << col << "\n"; 
			SO_2(matrix, temp, sr+1, sc+1, row-2, col-2);
			
		}
	}

};

int main()
{
	std::vector<std::vector<int> > example(M, std::vector<int>(N, -1));
	std::ostream_iterator<int> os(std::cout, "  ");
	int count = 1;
	for(int i=0;i<example.size();i++)
	{
		for(int j = 0; j < example[i].size(); j++)
			example[i][j] = count++;
        copy(example[i].begin(),example[i].end(),os);
        std::cout<<std::endl ;
    }
	std::cout<< std::endl;
	//std::vector<int> ret = Solution::spiralOrder(example);
	//copy(ret.begin(), ret.end(), os);
	std::cout << std::endl;

	Solution ss;
	std::vector<int> ret2 = ss.spiralOrder2(example);
	copy(ret2.begin(), ret2.end(), os);
	std::cout << std::endl;
}


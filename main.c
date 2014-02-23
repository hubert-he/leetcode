#include "Maximum_depth_of_binary_tree.h"
#include <iostream>
int main()
{
	BinTree aaa("abc@@d@@e@fg@@h@@");
	Solution bb;
	int ret = bb.maxDepth(aaa.getTree());
	std::cout << "xxx = " << ret << std::endl;
	return 0;
}

#include "Maximum_depth_of_binary_tree.h"
#include "sametree.h"
#include <iostream>
int main()
{
#ifdef Max_Depth
	BinTree aaa("abc@@d@@e@fg@@h@@");
	Solution bb;
	int ret = bb.maxDepth(aaa.getTree());
	std::cout << "xxx = " << ret << std::endl;
#endif
#ifdef Simple_Tree
	BinTree aaa("abc@@d@@e@fg@@h@@");
	//aaa.printTree(1);
	BinTree ccc("abc@@d@@e@fg@@h@@");
	BinTree ddd("abd@@d@@e@fg@@h@@");
	BinTree eee("abc@@d@@e@fg@@m@@");
	BinTree fff("abc@@de@@@f@@");
	Solution xx;
	int ret = xx.isSameTreebyStack(aaa.getTree(), ccc.getTree());
	std::cout << "\n++ret0 = " << ret << std::endl;
	ret = xx.isSameTreebyStack(aaa.getTree(), ddd.getTree());
	std::cout << "\n++ret0 = " << ret << std::endl;
	ret = xx.isSameTreebyStack(aaa.getTree(), eee.getTree());
	std::cout << "\n++ret0 = " << ret << std::endl;
	ret = xx.isSameTreebyStack(aaa.getTree(), fff.getTree());
	std::cout << "\n++ret0 = " << ret << std::endl;
#endif
	return 0;
}

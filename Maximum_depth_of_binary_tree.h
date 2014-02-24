#ifndef Maximum_depth_of_binary_tree_H
#define Maximum_depth_of_binary_tree_H
struct TreeNode
{
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x): val(x), left(0), right(0) {}
};

class BinTree
{
	private:
		char *str;
		TreeNode *tree;
		void createTree();
		void createTree2();
	public:
		BinTree(char *formate);
		TreeNode * getTree() { return tree; }
};
#endif
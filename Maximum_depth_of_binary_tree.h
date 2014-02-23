struct TreeNode
{
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x): val(x), left(0), right(0) {}
};

class Solution
{
	public:
		int maxDepth(TreeNode *root)
		{
			int max0 = 0, max1 = 0;
			if(root == 0) return -1;
			if(root->left)
			{
				max0 = maxDepth(root->left);
			}
			if(root->right)
			{
				max1 = maxDepth(root->right);
			}
			return ( (max0>max1)? max0+1: max1+1);
		}
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

#include <iostream>
#include "Maximum_depth_of_binary_tree.h"
#include <stack>
#include <string.h>
#define FALSE 0



BinTree::BinTree(char *formate)
{
    unsigned int len = strlen(formate);
    if(len <= 0) str = NULL;
    else
    {
        str = new char[len];
        memcpy(str, formate, len);
    }
    tree = NULL;
    createTree2();
}
void BinTree::printTree(int opt)
{
    if(!tree) return;
	std::cout << "Tree: ";
	switch(opt)
	{
	case 0:
		pre();
		break;
	case 1:
		mid();
		break;
	case 2:
		post();
		break;
	default:
		break;
	}
	std::cout << std::endl;
}

void BinTree::pre()
{

}

void BinTree::post()
{
}

void BinTree::mid()
{
    std::stack<TreeNode *> st;
    TreeNode *tp = tree;
    while(tp || !st.empty())
    {
        if(tp)
        {
            // tp = NULL means that you should turn right right now
            st.push(tp);
            tp = tp->left;
        }
        else
        {
        	tp = st.top()->right;
        	std::cout << " " << st.top()->val;
			st.pop();
        }
    }
}
void BinTree::createTree()
{
    if(str == NULL) return;
    std::stack<TreeNode*> st;
    bool isCreateR = FALSE;
    unsigned int len = strlen(str), i;
    TreeNode *tp;
    TreeNode *tpp;
    for(i = 0; i < len; i++)
    {
        if(str[i] == '@')
        {
            if(!isCreateR)
                isCreateR = 1;
            else
            {
                st.pop();
                if(!st.empty())
                    tpp = st.top();
            }
        }
        else
        {
            tp = new TreeNode(str[i]);
            if(i == 0)
            {
                tree = tp;
            }
            else
            {
                if(!isCreateR)
                    tpp->left = tp;
                else
                {
                    tpp->right = tp;
                    isCreateR = 0;
                    st.pop();
                }
            }
            st.push(tp);
            tpp = tp;
        }

    }
}
/*
**
    左：
        为空(==@)
            符号变量赋值 向右
        非空
            链入二叉树， 入栈
    右： // 一旦向右 必出栈
        为空(==@)
            出栈
        非空
            出栈，链入新元素到二叉树，入栈新元素， 符号变量赋值 向左
**
*/
void BinTree::createTree2()
{
    if(str == NULL) return;
    std::stack<TreeNode*> st;
    bool isCreateR = FALSE;
    unsigned int len = strlen(str), i;
    TreeNode *tp;
    TreeNode *tpp;

    tree = new TreeNode(str[0]);
    st.push(tree);
    for(i = 1; i < len; i++)
    {
        if(str[i] == '@')
        {
            // NULL
            if(!isCreateR)
            {
                // 左
                isCreateR = 1;
            }
            else
            {
                // 右
                st.pop();
            }
        }
        else
        {
            // !NULL
            tp = new TreeNode(str[i]);
            if(!isCreateR)
            {
                // 左
                tpp = st.top();
                tpp->left = tp;
                st.push(tp);
            }
            else
            {
                // 右
                tpp = st.top();
                st.pop();
                tpp->right = tp;
                st.push(tp);
                isCreateR = 0;
            }
        }

    }
}



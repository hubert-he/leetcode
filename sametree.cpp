#include <stack>
#include "sametree.h"
#include <iostream>
bool Solution::isSameTreebyStack(TreeNode *p, TreeNode *q)
{
#if 1
    // Note: The Solution object is instantiated only once and is reused by each test case.
    std::stack<TreeNode*> st0, st1;
    TreeNode *node0 = p, *node1 = q;
    while(node0 || !st0.empty())
    {
        if(node0)
        {
            if(node0)
            {
            	if(node1) return 0;
            	if(node0->val != node1->val) return 0;
            }
            else return 0;
			// go on left
			st0.push(node0);
			st1.push(node1);
			node0 = node0->left;
			node1 = node1->left;
        }
        else
        {
        	if(node1) return 0;
        	node0 = st0.top();
			st0.pop();
			node0 = node0->right;
			if(st1.empty()) return 0;
			node1 = st1.top();
			st1.pop();
			node1 = node1->right;
        }
    }
    if(st0.empty() && st1.empty())
        return 1;
    else return 0;
#endif
#if 0
    if(node0 && node1)
    {
        // !NULL
        if(node0->val != node1->val) break;
        if(node0->left && node1->left)
        {
            st0.push(node0->left);
            st1.push(node1->left);
        }
        else if(!node0->left && !node1->left)
        {
            st0.pop();
            st1.pop();
            if(node0->right && node1->right)
            {
                st0.push(node0->right);
                st1.push(node1->right);
            }
            else if(!node0->right && !node1->right)
            {
                st0.pop();
                st1.pop();
            }
            else break;
        }
        else break;
    }
    else break;
}
if(!st0.empty() || !st1.empty())
    return 0;
else return 1;
#endif
}

bool Solution::isSameTreebyRecursion(TreeNode *p, TreeNode *q)
{
    bool left = 0, right = 0;
    if(p && q)
    {
        //std::cout << "var=" << p->val << " " << q->val << std::endl;
        if(p->val == q->val)
        {
            left = isSameTreebyRecursion(p->left, q->left);
            //std::cout << "left=" << left << std::endl;
            if (left == 0) return 0;
            right = isSameTreebyRecursion(p->right, q->right);
            //std::cout << "right=" << right << std::endl;
            if(right == 0) return 0;
            return 1;
        }
        else return 0;
    }
    else if(!p && !q) return 1;
    else return 0;
}


#include <stdio.h>

int searchInsertByseq(int A[], int n, int target)
{
    int i;
    for(i = 0; i < n; i++)
    {
    	//printf("%d:%d\n", target, A[i]);
    	if(A[i] < target) continue;
    	if(A[i] >= target) return i;
    }
}

int searchInsert(int A[], int n, int target)
{
  	int l, r, mid;
	for(l = 0, r = n-1, mid = (r+l)/2; l < r; mid = (r+l)/2)
	{
		if(A[mid] == target) return mid;
		else if(A[mid] > target)
			{
				r = mid - 1;
			}
		else
			l = mid + 1;
	}
	if(A[mid] >= target) return mid;
	else return mid + 1;
}

int main()
{

    int a[4] = {1,3,5,6};
    int ret;
#if 1
    ret = searchInsert(a, 4, 5);
    printf("ret:%d\n", ret);
    ret = searchInsert(a, 4, 2);
    printf("ret:%d\n", ret);
    ret = searchInsert(a, 4, 7);
    printf("ret:%d\n", ret);
    ret = searchInsert(a, 4, 0);
    printf("ret:%d\n", ret);
#endif
	int b[1] = {1};
	ret = searchInsertByseq(b, 1, 1);
	printf("ret: %d\n", ret);
}


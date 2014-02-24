#include <stdio.h>
#include <math.h>
int main()
{
	int x;
	int negative = 0;
	unsigned int ret = 0;
	scanf("%d", &x);
	if(x < 0)
	{
		x = -x;
		negative = 1;
	}
	unsigned int count = 0;
	unsigned int xx = x;
	while(xx/10 > 0)
	{
		count++;
		xx = xx/10;
	}
	//if(xx > 0) count++;
	printf("%d\n", count);
	while(count > 0)
	{
		ret += (x%10)*(pow(10,count));
		count--;
		x = x/10;
	}
	if(x>0) ret += x;
	printf("%d %u\n", x, negative?-ret:ret);
}


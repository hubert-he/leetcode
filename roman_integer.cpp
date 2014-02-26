#include <iostream>
#include <string>

#define I 1
#define V 5
#define X 10
#define L 50
#define C 100
#define D 500
#define M 1000

int val[]= {1000,900,500,400,100,90,50,40,10,9,5,4,1};
std::string r[]= {"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"};

std::string intToRoman(int num)
{
    std::string roman_num;
    int k = num, i = 0;
    for(k = num, i = 0; k > 0; i++)
    {
        while(k>=val[i])
        {
            roman_num += r[i];
			std::cout << "+ " << roman_num << " -" << k << std::endl;
            k -= val[i];
        }
    }
	return roman_num;
}

int romanToInt(std::string s)
{
    int ret = 0;
    for(int i = 0; i < s.length(); i++)
    {
        std::cout << "ret=" << ret << std::endl;
        switch(s[i])
        {
            case 'I':
                ret += I;

                if(s[i+1] == 'V')
                {
                    ret += V-I-I;
                    i++;
                }
                else if(s[i+1] == 'X')
                {
                    ret += X-I-I;
                    i++;
                }
                else if (s[i+1] != 'I' )
                {
                    std::cout << "ERROR:wrong Format!!!\n";
                    return 0;
                }
				std::cout << "++ret=" << ret << std::endl;
                break;
            case 'X':
                ret += X;
                std::cout << "+ret=" << ret << std::endl;

                if(s[i+1] == 'L')
                {
                    ret += (L-X-X);
                    i++;
                }
                else if(s[i+1] == 'C')
                {
                    ret += (C-X-X);
                    i++;
                }
                else if (s[i+1] != 'X' && s[i+1] != 'I')
                {
                    std::cout << "ERROR:wrong Format!!!\n";
                    return 0;
                }
                std::cout << "+ret=" << ret << std::endl;
                break;

            case 'V':
                ret += V;
                break;
            case 'C':
                ret += C;
                if(s[i+1] == 'D')
                {
                    ret += D-C-C;
                    i++;
                }
                else if(s[i+1] == 'M')
                {
                    ret += M-C-C;
                    i++;
                }
                else if (s[i+1] != 'X' && s[i+1] != 'C')
                {
                    std::cout << "ERROR:wrong Format!!!\n";
                    return 0;
                }
                break;

            case 'L':
                ret += L;
                break;
            case 'D':
                ret += D;
                break;
            case 'M':
                ret += M;
                break;
            default:
                return 0;
        }

    }
    return ret;
}

void romanToInt(std::string s, int *p)
{
    int ret = 0, j = 0;
    for(int i = 0; i < s.length(); i++)
    {
        switch(s[i])
        {
            case 'I':
                ret += I;
                if(s[i+1] == 'V')
                {
                    ret += V - I - I;
                    i++;
                }
                else if(s[i+1] == 'X')
                {
                    ret += X - I - I;
                    i++;
                }
                break;
            case 'X':
                ret += X;
                if(s[i+1] == 'L')
                {
                    ret += L - X - X;
                    i++;
                }
                else if(s[i+1] == 'C')
                {
                    ret += C - X - X;
                    i++;
                }
                break;
            case 'V':
                ret += V;
                break;
            case 'C':
                ret += C;
                if(s[i+1] == 'D')
                {
                    ret += D - C - C;
                    i++;
                }
                else if(s[i+1] == 'M')
                {
                    ret += M - C - C;
                    i++;
                }
                break;
            case 'L':
                ret += L;
                break;
            case 'D':
                ret += D;
                break;
            case 'M':
                ret += M;
                break;
            default:
                *p = 0;
                return;
        }
    }
    *p = ret;
}
int main()
{
    std::string s("XCIX");
    //std::string s("MDCCCIC");
    //std::string s("XXIV");
    //std::string s("MDCCCXCIX");
    int ret = romanToInt(s);
    std::cout << "ret=" << ret << std::endl;
    romanToInt(s, &ret);
    std::cout << "ret=" << ret << std::endl;
	std::string xx = intToRoman(15);
	std::cout << "roman = " << xx << std::endl;
    return 0;
}

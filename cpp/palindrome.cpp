#include <iostream>
#include <cstring>

using namespace std;

const int maxn=1000;

void fillP(char s[maxn],char se[], int p[]){
	int slen=strlen(s);
	for (int i = 0; i < slen; ++i)
	{
		se[2*i]='#';
		se[2*i+1]=s[i];
	}
	se[2*slen]='#';
	se[2*slen+1]='\0';

	slen=strlen(se);
	int R=0,C=0;
	for (int i = 1; i < slen-1; ++i)
	{
		int i_mirr=2*C-i;
		p[i]=(R>i)?min(p[i_mirr],R-i):0;

		while(i-1-p[i]>=0 && se[i+1+p[i]]==se[i-1-p[i]]) p[i]++;

		if(i+p[i]>R){
			C=i;
			R=i+p[i];
		}
	}
}

int main(){
	char s[maxn]="aabbaa";

	int p[2*maxn];
	char se[2*maxn];
	fillP(s,se,p);
	cout<<se<<"\n";
	int len=strlen(se);
	for (int i = 0; i < len; ++i)
	{
		cout<<p[i];
	}
	cout<<"\n";

	return 0;
}

#include <iostream>
#include <cstring>
#include <map>
#include <vector>
#include <algorithm>
#include <cstdio>
#include <cmath>

using namespace std;

/*
//naive n^3 approach
int main(){
	map<string,int> m;
	string s("akshit");
	vector<string> v;
	for (int i = 0; i < s.size(); ++i)
	{
		v.push_back(s.substr(i, s.size()-i));
		m[s.substr(i,s.size()-i)]=i;
	}
	sort(v.begin(),v.end());
	for (int i = 0; i < v.size(); ++i)
	{
		cout<<m[v[i]]<<"\n";
	}
}
*/

const int maxn=100000;
const int maxlg=17;

struct entry
{
	int nr[2];
	int p;
}L[maxn];

int P[maxlg][maxn];
int sA[maxn];
int lcp[maxn];//gives lcp length b\w prefix i and i+1 in sorted suffix array

bool cmp(entry a, entry b){
	return a.nr[0]==b.nr[0] ? (a.nr[1]<b.nr[1]?1:0) : (a.nr[0]<b.nr[0]?1:0);
}

int findLCP(int x,int y,int n){
	int stps=ceil(log2(n));
	int len=0;
	for (int k = stps; k >=0; --k)
	{
		if(P[k][x]==P[k][y]){
			len+=1<<k;
			x+=1<<k;
			y+=1<<k;
		}
	}
	return len;
}

void suffixArray(char s[maxn]){
	int N=strlen(s);
	for (int i = 0; i < N; ++i)
	{
		P[0][i]=s[i]-'a';
	}

	int cnt,stp;
	for(cnt=1, stp=1; cnt<N; cnt<<=1,stp++){
		for (int i = 0; i < N; ++i)
		{
			L[i].nr[0]=P[stp-1][i];
			L[i].nr[1]=i+cnt <N ? P[stp-1][i+cnt] : -1;
			L[i].p=i;
		}
		sort(L,L+N,cmp);
		for (int i = 0; i < N; ++i)
		{
			P[stp][L[i].p] = i>0 && L[i].nr[0]==L[i-1].nr[0] && L[i].nr[1]==L[i-1].nr[1] ? P[stp][L[i-1].p] : i;
		}
	}


	for (int i = 0; i < N; ++i)
	{
		sA[i]=L[i].p;//sA[i] stores index of the ith largest suffix	
	}

	//finding lcp
	for (int i = 0; i+1 < N; ++i)
	{
		lcp[i]=findLCP(sA[i],sA[i+1],N);
	}
}

int main(){
	char s[maxn]="mississippi";
	suffixArray(s);
	int N=strlen(s);
	for (int i = 0; i < N; ++i)
	{
		printf("%s-%d\n",s+sA[i],lcp[i]);
	}

	return 0;

}

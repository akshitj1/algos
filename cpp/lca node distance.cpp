//https://www.codechef.com/COOK38/problems/RRTREE

#include <iostream>
#include <cstdio>
#include <algorithm>

using namespace std;

#define two(x) (1<<(x))
#define contain(S,x) (((S)&two(x))!=0)

const int maxn=100000+10;

int n;
int p[maxn];
int d[maxn];
int f[maxn][20]; //ceil(log(maxn))

int getf(int key,int len){
	
	for (int k=0; two(k) <= len ; k++)
		if(contain(len,k)) key=f[key][k];

	return key;
}

int get_lca(int a,int b){
	if(d[a]>d[b]) swap(a,b);
	b=getf(b,d[b]-d[a]);
	int H=-1,T=d[a];
	while(H+1<T){
		int M=(H+T)/2;
		if(getf(a,M)==getf(b,M)) T=M;
		else H=M;
	}
	return getf(a,T);
}

int get_dist(int a, int b){
	return d[a]+d[b]-2*d[get_lca(a,b)];
}

void process(){
	scanf("%d",&n);
	//cin>>n;
	for (int i = 1; i < n; ++i){
		scanf("%d",&p[i]);
		//cin>>p[i];
		p[i]--;
	}
	if(n<=1) return;

	d[0]=0;
	for (int i = 1; i < n; ++i)
		d[i]=d[p[i]]+1;

	for (int i = 0; i < n; ++i)
	{
		f[i][0]=p[i];
		for (int k = 1; two(k)<= d[i]; ++k)
			f[i][k]=f[f[i][k-1]][k-1];
	}

	int p1=0;
	int p2=1;
	int length=1;
	cout<<1<<"\n";
	for (int i = 2; i < n; ++i)
	{
		int l1=get_dist(i,p1);
		int l2=get_dist(i,p2);
		if(l1>=l2 && l1>length){
			p2=i;
			length=l1;
		}else if(l2>=l1 && l2>length){
			p1=i;
			length=l2;
		}
		printf("%d\n", length);
		//cout<<length<<"\n";
	}
}

int main(){
	int t;
	cin>>t;
	while(t--) process();
}

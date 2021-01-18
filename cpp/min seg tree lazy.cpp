#include <iostream>
#include <cstring>
#include <map>
#include <vector>
#include <algorithm>
#include <cstdio>
#include <cmath>

using namespace std;

#define pb push_back;

typedef pair<int,int> ii;
template<class T> checkmax(T &a,T b){ if(b>a) a=b;}

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
int rank[maxn];

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

	stp--;
	for (int i = 0; i < N; ++i)
	{
		rank[i]=P[stp][i];
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

int tree[1<<20];

//seg tree lazy update with no < operator as elements comming in decreasing order
void update(int p,int s,int t, int l, int r, int key){
	if(l<=s && t<=r) {tree[p=key]; return;}
	if(tree[p]>0){
		tree[2*p]=tree[2*p+1]=tree[p];
		tree[p]=-1;
	}
	int m=(s+t)/2;
	if(l<=m) update(2*p+1,s,m,l,r,key);
	if(m<r) update(2*p+2,m+1,t,l,r,key);//why not equality bcs m+1
}

void getAll(int p,int s, int t,int r[]){
	if(tree[p]>0){
		for (int i = s; i <= t; ++i) r[i]=tree[p];
		return;
	}

	if(s==t) {r[s]=-1;return;}
	int m=(s+t)/2;
	getAll(2*p+1,s,m,r);
	getAll(2*p+2,m+1,t,r);
}

int main(){
	char s[maxn]="mississippi";
	suffixArray(s);
	int N=strlen(s);
	for (int i = 0; i < N; ++i)
	{
		printf("%s-%d\n",s+sA[i],lcp[i]);
	}

	vector<ii> q;

	for (int i = 0; i < N; ++i)
		q.pb(ii(lcp[i],i));

	memset(tree,255,sizeof(tree));//maximum negative value
	sort(q.begin(),q.end());
	for (int i = N-1; i >= 0; --i)
	{
		int pos=q[i].second;
		int p=sA[pos];
		if(p+lcp[pos]-1<N) update(0,0,N-1,p,p+lcp[pos]-1,p);//why -1
	}
	int pos[maxn];
	getAll(0,0,N-1,pos);


	int g[maxn];
	for (int i = 0; i < N; ++i) pos[i]<0?N+1:lcp[rank[pos[i]]];//why n+1 not n
	
	memset(g,255,sizeof(g));
	for (int i = 0; i < N; ++i) g[sA[i]+lcp[i]]=sA[i];

	for (int t=-1,i = 0; i < n; ++i)
	{
		checkmax(t,g[i]);
		if(t<0) continue;
		int l=i-t+1;
		if(l<length[i] || (l==length[i] && rank[t]<rank[pos[i]])) pos[i]=t,length[i]=l;
	}
	return 0;
}

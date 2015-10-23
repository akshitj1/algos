#include <iostream>
#include <queue>
#include <vector>
#include <cstring>
#include <map>

using namespace std;

#define pb push_back
#define clr(a) memset(a, 0, sizeof(a))
#define forl(i,l,u) for(int i=(l);i<(u);i++)
#define forn(i,n) forl(i,0,n)

typedef pair<int, int> ii;
typedef map<int, int> mp;

const int maxn=200000+11;
const int maxans=1000*maxn;

struct sNode
{
	int p, rank, val;
} nd[maxn];

void createSet(int n){
	forn(i, n+1)
		nd[i].rank=0, nd[i].p=i;
}

int findSet(int x){
	if(nd[x].p!=x) nd[x].p=findSet(nd[x].p);
	return nd[x].p;
}

void mergeSet(int x, int y){
	int px=findSet(x);
	int py=findSet(y);
	if(px==py)
		return;
	if(nd[px].rank > nd[py].rank)
		nd[py].p=px, nd[px].val+=nd[py].val;
	else
		nd[px].p=py, nd[py].val+=nd[px].val;
	
	if(nd[px].rank == nd[py].rank)	nd[py].rank++;
}

int main(){
	int n, q;
	cin>>n>>q;
	createSet(n);
	forl(i, 1, n+1){
		cin>>nd[i].val;
	}
	while(q--){
		int x, y;
		cin>>x>>y;
		mergeSet(x,y);
	}
}
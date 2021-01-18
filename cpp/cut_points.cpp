/* find articulation/cur points
implmentation: http://e-maxx.ru/algo/cutpoints,
tutorial: http://www.geeksforgeeks.org/articulation-points-or-cut-vertices-in-a-graph/
5 5
1 2
2 3
1 3
3 4
4 5
*/
#include <iostream>
#include <vector>

using namespace std;

#define pb push_back
#define forl(i,l,u) for(int i=(l);i<(u);i++)
#define forn(i,n) forl(i,0,n)

const int maxn=1e5+11;
int timer=0;

struct Node{
	vector<int> adj;
	bool vis, isAP;
	int dTim, lDTim;
	Node(){
		vis=isAP=false;
	}
}gr[maxn];

void dfs(int u, int p=-1){
	gr[u].vis=true;
	gr[u].dTim=gr[u].lDTim=timer++;
	int children=0;
	forn(i, gr[u].adj.size()){
		int v=gr[u].adj[i];
		if(v==p)	continue;
		if(!gr[v].vis){
			children++;
			dfs(v, u);
			gr[u].lDTim=min(gr[v].lDTim, gr[u].lDTim);
			if(gr[v].lDTim >= gr[u].dTim && p!=-1)
				gr[u].isAP=true;
		}else
			gr[u].lDTim=min(gr[v].lDTim, gr[u].lDTim);
	}
	if(p==-1 && children>1)
		gr[u].isAP=true;
}

int main(){
	int n, m;
	cin>>n>>m;
	forn(i, m){
		int u, v;
		cin>>u>>v;
		gr[u].adj.pb(v);
		gr[v].adj.pb(u);
	}
	timer=0;
	dfs(1);
	forl(i, 1, n+1)
		if(gr[i].isAP)	cout<<i<<" ";
}
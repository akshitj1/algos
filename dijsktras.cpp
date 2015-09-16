/* test input
6 9 1
1 2 7
1 6 14
1 3 9
2 3 10
2 4 15
3 6 2
3 4 11
6 5 9
4 5 6
*/

#include <iostream>
#include <queue>
#include <vector>
#include <cstring>
#include <functional>//for greater

using namespace std;

#define pb push_back
#define clr(a) memset(a, 0, sizeof(a))
#define forl(i,l,u) for(int i=(l);i<(u);i++)
#define forn(i,n) forl(i,0,n)

typedef pair<int,int> ii;

const int maxn=1000+11;
const int maxw=1000+11;
const int maxd=maxn*maxw;

vector<ii> gr[maxn];
int minDist[maxn];
bool vis[maxn];

void djInit(int n, int (&minDist)[maxn]){
	forn(i,n)
		minDist[i]=maxd;
	clr(vis);
}

void dijsktras(int s, int n,const vector<ii> (&gr)[maxn], int (&minDist)[maxn]){
	minDist[s]=0;
	djInit(n, minDist);
	priority_queue<ii, vector<ii>, greater<ii> > pq;
	pq.push(ii(0, s));
	while(!pq.empty()){
		ii top=pq.top();
		pq.pop();
		int u=top.second;
		if(vis[u]) continue;
		int d=top.first;
		if(minDist[u]>d)
			minDist[u]=d;
		forn(i, gr[u].size()){
			pq.push(ii(minDist[u]+gr[u][i].second, gr[u][i].first));
		}
		vis[u]=true;
	}
}

int main(){
	int n,m,s;
	cin>>n>>m>>s;
	s--;
	forn(i,m){
		int u,v,d;
		cin>>u>>v>>d; u--,v--;
		gr[u].pb(ii(v,d));
		gr[v].pb(ii(u,d));
	}
	dijsktras(s, n, gr, minDist);
	forn(i, n)
		cout<<s+1<<","<<i+1<<" : "<<minDist[i]<<endl;
}
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
int prev[maxn];
bool vis[maxn];

void djInit(int n, int (&minDist)[maxn], int (&prev)[maxn]){
	forn(i,n){
		minDist[i]=maxd;
		prev[i]=-1;
	}
	clr(vis);
}

void dijsktras(int s, int t, int n,const vector<ii> (&gr)[maxn], int (&minDist)[maxn], int (&prev)[maxn]){
	djInit(n, minDist, prev);
	minDist[s]=0;
	priority_queue<ii, vector<ii>, greater<ii> > pq;
	pq.push(ii(0, s));
	while(!pq.empty()){
		ii top=pq.top();
		pq.pop();
		int u=top.second;
		if(u==t)	break;
		if(vis[u]) continue;
		forn(i, gr[u].size()){
			int v=gr[u][i].first;
			int alt= minDist[u]+gr[u][i].second;
			if(minDist[v]>alt){
				minDist[v]=alt;
				prev[v]=u;
			}
			pq.push(ii(minDist[v], v));
		}
		vis[u]=true;
	}
	if(minDist[t]==maxd)
		return -1;
	return minDist[t];
}

ll prims(int n){
	priority_queue<ii, vector<ii>, greater<ii> > pq;
	memset(vis, 0, n*sizeof(bool));
	forn(i, n)
		dist[i]=maxw;
	dist[0]=0;
	pq.push(ii(0, 0));
	while(!pq.empty()){
		int u=pq.top().second; pq.pop();
		if(vis[u])
			continue;
		forn(i, adj[u].size()){
			int v=adj[u][i].first;
			int w=adj[u][i].second;
			if(!vis[v] && dist[v]>w){
				dist[v]=w;//insert prev for backtracking
				pq.push(ii(w, v));
			}
		}
		vis[u]=true;
	}
	//total weight of spanning tree
	ll ans=0;
	forl(i, 1, n)
		if(dist[i]==maxw) return -1;
		else ans+=dist[i];
	return ans;
}

//v2
ll dijsktras(int s,int t, int n){
	djInit(n);
	dist[s]=0;
	priority_queue<lip, vector<lip>, greater<lip> > pq;
	pq.push(lip(0, s));
	while(!pq.empty()){
		lip top=pq.top();
		pq.pop();
		int u=top.second;
		if(u==t)
			break;
		if(vis[u]) continue;
		forn(i, adj[u].size()){
			int v=adj[u][i].first;
			if(vis[v])	continue;
			ll alt=dist[u]+adj[u][i].second;
			if(dist[v]>alt){
				dist[v]=alt;
				pq.push(lip(dist[v], v));
			}
		}
		vis[u]=true;
	}
	if(dist[t]==maxd)
		return -1;
	return dist[t];
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
	dijsktras(s, n, gr, minDist, prev);
	forn(i, n)
		cout<<s+1<<"::"<<prev[i]+1<<" -> "<<i+1<<" : "<<minDist[i]<<endl;
}
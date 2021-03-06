//https://www.codechef.com/COOK39/problems/PPTREE

#include <iostream>
#include <vector>
#include <cstring>

using namespace std;

#define pb push_back
#define MEM(a,b) memset(a,(b),sizeof(a))

const int maxn=100000;

struct Node
{
	int v,wt;
	Node(int _v,int _wt){
		v=_v;
		wt=_wt;
	}
};

vector<Node> adj[maxn+5];
int rxor[maxn+5];
int trie[30*maxn][2];
//int trie[30*maxn][2];

int nid;

void insert_trie(int node,int d,int v){
	if(d==30)
		return;
	int p=29-d;
	int c=((1<<p)&v) ? 1:0;
	if(trie[node][c]==-1)
		trie[node][c]=++nid;

	insert_trie(trie[node][c],d+1,v);
}

int solve(int node,int d,int v){
	if(d==30)
		return 0;
	int p=29-d;
	int c=((1<<p)&v) ? 0:1;
	if(trie[node][c]!=-1)
		return (1<<p) + solve(trie[node][c],d+1,v);
	else
		return solve(trie[node][!c],d+1,v);
}

void dfs(int u,int p, int cxor){
	rxor[u]=cxor;
	for (int i = 0; i < adj[u].size(); ++i){
		if(adj[u][i].v!=p)
			dfs(adj[u][i].v,u,adj[u][i].wt^cxor);
	}
}

int main(){
	int tc;
	cin>>tc;
	while(tc--){
		int n;
		cin>>n;

		for (int i = 1; i <= n; ++i)
		{
			adj[i].clear();
		}

		for (int i = 0; i < n-1; ++i)
		{
			int u,v,wt;
			cin>>u>>v>>wt;
			adj[u].pb(Node(v,wt));
			adj[v].pb(Node(u,wt));
		}
		dfs(1,0,0);
		nid=0;
		MEM(trie,-1);
		for (int i = 1; i <= n; ++i)
		{
			insert_trie(0,0,rxor[i]);
		}

		int res=0;
		for (int i = 1; i <= n; ++i)
		{
			int q=solve(0,0,rxor[i]);
			res=q>res?q:res;
		}
		cout<<res<<"\n";
	}
}
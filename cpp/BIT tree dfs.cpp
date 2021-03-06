// BIT dfs tree tree to range
//https://www.codechef.com/COOK35/problems/TYTACTIC

#include <iostream>
#include <vector>

using namespace std;

#define MAX_V 111111

int num[MAX_V];
int vIdx[MAX_V];
vector<int> adj[MAX_V];
int weight[MAX_V];
int BIT[MAX_V];
int _idx=0;
int n;

void dfs(int u){
	vIdx[u]=++_idx;
	num[u]=1;
	for (int i = 0; i < adj[u].size(); ++i)
	{
		int v=adj[u][i];
		if(num[v]==0){
			dfs(v);
			num[u]+=num[v];
		}
	}
}

//add the difference
void update(int idx,int diff){
	while(idx <= n){
		BIT[idx] += diff;
		idx += idx & -idx;
	}
}

//get cumulative sum
int get(int idx){
	int sum=0;
	while(idx){
		sum+=BIT[idx];
		idx -= idx & -idx;
	}
	return sum;
}

int main(){
	int m;
	cin>>n>>m;
	for (int i = 1; i <= n; ++i)
		cin>>weight[i];

	int a,b;
	for (int i = 1; i < n; ++i)
	{
		cin>>a>>b;
		adj[a].push_back(b);
		adj[b].push_back(a);
	}
	dfs(1);

	//initialize BIT
	for (int i = 1; i <= n; ++i)
	{
		update(vIdx[i],weight[i]);
	}

	char type[3];
	while(m--){
		cin>>type;
		if(type[0]=='U'){
			int v,val;
			cin>>v>>val;
			update(vIdx[v],val-weight[v]);
			weight[v]=val;
		}else{
			int v;
			cin>>v;
			cout<<(get(vIdx[v]+num[v]-1) - get(vIdx[v]-1))<<"\n";
		}
	}

	return 0;
}
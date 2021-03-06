//FIND CYCLES
//https://www.codechef.com/problems/INVITES/

#include <iostream>
#include <vector>

//detect cycles
//dfs
//break cycle

using namespace std;
#define MAX_N 100005
#define pb(x) push_back(x)
#define MOD 1000000007
#define ll long long

vector<int> adj[MAX_N];
vector< vector<int> > cycles;
ll nCh[MAX_N];
ll Ch[MAX_N];
int enemy[MAX_N];
int visited[MAX_N];
bool inCycle[MAX_N];

int detectCycle(int v,vector<int> &C){
	if(visited[v]==2)
		return -1;
	if(visited[v]==1)
		return v;

	visited[v]++;
	int ret=detectCycle(enemy[v],C);
	if(ret!=-1){
		C.pb(v);
		inCycle[v]=true;
	}else{
		inCycle[v]=false;
	}
	visited[v]++;//pop
	if(ret==v)
		return -1;
	else
		return ret;
}

void dfs(int v,int p){

	nCh[v]=1;
	Ch[v]=1;
	for (int i = 0; i < adj[v].size(); ++i)
	{
		int w=adj[v][i];
		if( inCycle[w] || w==p)
			continue;

		dfs(w,v);
		nCh[v]= (nCh[v]*(nCh[w]+Ch[w]))%MOD;
		Ch[v]=(Ch[v] * nCh[w])%MOD;
	}
}

ll solveCycle(vector<int> cy){
	ll ans1=0;
	//choosing 1st vertex
	ll nch=0;
	ll ch=Ch[cy[0]];
	for (int i = 1; i < cy.size(); ++i)
	{
		int v=cy[i];
		ll _ch,_nch;
		if(i==cy.size()-1){
			_nch=((ch+nch)*nCh[v])%MOD;
			_ch=0;
		}else{
			_nch=((ch+nch)*nCh[v])%MOD;
			_ch=(nch*Ch[v])%MOD;
		}
		ch=_ch;
		nch=_nch;
	}
	ans1=(ch+nch)%MOD;

	//not choosing 1st vertex
	nch=nCh[cy[0]];
	ch=0;
	for (int i = 1; i < cy.size(); ++i)
	{
		int v=cy[i];
		ll _ch,_nch;
		_nch=((ch+nch)*nCh[v])%MOD;
		_ch=(nch*Ch[v])%MOD;
		ch=_ch;
		nch=_nch;
	}

	ll ans2=(nch+ch)%MOD;

	return (ans1+ans2)%MOD;
}

ll solveCycles(){
	ll ans=1;
	for (int i = 0; i < cycles.size(); ++i)
	{
		ans=(ans*solveCycle(cycles[i]))%MOD;
	}
	return ans;
}

void reinit(){
	for (int i = 0; i < MAX_N; ++i)
	{
		adj[i].clear();
		visited[i]=0;
		nCh[i]=Ch[i]=0;
	}
	cycles.clear();
}

int main(){
	int t;
	cin>>t;
	while(t--){
		int n;
		cin>>n;
		for (int i = 1; i <= n; ++i)
		{
			int v;
			cin>>v;
			enemy[i]=v;
			adj[v].pb(i);
			adj[i].pb(v);
		}

		//detect all cycles
		for (int i = 1; i <= n; ++i)
		{
			vector<int> tmp;
			detectCycle(i,tmp);
			if(tmp.size()>0)
				cycles.push_back(tmp);
		}

		//dfs for all cycle elements
		for (int i = 1; i <= n; ++i)
		{
			if(inCycle[i])
				dfs(i,i);
		}
		cout<<solveCycles()<<"\n";
		reinit();
	}
}
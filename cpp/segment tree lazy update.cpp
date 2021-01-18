//https://www.codechef.com/COOK38/problems/RRBIGNUM

#include <iostream>
#include <cmath>

using namespace std;

#define ll long long

int maxn = 1000005;
int SZ = 1<<(ceil(log(maxn)/log(2))+1);
const int mod=1000000007;

struct Node
{
	int l,r;

	ll my_sum;
	ll tot_sum;

	int parity;
	bool flipped;
};


void init(Node *tr, int v[], int n){
	tr=new Node[SZ];

	ll t=1;

	for (int i = 0; i < SZ/2; ++i)
	{
		//initialize leafs
		tr[i+SZ/2].l=i;
		tr[i+SZ/2].r=i;

		tr[i+SZ/2].tot_sum=t;

		if(i<v.size()){
			tr[i+SZ/2].my_sum=t*v[i];
			tr[i+SZ/2].parity= v[i] & 1;
		}else{
			//leaf does not exists
			tr[i+SZ/2].my_sum=0;
			tr[i+SZ/2].parity= 0;
		}

		tr[i+SZ/2].flipped=false;
		t=(t<<1)%mod;
	}

	for (int i = SZ/2; i >0 ; --i)//1 is the root
	{
		//initialize non leaf nodes
		int x=2*i;
		int y=x+1;

		tr[i].l=tr[x].l;
		tr[i].r=tr[y].r;

		tr[i].tot_sum=(tr[x].tot_sum + tr[y].tot_sum)%mod;
		tr[i].my_sum=(tr[x].my_sum + tr[y].my_sum)%mod;

		tr[i].parity=(tr[x].parity+tr[y].parity)&1;
	}
}


void flip(Node *tr, int idx, int l,int r){
	if(tr[idx].l > r || tr[idx].r<l)
		return;

	if(l<=tr[idx].l && tr[idx].r<=r){
		tr[idx].flipped ^=1;

		tr[idx].my_sum=tr[idx].tot_sum + mod - tr[idx].my_sum;

		if(tr[idx].my_sum >= mod)
			tr[idx].my_sum-=mod;

		int x=(tr[idx].r - tr[idx].l +1)&1;
		tr[idx].parity^=x;

		return;
	}

	if(p[idx].flipped){
		//unflip ie.push changes
		p[idx].flipped=false;

		tr[idx].my_sum=tr[idx].tot_sum + mod - tr[idx].my_sum;

		if(tr[idx].my_sum >= mod)
			tr[idx].my_sum-=mod;

		int x=(tr[idx].r - tr[idx].l +1)&1;
		tr[idx].parity^=x;

		if(l>tr[idx].l){
			flip(tr, 2*idx, tr[idx].l,l-1);
			flip(tr, 2*idx+1, tr[idx].l,l-1);
		}

		if(r<tr[idx].r){
			flip(tr, 2*idx, r-1,tr[idx].r);
			flip(tr, 2*idx+1, r-1,tr[idx].r);
		}
	}else{
		flip(tr,2*idx,l,r);
		flip(tr,2*idx+1,l,r);
	}

	tr[idx].my_sum = (tr[2*idx].my_sum + tr[2*idx+1].my_sum)%mod;
	tr[idx].parity = (tr[2*idx].parity + tr[2*idx+1].parity) & 1;
}

int get_val(ll sum, int parity, int eo){
	//we want to return sum/2
	ll q=(mod+1)/2;
	if(parity==1 && eo==0)
		return (sum*q)%mod;

	sum++;

	if(eo == 0)
		sum++;

	return (sum*q)%mod;
}

int main(){
	char buf[maxn];
	cin>>buf;
	strig s(buf);
	int n=s.size();
	int v[maxn];
	for (int i = 0; i < ; ++i)
		v[i]=s[n-i-1]-'0';

	Node* tr;
	init(tr, v, n);

	int m;
	cin>>m;

	int eo=v[0] & 1;
	cout<<get_val(tr[])<<"\n";

	for (int i = 0; i < m; ++i)
	{
		/* code */
	}
}

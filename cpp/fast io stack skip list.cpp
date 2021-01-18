#include <stack>
#include <vector>
#include <cstdio>
#include <algorithm>
#include <iostream>	

using namespace std;

#define MAX_N 1000005
#define ll long long
#define DEBUG 0
#define gc getchar_unlocked

int read_int() {
  char c = gc();
  while(c<'0' || c>'9') c = gc();
  int ret = 0;
  while(c>='0' && c<='9') {
    ret = 10 * ret + c - 48;
    c = gc();
  }
  return ret;
}

char read_char(){
	char c= gc();
	while(c!='C' && c!='D' && c!='<' && c!='>'&& c!='=') c=gc();
	return c;
}

struct Node
{
	int idx;
	int val;
	int l;
	int r;
};

bool nodeCmp(Node a,Node b){
	return a.val<b.val;
}

struct qNode
{
	int num;
	int eq,ls,gr;
};

bool qNodeCmp(qNode a, qNode b){
	return a.num<b.num;
}

struct iNode
{
	char op,pl;
	int num,idx;
} inp[MAX_N];

bool iNodeCmp(iNode a, iNode b){
	return a.num<b.num;
}

Node a[MAX_N];
stack<int> s;
vector<qNode> qr;
char out[MAX_N];


int main(){
	int n,m;
	n=read_int();
	m=read_int();

	for (int i = 0; i < n; ++i)
	{
		a[i].val=read_int();
		a[i].idx=i;
		while((!s.empty()) && (a[s.top()].val<=a[i].val)){
			int t=s.top();
			a[t].r=i;
			s.pop();
		}

		if(!s.empty()){
			int t=s.top();
			a[i].l=t;
		}else
			a[i].l=-1;
		s.push(i);
	}

	//empty stack
	while(!s.empty()){
		a[s.top()].r=n;
		s.pop();
	}

	sort(a,a+n,nodeCmp);

	ll oc;
	qNode tmp;
	//fill equal map
	for (int i = 0; i < n; ++i)
	{
		if(DEBUG) cout<<a[i].val<<":"<<a[i].l<<":"<<a[i].r<<"\n";
		oc=(a[i].r-a[i].idx)*((ll)(a[i].idx-a[i].l));
		oc%=2;
		if((!qr.empty()) && (qr.back().num==a[i].val))
			qr.back().eq=(qr.back().eq+oc)%2;
		else{
			tmp.num=a[i].val;
			tmp.eq=oc;
			qr.push_back(tmp);
		}
	}
	
	if(DEBUG) cout<<"size"<<qr.size();
	if(DEBUG)
		for (int i = 0; i < qr.size(); ++i)
			cout<<qr[i].num<<":"<<qr[i].eq<<"\n";

	//fill less map
	int less=0;
	for (int i=0;i<qr.size();i++)
	{
		less= (less + qr[i].eq)%2;
		qr[i].ls=less;
	}
	

	int big=0;
	for (int i = qr.size()-1; i >=0; --i)
	{
		big= (big + qr[i].eq)%2;
		qr[i].gr=big;
	}

	//---------------------------------------------------------------------------------------
	for (int i = 0; i < m; ++i)
	{
		inp[i].idx=i;
		inp[i].op=read_char();
		inp[i].num=read_int();
		inp[i].pl=read_char();
	}

	sort(inp, inp+m, iNodeCmp);

	int ans;
	ll tSub=n;
	tSub=(tSub*(tSub+1))/2;
	tSub%=2;
	vector<qNode>::iterator it= qr.begin();
	for (int i = 0; i < m; ++i)
	{
		while((it!=qr.end()) && (it->num < inp[i].num))
			++it;

		if(it!=qr.end()){
			if(inp[i].op=='='){
				if(inp[i].num == it->num)
					ans=it->eq;
				else
					ans=0;
			}else if(inp[i].op=='<'){
				ans=it->ls;
				ans=abs(ans-it->eq)%2;
			}else{
				ans=it->gr;
				if(inp[i].num == it->num)
					ans=abs(ans-it->eq)%2;
			}
		}else{
			ans=0;
			if(inp[i].op=='<')
				ans=tSub;
		}

		if(ans==0){
			out[inp[i].idx]=inp[i].pl=='C'?'D':'C';
		}else
			out[inp[i].idx]=inp[i].pl;
	}

	out[m]='\0';
	printf("%s\n", out);
}

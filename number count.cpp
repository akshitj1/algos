// problem find count of numbers from 1 to N which does not has digit d for each in 0 to 9
//algo: 3215 -> 
//1-999
//1000-3000
//3000-3200
//3200-3210
//3210-3215

#include <iostream>
#include <algorithm>

using namespace std;

typedef long long ll;

ll pow9[20];
ll num[11];

void process(){
	ll _tmp=1;
	for (int i = 0; i < 20; ++i,_tmp*=9)
		pow9[i]=_tmp;
}

void clr(){
	for (int i = 0; i <= 9; ++i)
		num[i]=0;
}

int solve(ll val){
	clr();
	ll _val=val;
	int cnum[20];
	int pos=1;
	while(_val){
		cnum[pos]=_val%10;
		_val/=10;
		pos++;
	}
	int N=pos-1;

	//from 1 to 9999..(N-1) times
	num[0]+=(pow9[N]-9)/8;
	for (int d = 1; d <= 9; ++d)
		num[d]+=pow9[N-1]-1;

	bool prevD[10]={};
	for (int p = N; p >= 1; --p)
	{
		int cd=cnum[p];
		if(cd==0){//no meaning in counting from 3200-3200
			prevD[cd]=true;
			continue;
		}
		for (int d = 0; d <= 9; ++d)
		{
			if(prevD[d])
				continue;
			if(d!=cd && !(p!=1 && d==0))//add current number itself ie. r in l-r
				num[d]++;

			if(p==N){
				if(d==0 || cd<=d)
					num[d]+=(cd-1)*pow9[p-1];
				else
					num[d]+=(cd-2)*pow9[p-1];
			}else{
				if(d<cd)
					num[d]+=(cd-1)*pow9[p-1];
				else
					num[d]+=cd*pow9[p-1];

				if(d!=0)
					num[d]-=1;//already counted trailing zeroes number previously ie. l in l-r
			}
		}
		prevD[cd]=true;
	}
}

int main(){
	int t;
	cin>>t;
	process();
	while(t--){
		ll val;
		cin>>val;
		solve(val);
	}
}
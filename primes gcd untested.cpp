#include <iostream>
#include <cstring>
#include <climits>
#include <cmath>
#include <cstdlib>
#include <vector>

using namespace std;

#define MAX_A 1000000000
#define pb push_back
#define DEBUG 1
#define ID if(DEBUG)

vector<int> p;

void findPrimes(){
	int maxL=sqrt(MAX_A);
	bool isPrime;
	for (int i = 2; i <= maxL; ++i)
	{	isPrime=true;
		int sqi=sqrt(i);
		for (int pi = 0; pi < p.size() && p[pi]<=sqi ; ++pi)
		{
			if(i%p[pi]==0){
				isPrime=false;
				break;
			}
		}
		if(isPrime)
			p.pb(i);
	}
}

int gcd(int a,int b){
	if(a<b)
		return gcd(b,a);
	if(b==0)
		return a;
	return gcd(b,a%b);
}

int factorize(int a){
	int ans=1;
	int sqa=sqrt(a);
	for (int i = 0; i < p.size() && p[i]<=sqrt(a); ++i)
	{
		int occ=0;
		while(a>0 && a%p[i]==0){
			a/=p[i];
			occ++;
		}
		ans*=(occ+1);
	}
	if(a>1)
		ans*=2;

	return ans;
}


int main(){
	int t;
	int n,m;
	cin>>t;
	while(t--){
		cin>>n>>m;
		int d=gcd(n-1,m-1);
		ID cout<<d<<"\n";
		cout<<factorize(d)+1<<"\n";
	}
}
#include <iostream>

using namespace std;
//const int mod=1299709;

//usage: inverse(num,mod,mod)
int inverse(int new_r, int r, int n, int new_t=1,int t=0){
	if(!new_r)
		return (t+n)%n;
	return inverse(r%new_r, new_r, n , t-(r/new_r)*new_t, new_t);
}

int main(){
	int n=2, mod=1000*1000*1000+7;
	int inv=inverse(n, mod, mod);
	cout<<inv<<" "<<(inv*n)%mod<<endl;
	int a=1341,b=1242;
	cout<<((a%mod * b%mod)*inv)%mod<<endl;
	cout<<((a*b)/2)%mod<<endl;
}

#include <iostream>
#include <vector>
#include <map>
#include <cstdlib>
#include <cstring>
#include <stack>
#include <queue>
#include <algorithm>
#include <string>
#include <sstream>

using namespace std;

#define pb push_back
#define clr(a) memset(a, 0, sizeof(a))
#define forl(i,l,u) for(int i=(l);i<(u);i++)
#define forn(i,n) forl(i,0,n)

typedef pair<int, int> ii;
typedef long long ll;

void permutation(string &s){
	cout<<s<<endl;
	while(true){
		int l=-1;
		for(int i=s.size()-2; i>=0; i--)
			if(s[i]<s[i+1]){
				l=i;
				break;
			}
		if(l==-1)
			break;
		int r=l+1;
		while(r<s.size() && s[l]<s[r])
			r++;
		swap(s[l], s[r-1]);
		reverse(s.begin()+l+1, s.end());
		cout<<s<<endl;
	}
}

int main(){
	ios_base::sync_with_stdio(false); cin.clear();
	string s="abc";
	permutation(s);
}
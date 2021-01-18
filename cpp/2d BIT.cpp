
int bit[maxn][maxn]={};

void updatey(int x,int y, int val){
	while(y<=n){
		bit[x][y]+=val;
		y+=(y & -y);
	}
}

void update(int x, int y, int val){
	while(x<=n){
		updatey(x,y,val);
		x+=(x & -x);
	}
}


int queryy(int x, int y){
	int ans=0;
	while(y>0){
		ans+=bit[x][y];
		y-=(y & -y);
	}
	return ans;
}

int query(int x, int y){
	int ans=0;
	while(x>0){
		ans+=queryy(x,y);
		x-=(x & -x);
	}
	return ans;
}

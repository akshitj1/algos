/*
N is no. of points
coeffs are for: coeff[0] x^3 + coeff[1] x^2 + coeff[2] x^1 + coeff[3]
*/
void gauss_jordan(){
	for (int r = 0; r < N; ++r)
	{
		int maxr=-1;
		for (int i = r; i < N; ++i)
			if(!isEqual(A[i][r],0) && (maxr==-1 || A[i][r] > A[maxr][r]))
				maxr=i;
		
		for (int j = 0; j < N+1; ++j)
			swap(A[r][j], A[maxr][j]);

		ld dv=A[r][r];
		for (int j = r; j < N+1; ++j)
			A[r][j]/=dv;

		for (int i = 0; i < N; ++i)
			if((!isEqual(A[i][r],0)) && i!=r){
				ld mul=A[i][r];
				for (int j = r; j < N+1; ++j)
					A[i][j]-= A[r][j] * mul;
			}
	}

	for (int i = 0; i < N; ++i){
		if(isEqual(A[i][N], 0))
			coeff[i]=0;
		else
			coeff[i]=A[i][N];

		cerr<<coeff[i]<<"x^"<<N-i-1<<" ";
	}
	cerr<<"\n";
}

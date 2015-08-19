#define MAX_V 50
#define MAX_D MAX_V


//for undirected
void FloydWarshal(int numV,char cGraph[MAX_V][MAX_V]){
	int dist[MAX_V][MAX_V];
	for (int i = 0; i < numV; ++i)
	{
		for (int j = i; j < numV; ++j)
		{
			if(i==j)
				dist[i][j]=0;
			else if(cGraph[i][j]=='Y')//if connected
				dist[j][i]=dist[i][j]=1;
			else
				dist[j][i]=dist[i][j]=MAX_D;
		}
	}

	for (int k = 0; k < numV; ++k){
		for (int i = 0; i < numV; ++i){
			for (int j = i+1; j < numV; ++j){//for directed graph
				
				if(dist[i][j] > (dist[i][k]+dist[k][j]))
					dist[j][i]=dist[i][j] = dist[i][k]+dist[k][j];
			}
		}
	}
}

//for directed
void FloydWarshal(int numV,char cGraph[MAX_V][MAX_V]){
	int dist[MAX_V][MAX_V];
	for (int i = 0; i < numV; ++i)
	{
		for (int j = 0; j < numV; ++j)
		{
			if(i==j)
				dist[i][j]=0;
			else if(cGraph[i][j]=='Y')//if connected
				dist[i][j]=1;
			else
				dist[i][j]=MAX_D;
		}
	}

	for (int k = 0; k < numV; ++k){
		for (int i = 0; i < numV; ++i){
			for (int j = 0; j < numV; ++j){//for directed graph
				
				if(dist[i][j] > (dist[i][k]+dist[k][j]))
					dist[i][j] = dist[i][k]+dist[k][j];
			}
		}
	}
}


void statistics(dr,n){
        int i,j,sum=0,sort_num=0;

        int flag=0;
        char drugsort[N];
        char drugs[N][N];
        drugsort=dr[0].sort;
        for(i=0;i<n;i++){
        		for(j=0;j<n;j++){
        				if(strcmp(dr[i].sort,drugs[j])==0){
        					break
        				}else{
        					drugs[i]=dr[i].sort;
        					sort_num++;
        				}
        		}
        }
        for(j=0;j<sort_num;j++){
        	for(i=0;i<n;i++){
    			    if(drugs[j]!=dr[i].sort){
                        	sum++;
                	}	
        	}
	      	printf("sort : %s have %d",drugs[j],sum);

        } 
} 

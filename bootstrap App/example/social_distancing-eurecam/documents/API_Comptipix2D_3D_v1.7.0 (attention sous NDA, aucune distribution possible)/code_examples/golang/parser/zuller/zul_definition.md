#Principle

Generation of a file with same format as *easy_comptage*  
It is quite close to **VCS timeless format**, except for aggregation time set to 30 min

##File name

*MMDDHHMM.ZUL*

**10091400.ZUL**  
1. 10	*month*  
2. 09	*day*  
3. 14	*hour*  
4. 00 *minuts*  

Extension is **ZUL**

##Content

Agregation time is **30 min**  
1. Field 1		*door number*  
2. Field 2		*date DD/MM/YYYY*  
3. Field 3 		*time HH:MM*  
4. Field 4		*entrances 4 digits*  
5. Field 5		*exits 4 digits*  
6. Field 6		*error code 1*  

###Example 1

*One door*  
	1	09/10/2013	14:00	0017	0012	1  

###Example 2

*3 doors*  
	1	09/10/2013	14:00	0001	0002	1  
	2	09/10/2013	14:00	0011	0012	1  
	3	09/10/2013	14:00	0021	0022	1  
	1	09/10/2013	14:30	0003	0004	1  
	2	09/10/2013	14:30	0013	0014	1  
	3	09/10/2013	14:30	0023	0024	1  
	1	09/10/2013	15:00	0005	0006	1  
	2	09/10/2013	15:00	0015	0016	1  
	3	09/10/2013	15:00	0025	0026	1  
	1	09/10/2013	15:30	0007	0008	1  
	2	09/10/2013	15:30	0017	0018	1  
	3	09/10/2013	15:30	0027	0028	1  

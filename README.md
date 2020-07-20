# cryptodini

#### Note:
this project is suppose to show only implementation logic without building pipeline (Jenkinfile, dockerfile, makefile) please let's me know if it needed.
 

### First part for "Cryptodini Server"
It represent the implementation of main SDK function 
- "Adjust" which is call by Each Robo Server when ports is need to be adjusted.
- "GetPort", it should be logic to be called before Robo Server logic determine the adjust needed.

### Second part for example of "Robo Server"
- "Deposit", it should return to Cryptodini which assets it should buy.
- "Withdraw", it should return to Cryptodini which assets it should sell.
- "GetPort", return current port or the specific user.

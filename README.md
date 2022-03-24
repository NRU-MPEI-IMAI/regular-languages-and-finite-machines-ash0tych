# Theorethical models of compututition

## Homework №1

### Task 1. Create a FA that recognizes a given language. Automat must be deterministic

##### Automat №1

![L = \{w \in \{a,b,c \}^* | |w|_c = 1 \}](https://latex.codecogs.com/png.image?%5Cdpi%7B110%7DL%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b,c%20%5C%7D%5E*%20%7C%20%7Cw%7C_c%20=%201%20%5C%7D)

   <img title="" src="file:///D:/6 сем/TMOC/HW 1/task 1/graph_1_1.svg" alt="" data-align="center" width="717">    

##### Automat №2

![L = \{w \in \{a,b\}^* | |w|_a \leq  2; |w|_b \geq 2\}](https://latex.codecogs.com/png.image?%5Cdpi%7B110%7DL%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%5Cleq%20%202;%20%7Cw%7C_b%20%5Cgeq%202%5C%7D)

<img title="" src="file:///D:/6 сем/TMOC/HW 1/task 1/graph_1_2.svg" alt="" data-align="center" width="735">

##### Automat №3

![L = \{w \in \{a,b\}^* | |w|_a  \neq |w|_b \}](https://latex.codecogs.com/png.image?%5Cdpi%7B110%7DL%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%20%5Cneq%20%7Cw%7C_b%20%5C%7D)

<img title="" src="file:///D:/6 сем/TMOC/HW 1/task 1/graph_1_3.svg" alt="" data-align="center" width="743">

##### Automat №4

![L = \{w \in \{a,b\}^* | ww = www \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20ww%20=%20www%20%5C%7D)

<img title="" src="file:///D:/6 сем/TMOC/HW 1/task 1/graph_1_4.svg" alt="" data-align="center" width="687">

### Task 2. Create a FA that recognizes a given language. Automat must be constructed by direct production of DFA

##### Automat №1

$L = \{w \in \{a,b\}^* ||w|_a \geq 2 \wedge |w|_b \geq 2\}$

   First, let's create DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_a \geq 2\}$

<img title="" src="file:///D:/6 сем/TMOC/HW 1/task 2/1/graph_2_1_1.svg" alt="" data-align="center" width="827">

   Now, let's create second DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_b \geq 2\}$

<img title="" src="file:///D:/6 сем/TMOC/HW 1/task 2/1/graph_2_1_2.svg" alt="" data-align="center" width="784">

   And our final result:   

![](D:\6%20сем\TMOC\HW%201\task%202\1\graph_2_1_3.svg)

##### Automat №2

L = $\{w \in \{a,b\}^* ||w|_a \geq 3 \wedge |w|_b \;is \; odd\; \}$

   First, let's create DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_a \geq 3\}$

<img title="" src="file:///D:/6 сем/TMOC/HW 1/task 2/2/graph_2_2_1.svg" alt="" data-align="center" width="791">

   Now, let's create second DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_b \;is \; odd;\}$

<img title="" src="file:///D:/6%20сем/TMOC/HW%201/task%202/2/graph_2_2_2.svg" alt="" data-align="center" width="652">

   And our final result:

<img title="" src="file:///D:/6%20сем/TMOC/HW%201/task%202/2/graph_2_2_3.svg" alt="" data-align="center" width="735">

##### Automat №3

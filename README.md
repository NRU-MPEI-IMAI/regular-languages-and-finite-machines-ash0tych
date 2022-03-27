# Theorethical models of compututition

## Homework №1

### Task 1. Create a FA that recognizes a given language. Automat must be deterministic

##### Automat №1

![L = \{w \in \{a,b,c \}^* | |w|_c = 1 \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b,c%20%5C%7D%5E*%20%7C%20%7Cw%7C_c%20=%201%20%5C%7D)

   <img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%201/graph_1_1.svg" alt="" data-align="center" width="717">    

##### Automat №2

![L = \{w \in \{a,b\}^* | |w|_a \leq  2; |w|_b \geq 2\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%5Cleq%20%202;%20%7Cw%7C_b%20%5Cgeq%202%5C%7D)

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%201/graph_1_2.svg" alt="" data-align="center" width="735">

##### Automat №3

![L = \{w \in \{a,b\}^* | |w|_a  \neq |w|_b \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%20%5Cneq%20%7Cw%7C_b%20%5C%7D)

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%201/graph_1_3.svg" alt="" data-align="center" width="743">

##### Automat №4

![L = \{w \in \{a,b\}^* | ww = www \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20ww%20=%20www%20%5C%7D)

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%201/graph_1_4.svg" alt="" data-align="center" width="687">

### Task 2. Create a FA that recognizes a given language. Automat must be constructed by direct production of DFA

##### Automat №1

![L = \{w \in \{a,b\}^* ||w|_a \geq 2 \wedge |w|_b \geq 2\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%202%20%5Cwedge%20%7Cw%7C_b%20%5Cgeq%202%5C%7D)

   First, let's create DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_a \geq 2\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%202%5C%7D)

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/1/graph_2_1_1.svg" alt="" data-align="center" width="827">

   Now, let's create second DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_b \geq 2\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_b%20%5Cgeq%202%5C%7D)

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/1/graph_2_1_2.svg" alt="" data-align="center" width="784">

   And our final result:   

![](https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/1/graph_2_1_3.svg)

##### Automat №2

![L = \{w \in \{a,b\}^* ||w|_a \geq 3 \wedge |w|_b \;is \; odd\; \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%203%20%5Cwedge%20%7Cw%7C_b%20%5C;is%20%5C;%20odd%5C;%20%5C%7D)

   First, let's create DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_a \geq 3\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%203%5C%7D)

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/2/graph_2_2_1.svg" alt="" data-align="center" width="791">

   Now, let's create second DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_b \;is \; odd;\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_b%20%5C;is%20%5C;%20odd;%5C%7D)

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/2/graph_2_2_2.svg" alt="" data-align="center" width="652">

   And our final result:

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/2/graph_2_2_3.svg" alt="" data-align="center" width="735">

##### Automat №3

![L = \{w \in \{a,b\}^* | |w|_a \;is \; even\; \wedge |w|_b \; is \; multiple \; of \; 3\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%5C;is%20%5C;%20even%5C;%20%5Cwedge%20%7Cw%7C_b%20%5C;%20is%20%5C;%20multiple%20%5C;%20of%20%5C;%203%5C%7D)

First, let's create DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_a \;is \; even\;\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5C;is%20%5C;%20even%5C;%5C%7D)

<img src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/3/graph_2_3_1.svg" title="" alt="" width="657">

Now, let's create second DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_b \; is \; multiple \; of \; 3\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_b%20%5C;%20is%20%5C;%20multiple%20%5C;%20of%20%5C;%203%5C%7D)

<img src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/3/graph_2_3_2.svg" title="" alt="" width="751">

And our final result:

<img title="" src="https://github.com/NRU-MPEI-IMAI/regular-languages-and-finite-machines-ash0tych/blob/main/task%202/3/graph_2_3_3.svg" alt="" width="771">

##### Automat №4

$L = \overline{L_3}$

Well, this is easy

<img src="file:///D:/6%20сем/TMOC/HW%201/task%202/4/graph_2_4.svg" title="" alt="" width="726">

##### Automat №5

$L = L_2 / L_3$

So, $L_2$ has 8 nodes and $L_3$ has 6 nodes. This mean's $L_5$ will have 48 nodes.... Screw this


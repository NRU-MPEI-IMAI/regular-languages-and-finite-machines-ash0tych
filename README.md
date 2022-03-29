# Theorethical models of compututition

## Homework №1

### Task 1. Create a FA that recognizes a given language. Automat must be deterministic

##### Automat №1

![L = \{w \in \{a,b,c \}^* | |w|_c = 1 \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b,c%20%5C%7D%5E*%20%7C%20%7Cw%7C_c%20=%201%20%5C%7D)

   <img title="" src="./task 1/graph_1_1.svg" alt="" data-align="inline" width="746">    

##### Automat №2

![L = \{w \in \{a,b\}^* | |w|_a \leq  2; |w|_b \geq 2\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%5Cleq%20%202;%20%7Cw%7C_b%20%5Cgeq%202%5C%7D)

<img title="" src="./task 1/graph_1_2.svg" alt="" data-align="inline" width="779">

##### Automat №3

![L = \{w \in \{a,b\}^* | |w|_a  \neq |w|_b \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%20%5Cneq%20%7Cw%7C_b%20%5C%7D)

<img title="" src="./task 1/graph_1_3.svg" alt="" data-align="inline" width="737">

##### Automat №4

![L = \{w \in \{a,b\}^* | ww = www \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20ww%20=%20www%20%5C%7D)

<img title="" src="./task 1/graph_1_4.svg" alt="" data-align="inline" width="710">

### Task 2. Create a FA that recognizes a given language. Automat must be constructed by direct production of DFA

##### Automat №1

![L = \{w \in \{a,b\}^* ||w|_a \geq 2 \wedge |w|_b \geq 2\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%202%20%5Cwedge%20%7Cw%7C_b%20%5Cgeq%202%5C%7D)

   First, let's create DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_a \geq 2\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%202%5C%7D)

<img title="" src="./task 2/1/graph_2_1_1.svg" alt="" data-align="inline" width="827">

   Now, let's create second DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_b \geq 2\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_b%20%5Cgeq%202%5C%7D)

<img title="" src="./task 2/1/graph_2_1_2.svg" alt="" data-align="inline" width="784">

   And our final result:   

<img src="./task 2/1/graph_2_1_3.svg" title="" alt="" data-align="inline">

##### Automat №2

![L = \{w \in \{a,b\}^* ||w|_a \geq 3 \wedge |w|_b \;is \; odd\; \}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%203%20%5Cwedge%20%7Cw%7C_b%20%5C;is%20%5C;%20odd%5C;%20%5C%7D)

   First, let's create DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_a \geq 3\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5Cgeq%203%5C%7D)

<img title="" src="./task 2/2/graph_2_2_1.svg" alt="" data-align="inline" width="815">

   Now, let's create second DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_b \;is \; odd\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_b%20%5C;is%20%5C;%20odd%5C%7D)

<img title="" src="./task 2/2/graph_2_2_2.svg" alt="" data-align="inline" width="773">

   And our final result:

<img title="" src="./task 2/2/graph_2_2_3.svg" alt="" data-align="inline" width="771">

##### Automat №3

![L = \{w \in \{a,b\}^* | |w|_a \;is \; even\; \wedge |w|_b \; is \; multiple \; of \; 3\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%20%7Cw%7C_a%20%5C;is%20%5C;%20even%5C;%20%5Cwedge%20%7Cw%7C_b%20%5C;%20is%20%5C;%20multiple%20%5C;%20of%20%5C;%203%5C%7D)

First, let's create DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_a \;is \; even\;\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_a%20%5C;is%20%5C;%20even%5C;%5C%7D)

<img src="./task 2/3/graph_2_3_1.svg" title="" alt="" width="657">

Now, let's create second DFA that recognizes ![L = \{w \in \{a,b\}^* ||w|_b \; is \; multiple \; of \; 3\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bw%20%5Cin%20%5C%7Ba,b%5C%7D%5E*%20%7C%7Cw%7C_b%20%5C;%20is%20%5C;%20multiple%20%5C;%20of%20%5C;%203%5C%7D)

<img src="./task 2/3/graph_2_3_2.svg" title="" alt="" width="751">

And our final result:

<img title="" src="./task 2/3/graph_2_3_3.svg" alt="" width="771">

##### Automat №4

![L = \overline{L_3}](https://latex.codecogs.com/svg.image?L%20=%20%5Coverline%7BL_3%7D)

Well, this is easy

<img src="./task 2/4/graph_2_4.svg" title="" alt="" width="726">

##### Automat №5

![L = L_2 / L_3](https://latex.codecogs.com/svg.image?L%20=%20L_2%20/%20L_3)

So, ![L_2](https://latex.codecogs.com/svg.image?L_2) has 8 nodes and ![L_3](https://latex.codecogs.com/svg.image?L_3) has 6 nodes. This mean's ![L_5](https://latex.codecogs.com/svg.image?L_5) will have 48 nodes.... Screw this

### Task 3. Create a minimal DFA that recognizes the same language as described by a regular expression.

##### Regular expression №1

![(ab+aba)^* a](https://latex.codecogs.com/svg.image?(ab&plus;aba)%5E*%20a)

Btw, since now i will only compile .dot in Golang, because writing there something more than 6 nodes is realy hard.

Cuz we have "+" first of all we must define NFA.

<img title="" src="task 3/rendered/graph_3_1_1.svg" alt="" data-align="inline" width="697">

|          | a             | b             |
|:--------:|:-------------:|:-------------:|
| 1        | 3,6,10        | ![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) |
| 3,6,10   | ![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) | 4,7           |
| 4,7      | 8,3,6,10      | ![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) |
| 8,3,6,10 | 3,6,10        | 4,7           |

Now, using this table, we can convert NFA to DFA

<img title="" src="task 3/rendered/graph_3_1_2.svg" alt="" data-align="inline" width="683">

##### Regular expression №2

![ a(a(ab)^∗b )^∗ (ab)^∗ ](https://latex.codecogs.com/svg.image?a(a(ab)%5E*b)%5E*(ab)%5E*)

NFA:

<img title="" src="./task 3/rendered/graph_3_2_1.svg" alt="" data-align="inline" width="748">

Table to conversion:

|     | a             | b             |
| --- | ------------- | ------------- |
| 1   | 2,8           | ![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) |
| 2,8 | 3,7           | ![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) |
| 3,7 | 4             | 6,8           |
| 4   |![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) | 5             |
| 6,8 | 3,7           | ![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) |
| 5   | 4             | 6             |
| 6   | 3,7           | ![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) |

And our automat:

<img title="" src="./task 3/rendered/graph_3_2_2.svg" alt="" data-align="inline" width="745">  

But there is a problem. This automat is not minimal and our task is to create **minimal**. So there it is

<img title="" src="./task 3/rendered/graph_3_2_3.svg" alt="" data-align="inline">

##### Regular expression №3

![(a + (a + b) (a + b)b)^* ](https://latex.codecogs.com/svg.image?(a%20&plus;%20(a%20&plus;%20b)%20(a%20&plus;%20b)b)%5E*%20)

NFA:

<img title="" src="./task 3/rendered/graph_3_3_1.svg" alt="" data-align="inline" width="804">

Table:

|       | a             | b     |
| ----- | ------------- | ----- |
| 1     | 1,2           | 2     |
| 1,2   | 1,2,3         | 2,3   |
| 2     | 3             | 3     |
| 1,2,3 | 1,2,3         | 1,2,3 |
| 2,3   | 3             | 1,3   |
| 3     | ![\varnothing](https://latex.codecogs.com/svg.image?%5Cvarnothing) | 1     |
| 1,3   | 1,2           | 1,2   |

DFA:

<img title="" src="./task 3/rendered/graph_3_3_2.svg" alt="" data-align="inline" width="766">

##### Regular expression №4

![(b + c) ((ab)^* c + (ba)^* )^* ](https://latex.codecogs.com/svg.image?(b%20&plus;%20c)%20((ab)%5E*%20c%20&plus;%20(ba)%5E*%20)%5E*%20)

We don't even need NFA for this case. Lets create DFA:

<img title="" src="./task 3/rendered/graph_3_4_2.svg" alt="" data-align="inline" width="709">

We can minimize this:

<img title="" src="./task 3/rendered/graph_3_4_3.svg" alt="" data-align="inline" width="688">

##### Regular expression №5

![(a + b)^+ (aa + bb + abab + baba) (a + b)^+ ](https://latex.codecogs.com/svg.image?(a%20&plus;%20b)%5E&plus;%20(aa%20&plus;%20bb%20&plus;%20abab%20&plus;%20baba)%20(a%20&plus;%20b)%5E&plus;%20) .... No

### Task 4. Determine whether the language is regular or not

##### Language 1

![L = {(aab)^nb(aba)^m \mid n \geqslant  0, m \geqslant  0}](https://latex.codecogs.com/svg.image?L%20=%20%7B(aab)%5Enb(aba)%5Em%20%5Cmid%20n%20%5Cgeqslant%20%200,%20m%20%5Cgeqslant%20%200%7D)

<img title="" src="./task 4/graph_4_1.svg" alt="" data-align="inline" width="695">

Yes it's regular

##### Language 2

![L = \{uaav \mid u \in \{a, b\}^*, v \in \{a, b\}^*, |u|_b \geqslant  |v|_a\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Buaav%20%5Cmid%20u%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%20v%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%20%7Cu%7C_b%20%5Cgeqslant%20%20%7Cv%7C_a%5C%7D)

![
\\
\omega = b^naaa^n, |\omega| \geq  n \\
\omega = xyz\\
x = b^i \quad y=b^j \quad i + j \leq n \quad j > 0 \\
z = b^{n-i-j}aaa^n \\
|xy| \leq n \quad |y| > 0 \\
xy^0z = b^ib^{n-i-j}aaa^n = b^{n-j}aaa^n\notin L
](https://latex.codecogs.com/svg.image?%5C%5C%5Comega%20=%20b%5Enaaa%5En,%20%7C%5Comega%7C%20%5Cgeq%20%20n%20%5C%5C%5Comega%20=%20xyz%5C%5Cx%20=%20b%5Ei%20%5Cquad%20y=b%5Ej%20%5Cquad%20i%20&plus;%20j%20%5Cleq%20n%20%5Cquad%20j%20%3E%200%20%5C%5Cz%20=%20b%5E%7Bn-i-j%7Daaa%5En%20%5C%5C%7Cxy%7C%20%5Cleq%20n%20%5Cquad%20%7Cy%7C%20%3E%200%20%5C%5Cxy%5E0z%20=%20b%5Eib%5E%7Bn-i-j%7Daaa%5En%20=%20b%5E%7Bn-j%7Daaa%5En%5Cnotin%20L

##### Language 3

![L = \{a^mw \mid w \in \{a, b\}^*, 1 \leqslant  |w|_b \leqslant  m\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Ba%5Emw%20%5Cmid%20w%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%201%20%5Cleqslant%20%20%7Cw%7C_b%20%5Cleqslant%20%20m%5C%7D)

![
\\
\omega = a^nb^n , |\omega| \geq  n \\
\omega = xyz \\
x = a^i \quad y = a^j \quad i+j \leq n \quad j > 0 \\
z = a^{n-i-j}b^n \\
|xy| \leq n \quad |y| > 0 \\
xy^0z = a^ia^{n-i-j}b^n = a^{n-j}b^n \notin L
](https://latex.codecogs.com/svg.image?%5C%5C%5Comega%20=%20a%5Enb%5En%20,%20%7C%5Comega%7C%20%5Cgeq%20%20n%20%5C%5C%5Comega%20=%20xyz%20%5C%5Cx%20=%20a%5Ei%20%5Cquad%20y%20=%20a%5Ej%20%5Cquad%20i&plus;j%20%5Cleq%20n%20%5Cquad%20j%20%3E%200%20%5C%5Cz%20=%20a%5E%7Bn-i-j%7Db%5En%20%5C%5C%7Cxy%7C%20%5Cleq%20n%20%5Cquad%20%7Cy%7C%20%3E%200%20%5C%5Cxy%5E0z%20=%20a%5Eia%5E%7Bn-i-j%7Db%5En%20=%20a%5E%7Bn-j%7Db%5En%20%5Cnotin%20L)

##### Language 4

![L = \{a^kb^ma^n \mid k = n \vee m > 0\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Ba%5Ekb%5Ema%5En%20%5Cmid%20k%20=%20n%20%5Cvee%20%20m%20%3E%200%5C%7D)

![
\\
\omega = a^nba^n , |\omega| \geq  n \\
\omega = xyz \\
x = a^i \quad y = a^j \quad i+j \leq n \quad j > 0 \\
z = a^{n-i-j}ba^n \\
|xy| \leq n \quad |y| > 0 \\
xy^kz = a^ia^{jk}a^{n-i-j}ba^n = a^{n-j(k-1)}ba^n \notin L \quad \forall k > 1
](https://latex.codecogs.com/svg.image?%5C%5C%5Comega%20=%20a%5Enba%5En%20,%20%7C%5Comega%7C%20%5Cgeq%20%20n%20%5C%5C%5Comega%20=%20xyz%20%5C%5Cx%20=%20a%5Ei%20%5Cquad%20y%20=%20a%5Ej%20%5Cquad%20i&plus;j%20%5Cleq%20n%20%5Cquad%20j%20%3E%200%20%5C%5Cz%20=%20a%5E%7Bn-i-j%7Dba%5En%20%5C%5C%7Cxy%7C%20%5Cleq%20n%20%5Cquad%20%7Cy%7C%20%3E%200%20%5C%5Cxy%5Ekz%20=%20a%5Eia%5E%7Bjk%7Da%5E%7Bn-i-j%7Dba%5En%20=%20a%5E%7Bn-j(k-1)%7Dba%5En%20%5Cnotin%20L%20%5Cquad%20%5Cforall%20k%20%3E%201)

##### Language 5

![L = \{ucv \mid u \in \{a, b\}^*, v \in \{a, b\}^*, u \neq v^R\}](https://latex.codecogs.com/svg.image?L%20=%20%5C%7Bucv%20%5Cmid%20u%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%20v%20%5Cin%20%5C%7Ba,%20b%5C%7D%5E*,%20u%20%5Cneq%20v%5ER%5C%7D)

![
\\
\omega = (ab)^nc(ba)^n = \alpha_1\alpha_2...\alpha_{4n+1}, |\omega| \geq n \\
\omega = xyz \\
x = \alpha_1\alpha_2...\alpha_i \quad y=\alpha_{i+1}\alpha_{i+2}...\alpha_{i+j} \quad i+j \leq n \quad j > 0\\
z = \alpha_{i+j+1}\alpha_{i+j+2}...\alpha_{2n}c(ba)^n \\
|xy| \leq n \quad |y| > 0 \\
xy^kz = \alpha_1...\alpha_i(\alpha_{i+1}...\alpha_{i+j})^k\alpha_{i+j+1}...\alpha_{2n}c(ba)^n \notin L \quad \forall k > 1\\
](https://latex.codecogs.com/svg.image?%5C%5C%5Comega%20=%20(ab)%5Enc(ba)%5En%20=%20%5Calpha_1%5Calpha_2...%5Calpha_%7B4n&plus;1%7D,%20%7C%5Comega%7C%20%5Cgeq%20n%20%5C%5C%5Comega%20=%20xyz%20%5C%5Cx%20=%20%5Calpha_1%5Calpha_2...%5Calpha_i%20%5Cquad%20y=%5Calpha_%7Bi&plus;1%7D%5Calpha_%7Bi&plus;2%7D...%5Calpha_%7Bi&plus;j%7D%20%5Cquad%20i&plus;j%20%5Cleq%20n%20%5Cquad%20j%20%3E%200%5C%5Cz%20=%20%5Calpha_%7Bi&plus;j&plus;1%7D%5Calpha_%7Bi&plus;j&plus;2%7D...%5Calpha_%7B2n%7Dc(ba)%5En%20%5C%5C%7Cxy%7C%20%5Cleq%20n%20%5Cquad%20%7Cy%7C%20%3E%200%20%5C%5Cxy%5Ekz%20=%20%5Calpha_1...%5Calpha_i(%5Calpha_%7Bi&plus;1%7D...%5Calpha_%7Bi&plus;j%7D)%5Ek%5Calpha_%7Bi&plus;j&plus;1%7D...%5Calpha_%7B2n%7Dc(ba)%5En%20%5Cnotin%20L%20%5Cquad%20%5Cforall%20k%20%3E%201%5C%5C)

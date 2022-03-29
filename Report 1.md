# Theorethical models of compututition

## Homework №1

### Task 1. Create a FA that recognizes a given language. Automat must be deterministic

##### Automat №1

$L = \{w \in \{a,b,c \}^* | |w|_c = 1 \}$

   <img title="" src="./task 1/graph_1_1.svg" alt="" data-align="inline" width="746">    

##### Automat №2

$L = \{w \in \{a,b\}^* | |w|_a \leq  2; |w|_b \geq 2\}$

<img title="" src="./task 1/graph_1_2.svg" alt="" data-align="inline" width="779">

##### Automat №3

$L = \{w \in \{a,b\}^* | |w|_a  \neq |w|_b \}$

<img title="" src="./task 1/graph_1_3.svg" alt="" data-align="inline" width="737">

##### Automat №4

$L = \{w \in \{a,b\}^* | ww = www \}$

<img title="" src="./task 1/graph_1_4.svg" alt="" data-align="inline" width="710">

### Task 2. Create a FA that recognizes a given language. Automat must be constructed by direct production of DFA

##### Automat №1

$L = \{w \in \{a,b\}^* ||w|_a \geq 2 \wedge |w|_b \geq 2\}$

   First, let's create DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_a \geq 2\}$

<img title="" src="./task 2/1/graph_2_1_1.svg" alt="" data-align="inline" width="827">

   Now, let's create second DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_b \geq 2\}$

<img title="" src="./task 2/1/graph_2_1_2.svg" alt="" data-align="inline" width="784">

   And our final result:   

<img src="./task 2/1/graph_2_1_3.svg" title="" alt="" data-align="inline">

##### Automat №2

$L = \{w \in \{a,b\}^* ||w|_a \geq 3 \wedge |w|_b \;is \; odd\; \}$

   First, let's create DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_a \geq 3\}$

<img title="" src="./task 2/2/graph_2_2_1.svg" alt="" data-align="inline" width="815">

   Now, let's create second DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_b \;is \; odd\}$

<img title="" src="./task 2/2/graph_2_2_2.svg" alt="" data-align="inline" width="773">

   And our final result:

<img title="" src="./task 2/2/graph_2_2_3.svg" alt="" data-align="inline" width="771">

##### Automat №3

$L = \{w \in \{a,b\}^* | |w|_a \;is \; even\; \wedge |w|_b \; is \; multiple \; of \; 3\}$

First, let's create DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_a \;is \; even\;\}$

<img src="./task 2/3/graph_2_3_1.svg" title="" alt="" width="657">

Now, let's create second DFA that recognizes $L = \{w \in \{a,b\}^* ||w|_b \; is \; multiple \; of \; 3\}$

<img src="./task 2/3/graph_2_3_2.svg" title="" alt="" width="751">

And our final result:

<img title="" src="./task 2/3/graph_2_3_3.svg" alt="" width="771">

##### Automat №4

$L = \overline{L_3}$

Well, this is easy

<img src="./task 2/4/graph_2_4.svg" title="" alt="" width="726">

##### Automat №5

$L = L_2 / L_3$

So, $L_2$ has 8 nodes and $L_3$ has 6 nodes. This mean's $L_5$ will have 48 nodes.... Screw this

### Task 3. Create a minimal DFA that recognizes the same language as described by a regular expression.

##### Regular expression №1

$(ab+aba)^* a$

Btw, since now i will only compile .dot in Golang, because writing there something more than 6 nodes is realy hard.

Cuz we have "+" first of all we must define NFA.

<img title="" src="task 3/rendered/graph_3_1_1.svg" alt="" data-align="inline" width="697">

|          | a             | b             |
|:--------:|:-------------:|:-------------:|
| 1        | 3,6,10        | $\varnothing$ |
| 3,6,10   | $\varnothing$ | 4,7           |
| 4,7      | 8,3,6,10      | $\varnothing$ |
| 8,3,6,10 | 3,6,10        | 4,7           |

Now, using this table, we can convert NFA to DFA

<img title="" src="task 3/rendered/graph_3_1_2.svg" alt="" data-align="inline" width="683">

##### Regular expression №2

$a(a(ab)^∗b )^∗ (ab)^∗$

NFA:

<img title="" src="./task 3/rendered/graph_3_2_1.svg" alt="" data-align="inline" width="748">

Table to conversion:

|     | a             | b             |
| --- | ------------- | ------------- |
| 1   | 2,8           | $\varnothing$ |
| 2,8 | 3,7           | $\varnothing$ |
| 3,7 | 4             | 6,8           |
| 4   | $\varnothing$ | 5             |
| 6,8 | 3,7           | $\varnothing$ |
| 5   | 4             | 6             |
| 6   | 3,7           | $\varnothing$ |

And our automat:

<img title="" src="./task 3/rendered/graph_3_2_2.svg" alt="" data-align="inline" width="745">  

But there is a problem. This automat is not minimal and our task is to create **minimal**. So there it is

<img title="" src="./task 3/rendered/graph_3_2_3.svg" alt="" data-align="inline">

##### Regular expression №3

$(a + (a + b) (a + b)b)^* $

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
| 3     | $\varnothing$ | 1     |
| 1,3   | 1,2           | 1,2   |

DFA:

<img title="" src="./task 3/rendered/graph_3_3_2.svg" alt="" data-align="inline" width="766">

##### Regular expression №4

$(b + c) ((ab)^∗ c + (ba)^∗ )^∗ $

We don't even need NFA for this case. Lets create DFA:

<img title="" src="./task 3/rendered/graph_3_4_2.svg" alt="" data-align="inline" width="709">

We can minimize this:

<img title="" src="./task 3/rendered/graph_3_4_3.svg" alt="" data-align="inline" width="688">

##### Regular expression №5

$(a + b)^+ (aa + bb + abab + baba) (a + b)^+ $ .... No

### Task 4. Determine whether the language is regular or not

##### Language 1

$L = {(aab)^nb(aba)^m \mid n ≥ 0, m ≥ 0}$

<img title="" src="./task 4/graph_4_1.svg" alt="" data-align="inline" width="695">

Yes it's regular

##### Language 2

$L = \{uaav \mid u ∈ \{a, b\}^∗, v ∈ \{a, b\}^∗, |u|_b ≥ |v|_a\}$

$$
\omega = b^naaa^n, |\omega| \geq  n \\
\omega = xyz\\
x = b^i \quad y=b^j \quad i + j \leq n \quad j > 0 \\
z = b^{n-i-j}aaa^n \\
|xy| \leq n \quad |y| > 0 \\
xy^0z = b^ib^{n-i-j}aaa^n = b^{n-j}aaa^n\notin L
$$

##### Language 3

$L = \{a^mw \mid w ∈ \{a, b\}^∗, 1 ≤ |w|_b ≤ m\}$

$$
\omega = a^nb^n , |\omega| \geq  n \\
\omega = xyz \\
x = a^i \quad y = a^j \quad i+j \leq n \quad j > 0 \\
z = a^{n-i-j}b^n \\
|xy| \leq n \quad |y| > 0 \\
xy^0z = a^ia^{n-i-j}b^n = a^{n-j}b^n \notin L
$$

##### Language 4

$L = \{a^kb^ma^n \mid k = n ∨ m > 0\}$

$$
\omega = a^nba^n , |\omega| \geq  n \\
\omega = xyz \\
x = a^i \quad y = a^j \quad i+j \leq n \quad j > 0 \\
z = a^{n-i-j}ba^n \\
|xy| \leq n \quad |y| > 0 \\
xy^kz = a^ia^{jk}a^{n-i-j}ba^n = a^{n-j(k-1)}ba^n \notin L \quad \forall k > 1
$$

##### Language 5

$L = \{ucv \mid u ∈ \{a, b\}^∗, v ∈ \{a, b\}^∗, u \neq v^R\}$

$$

\omega = (ab)^nc(ba)^n = \alpha_1\alpha_2...\alpha_{4n+1}, |\omega| \geq n \\
\omega = xyz \\
x = \alpha_1\alpha_2...\alpha_i \quad y=\alpha_{i+1}\alpha_{i+2}...\alpha_{i+j} \quad i+j \leq n \quad j > 0\\
z = \alpha_{i+j+1}\alpha_{i+j+2}...\alpha_{2n}c(ba)^n \\
|xy| \leq n \quad |y| > 0 \\
xy^kz = \alpha_1...\alpha_i(\alpha_{i+1}...\alpha_{i+j})^k\alpha_{i+j+1}...\alpha_{2n}c(ba)^n \notin L \quad \forall k > 1\\

$$

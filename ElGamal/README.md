# ElGamal暗号
離散対数問題の困難性に基づく公開鍵暗号系。

![\begin{align*}
c_2c_1^{p-1-x}&\equiv my^r\times g^{r(p-1-x)}\equiv mg^{rx}\times g^{r(p-1)-rx}\\
&\equiv mg^{r(p-1)}
\end{align*}
](https://render.githubusercontent.com/render/math?math=%5Cdisplaystyle+%5Cbegin%7Balign%2A%7D%0Ac_2c_1%5E%7Bp-1-x%7D%26%5Cequiv+my%5Er%5Ctimes+g%5E%7Br%28p-1-x%29%7D%5Cequiv+mg%5E%7Brx%7D%5Ctimes+g%5E%7Br%28p-1%29-rx%7D%5C%5C%0A%26%5Cequiv+mg%5E%7Br%28p-1%29%7D%0A%5Cend%7Balign%2A%7D%0A)


### 素数の選び方
+ pは離散対数問題が解けないくらい大きくなければいけない(pは1024bit以上)
+ p-1が小さな素数の積の場合、**Pohlig-Hellman**の方法で離散対数問題が解けてしまう

### 楕円ElGamal暗号
ElGamalECで実装（予定）。
楕円曲線上の離散対数問題に依存する。


### 拡張ElGamal暗号
ElGamalExで実装。CCAに対して耐性がある。
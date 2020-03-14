# Go 연산자

<table class="table">
<tbody><tr>
  <th>연산자</th>
  <th>기능</th>
  <th>설명</th>
</tr>
<tr>
  <td>=</td>
  <td>대입</td>
  <td>
변수나 상수에 값을 대입합니다. 변수는 변수끼리 대입할 수 있습니다.


<div class="highlight"><pre><code class="go"><span class="kd">var</span> <span class="nx">a</span> <span class="kt">int</span> <span class="p">=</span> <span class="mi">1</span>
<span class="kd">var</span> <span class="nx">b</span> <span class="kt">int</span> <span class="p">=</span> <span class="mi">2</span>
<span class="kd">var</span> <span class="nx">c</span> <span class="kt">int</span> <span class="p">=</span> <span class="nx">b</span>
<span class="kd">const</span> <span class="nx">d</span> <span class="kt">string</span> <span class="p">=</span> <span class="s">"Hello, world!"</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>:=</td>
  <td>변수 선언 및 대입</td>
  <td>변수를 선언하는 동시에 값을 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">1</span>               <span class="c1">// int</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mf">3.5</span>             <span class="c1">// float64</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="s">"Hello, world!"</span> <span class="c1">// string</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>+</td>
  <td>덧셈</td>
  <td>
  두 값을 더합니다. 사용할 수 있는 자료형은 정수, 실수, 복소수, 문자열입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">1</span> <span class="o">+</span> <span class="mi">2</span>                <span class="c1">// 3: 두 정수 더하기</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">2</span> <span class="o">+</span> <span class="mi">3</span>                <span class="c1">// 5: 두 정수 더하기</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="nx">a</span> <span class="o">+</span> <span class="nx">b</span>                <span class="c1">// 8: 두 변수 더하기</span>
<span class="nx">d</span> <span class="o">:=</span> <span class="s">"Hello, "</span> <span class="o">+</span> <span class="s">"world!"</span> <span class="c1">// Hello, world!: 두 문자열 붙이기</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>-</td>
  <td>뺄셈</td>
  <td>두 값의 차이를 구합니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span> <span class="o">-</span> <span class="mi">2</span> <span class="c1">// 1: 두 정수 빼기</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">4</span> <span class="o">-</span> <span class="mi">5</span> <span class="c1">// -1: 두 정수 빼기</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="nx">a</span> <span class="o">-</span> <span class="nx">b</span> <span class="c1">// 2: 두 변수 빼기</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>*</td>
  <td>곱셈</td>
  <td>두 값을 곱합니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">2</span> <span class="o">*</span> <span class="mi">3</span>  <span class="c1">// 6: 두 정수 곱하기</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">9</span> <span class="o">*</span> <span class="mi">21</span> <span class="c1">// 189: 두 정수 곱하기</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="nx">a</span> <span class="o">*</span> <span class="nx">b</span>  <span class="c1">// 1134: 두 변수 곱하기</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>/</td>
  <td>나눗셈</td>
  <td>두 값을 나눕니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">5</span> <span class="o">/</span> <span class="mi">2</span>  <span class="c1">// 2: 두 정수 나누기</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">12</span> <span class="o">/</span> <span class="mi">4</span> <span class="c1">// 3: 두 정수 나누기</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="nx">a</span> <span class="o">/</span> <span class="nx">b</span>  <span class="c1">// 0: 두 변수 나누기</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>%</td>
  <td>나머지</td>
  <td>두 값을 나눈 뒤 나머지를 구합니다. 사용할 수 있는 자료형은 정수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">5</span> <span class="o">%</span> <span class="mi">2</span> <span class="c1">// 1: 5를 2로 나누었을 때 나머지 구하기</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>+=</td>
  <td>덧셈 후 대입</td>
  <td>현재 변수와 값을 더한 다음 다시 변수에 대입합니다. 문자열은 현재 변수에 문자열을 붙인 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">5</span>         <span class="c1">// 5</span>
<span class="nx">a</span> <span class="o">+=</span> <span class="mi">2</span>         <span class="c1">// 7: a에 2를 더한 후 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 7</span>
</code></pre></div>



<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="s">"Hello, "</span> <span class="c1">// Hello,</span>
<span class="nx">a</span> <span class="o">+=</span> <span class="s">"world!"</span>  <span class="c1">// Hello, world!: a에 world! 문자열을 붙인 후 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// Hello, world!</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>-=</td>
  <td>뺄셈 후 대입</td>
  <td>현재 변수에서 값을 뺀 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">5</span>         <span class="c1">// 5</span>
<span class="nx">a</span> <span class="o">-=</span> <span class="mi">2</span>         <span class="c1">// 3: a에 2를 뺀 후 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 3</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>*=</td>
  <td>곱셈 후 대입</td>
  <td>현재 변수와 값을 곱한 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">5</span>         <span class="c1">// 5</span>
<span class="nx">a</span> <span class="o">*=</span> <span class="mi">2</span>         <span class="c1">// 10: a에 2를 곱한 후 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 10</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>/=</td>
  <td>나눗셈 후 대입</td>
  <td>현재 변수를 값으로 나눈 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">5</span>         <span class="c1">// 5</span>
<span class="nx">a</span> <span class="o">/=</span> <span class="mi">2</span>         <span class="c1">// 2: a에 2를 나눈 후 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 2</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>%=</td>
  <td>나머지를 구한 후 대입</td>
  <td>현재 변수와 값의 나머지를 구한 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">5</span>         <span class="c1">// 5</span>
<span class="nx">a</span> <span class="o">%=</span> <span class="mi">2</span>         <span class="c1">// 1: a에서 2를 나눈 후 나머지를 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 1</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&amp;</td>
  <td>AND 비트 연산</td>
  <td>두 값을 비트 단위로 AND 연산을 합니다. 사용할 수 있는 자료형은 정수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span>                <span class="c1">// 00000011</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">2</span>                <span class="c1">// 00000010</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="nx">a</span> <span class="o">&amp;</span> <span class="nx">b</span>            <span class="c1">// 00000010: a와 b의 AND 비트 연산</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">c</span><span class="p">)</span> <span class="c1">// 00000010</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>|</td>
  <td>OR 비트 연산</td>
  <td>두 값을 비트 단위로 OR 연산을 합니다. 사용할 수 있는 자료형은 정수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span>                <span class="c1">// 00000011</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">2</span>                <span class="c1">// 00000010</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="nx">a</span> <span class="p">|</span> <span class="nx">b</span>            <span class="c1">// 00000011: a와 b의 OR 비트 연산</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">c</span><span class="p">)</span> <span class="c1">// 00000011</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>^</td>
  <td>XOR 비트 연산(다항)</td>
  <td>두 값을 비트 단위로 XOR 연산을 합니다. 사용할 수 있는 자료형은 정수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span>                <span class="c1">// 00000011</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">2</span>                <span class="c1">// 00000010</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="nx">a</span> <span class="p">^</span> <span class="nx">b</span>            <span class="c1">// 00000001: a와 b의 XOR 비트 연산</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">c</span><span class="p">)</span> <span class="c1">// 00000001</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&amp;^</td>
  <td>AND NOT 비트 연산</td>
  <td>두 값을 비트 단위로 AND NOT 연산을 합니다. 즉 다음과 같이 특정 비트를 끕니다(Bit clear). 사용할 수 있는 자료형은 정수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">255</span>              <span class="c1">// 11111111</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">68</span>               <span class="c1">// 01000100</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="nx">a</span> <span class="o">&amp;^</span> <span class="nx">b</span>           <span class="c1">// 10111011: a와 b의 AND NOT 비트 연산</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">c</span><span class="p">)</span> <span class="c1">// 10111011</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&amp;=</td>
  <td>AND 비트 연산 후 대입</td>
  <td>현재 변수를 값으로 AND 연산한 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span>                <span class="c1">// 00000011</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">2</span>                <span class="c1">// 00000010</span>
<span class="nx">a</span> <span class="o">&amp;=</span> <span class="nx">b</span>                <span class="c1">// 00000010: a와 b를 AND 비트 연산 후 a에 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">a</span><span class="p">)</span> <span class="c1">// 00000010</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>|=</td>
  <td>OR 비트 연산 후 대입</td>
  <td>현재 변수를 값으로 OR 연산한 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span>                <span class="c1">// 00000011</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">2</span>                <span class="c1">// 00000010</span>
<span class="nx">a</span> <span class="o">|=</span> <span class="nx">b</span>                <span class="c1">// 00000011: a와 b를 OR 비트 연산 후 a에 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">a</span><span class="p">)</span> <span class="c1">// 00000011</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>^=</td>
  <td>XOR 비트 연산 후 대입</td>
  <td>현재 변수를 값으로 XOR 연산한 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span>                <span class="c1">// 00000011</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">2</span>                <span class="c1">// 00000010</span>
<span class="nx">a</span> <span class="p">^=</span> <span class="nx">b</span>                <span class="c1">// 00000001: a와 b를 XOR 비트 연산 후 a에 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">a</span><span class="p">)</span> <span class="c1">// 00000001</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&amp;^=</td>
  <td>AND NOT 비트 연산 후 대입</td>
  <td>현재 변수를 값으로 AND NOT 연산한 다음 다시 변수에 대입합니다. 이 연산자는 특정 플래그를 끌 때 주로 사용합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">255</span>              <span class="c1">// 11111111</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="mi">68</span>               <span class="c1">// 01000100</span>
<span class="nx">a</span> <span class="o">&amp;^=</span> <span class="nx">b</span>               <span class="c1">// 10111011: a와 b를 AND NOT 비트 연산 후 a에 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">a</span><span class="p">)</span> <span class="c1">// 10111011</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&lt;&lt;</td>
  <td>비트를 왼쪽으로 이동</td>
  <td>현재 값의 비트를 특정 횟수만큼 왼쪽으로 이동합니다. 사용할 수 있는 자료형은 정수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">7</span>                <span class="c1">// 00000111</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="nx">a</span> <span class="o">&lt;&lt;</span> <span class="mi">2</span>           <span class="c1">// 00011100: a의 비트를 오른쪽으로 2번 이동</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">b</span><span class="p">)</span> <span class="c1">// 00011100</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&gt;&gt;</td>
  <td>비트를 오른쪽으로 이동</td>
  <td>현재 값의 비트를 특정 횟수만큼 오른쪽으로 이동합니다. 사용할 수 있는 자료형은 정수입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">112</span>              <span class="c1">// 01110000</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="nx">a</span> <span class="o">&gt;&gt;</span> <span class="mi">3</span>           <span class="c1">// 00001110: a의 비트를 왼쪽으로 3번 이동</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">b</span><span class="p">)</span> <span class="c1">// 00001110</span>
</code></pre></div>
</td>
</tr>
<tr>
  <td>&lt;&lt;=</td>
  <td>비트를 왼쪽으로 이동 후 대입</td>
  <td>현재 변수를 특정 횟수만큼 왼쪽으로 이동한 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">7</span>                <span class="c1">// 00000111</span>
<span class="nx">a</span> <span class="o">&lt;&lt;=</span> <span class="mi">2</span>               <span class="c1">// 00011100: a의 비트를 왼쪽으로 2번 이동한 후 a에 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">a</span><span class="p">)</span> <span class="c1">// 00011100</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&gt;&gt;=</td>
  <td>비트를 오른쪽으로 이동 후 대입</td>
  <td>현재 변수를 특정 횟수만큼 오른쪽으로 이동한 다음 다시 변수에 대입합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">112</span>              <span class="c1">// 01110000</span>
<span class="nx">a</span> <span class="o">&gt;&gt;=</span> <span class="mi">3</span>               <span class="c1">// 00001110: a의 비트를 오른쪽으로 3번 이동한 a에 후 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">a</span><span class="p">)</span> <span class="c1">// 00001110</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>^</td>
  <td>비트 반전(단항)</td>
  <td>비트를 반전시킵니다(Bitwise complement, 1의 보수). 즉 0은 1로 1은 0으로 바꿉니다.


<div class="highlight"><pre><code class="go"><span class="kd">var</span> <span class="nx">a</span> <span class="kt">uint8</span> <span class="p">=</span> <span class="mi">3</span>       <span class="c1">// 00000011</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="p">^</span><span class="nx">a</span>               <span class="c1">// 11111100: a의 비트를 반전시김</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Printf</span><span class="p">(</span><span class="s">"%08b"</span><span class="p">,</span> <span class="nx">b</span><span class="p">)</span> <span class="c1">// 11111100</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>+</td>
  <td>양수 부호(단항)</td>
  <td>값에 양수 부호를 붙입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="o">-</span><span class="mi">2</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="o">+</span><span class="nx">a</span>        <span class="c1">// a에 양수 부호를 붙임</span>
<span class="nx">d</span> <span class="o">:=</span> <span class="o">+</span><span class="nx">b</span>        <span class="c1">// b에 양수 부호를 붙임</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">c</span><span class="p">)</span> <span class="c1">// 3: +(3)</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">d</span><span class="p">)</span> <span class="c1">// -2: +(-2)</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>-</td>
  <td>음수 부호 (단항)</td>
  <td>값에 음수 부호를 붙입니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">3</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="o">-</span><span class="mi">2</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="o">-</span><span class="nx">a</span>        <span class="c1">// a에 음수 부호를 붙임</span>
<span class="nx">d</span> <span class="o">:=</span> <span class="o">-</span><span class="nx">b</span>        <span class="c1">// b에 음수 부호를 붙임</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">c</span><span class="p">)</span> <span class="c1">// -3: -(3)</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">d</span><span class="p">)</span> <span class="c1">// 2: -(-2)</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>==</td>
  <td>같다</td>
  <td>두 값이 같은지 비교합니다.
  <ul>
  <li>실수형은 값을 연산한 뒤에는 오차가 발생하므로 ==로 비교할 때 주의해야 합니다.</li>
  <li>문자열은 내용이 같은지 비교합니다.</li>
  <li>배열은 요소의 내용이 모두 같은지 비교합니다.</li>
  <li>슬라이스와 맵은 배열과는 달리 내용을 비교할 수 없고, 메모리에 할당되어 있는지 확인합니다.</li>
  <li>포인터는 주소가 같은지 비교합니다.</li>
  </ul>


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mi">1</span> <span class="o">==</span> <span class="mi">1</span><span class="p">)</span>             <span class="c1">// true: 두 정수가 같으므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mf">3.5</span> <span class="o">==</span> <span class="mf">3.5</span><span class="p">)</span>         <span class="c1">// true: 두 실수가 같으므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="s">"Hello"</span> <span class="o">==</span> <span class="s">"Hello"</span><span class="p">)</span> <span class="c1">// true: 두 문자열이 같으므로 true</span>

<span class="nx">a</span> <span class="o">:=</span> <span class="p">[</span><span class="mi">3</span><span class="p">]</span><span class="kt">int</span><span class="p">{</span><span class="mi">1</span><span class="p">,</span> <span class="mi">2</span><span class="p">,</span> <span class="mi">3</span><span class="p">}</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="p">[</span><span class="mi">3</span><span class="p">]</span><span class="kt">int</span><span class="p">{</span><span class="mi">1</span><span class="p">,</span> <span class="mi">2</span><span class="p">,</span> <span class="mi">3</span><span class="p">}</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span> <span class="o">==</span> <span class="nx">b</span><span class="p">)</span> <span class="c1">// true: 두 배열이 같으므로 true</span>

<span class="nx">c</span> <span class="o">:=</span> <span class="p">[]</span><span class="kt">int</span><span class="p">{</span><span class="mi">1</span><span class="p">,</span> <span class="mi">2</span><span class="p">,</span> <span class="mi">3</span><span class="p">}</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">c</span> <span class="o">==</span> <span class="kc">nil</span><span class="p">)</span> <span class="c1">// false: 슬라이스를 nil과 비교하여 </span>
                      <span class="c1">// 메모리가 할당되었는지 확인</span>

<span class="nx">d</span> <span class="o">:=</span> <span class="kd">map</span><span class="p">[</span><span class="kt">string</span><span class="p">]</span><span class="kt">int</span><span class="p">{</span><span class="s">"Hello"</span><span class="p">:</span> <span class="mi">1</span><span class="p">}</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">d</span> <span class="o">==</span> <span class="kc">nil</span><span class="p">)</span> <span class="c1">// false: 맵을 nil과 비교하여 </span>
                      <span class="c1">// 메모리가 할당되었는지 확인</span>

<span class="nx">e</span> <span class="o">:=</span> <span class="mi">1</span>
<span class="kd">var</span> <span class="nx">p1</span> <span class="o">*</span><span class="kt">int</span> <span class="p">=</span> <span class="o">&amp;</span><span class="nx">e</span>
<span class="kd">var</span> <span class="nx">p2</span> <span class="o">*</span><span class="kt">int</span> <span class="p">=</span> <span class="o">&amp;</span><span class="nx">e</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">p1</span> <span class="o">==</span> <span class="nx">p2</span><span class="p">)</span> <span class="c1">// true: 포인터에 저장된 메모리 주소가 같으므로 true</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>!=</td>
  <td>같지 않다</td>
  <td>두 값이 다른지 비교합니다.


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mi">1</span> <span class="o">!=</span> <span class="mi">2</span><span class="p">)</span>             <span class="c1">// true: 두 정수가 다르므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mf">3.5</span> <span class="o">!=</span> <span class="mf">5.5</span><span class="p">)</span>         <span class="c1">// true: 두 실수가 다르므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="s">"Hello"</span> <span class="o">!=</span> <span class="s">"world"</span><span class="p">)</span> <span class="c1">// true: 두 문자열이 다르므로 true</span>

<span class="nx">a</span> <span class="o">:=</span> <span class="p">[</span><span class="mi">3</span><span class="p">]</span><span class="kt">int</span><span class="p">{</span><span class="mi">1</span><span class="p">,</span> <span class="mi">2</span><span class="p">,</span> <span class="mi">3</span><span class="p">}</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="p">[</span><span class="mi">3</span><span class="p">]</span><span class="kt">int</span><span class="p">{</span><span class="mi">3</span><span class="p">,</span> <span class="mi">2</span><span class="p">,</span> <span class="mi">1</span><span class="p">}</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span> <span class="o">!=</span> <span class="nx">b</span><span class="p">)</span> <span class="c1">// true: 두 배열이 다르므로 true</span>

<span class="nx">c</span> <span class="o">:=</span> <span class="p">[]</span><span class="kt">int</span><span class="p">{</span><span class="mi">1</span><span class="p">,</span> <span class="mi">2</span><span class="p">,</span> <span class="mi">3</span><span class="p">}</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">c</span> <span class="o">!=</span> <span class="kc">nil</span><span class="p">)</span> <span class="c1">// true: 슬라이스를 nil과 비교하여 </span>
                      <span class="c1">// 메모리가 할당되었는지 확인</span>

<span class="nx">d</span> <span class="o">:=</span> <span class="kd">map</span><span class="p">[</span><span class="kt">string</span><span class="p">]</span><span class="kt">int</span><span class="p">{</span><span class="s">"Hello"</span><span class="p">:</span> <span class="mi">1</span><span class="p">}</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">d</span> <span class="o">!=</span> <span class="kc">nil</span><span class="p">)</span> <span class="c1">// true: 맵을 nil과 비교하여 </span>
                      <span class="c1">// 메모리가 할당되었는지 확인</span>

<span class="nx">e</span> <span class="o">:=</span> <span class="mi">1</span>
<span class="nx">f</span> <span class="o">:=</span> <span class="mi">1</span>
<span class="kd">var</span> <span class="nx">p1</span> <span class="o">*</span><span class="kt">int</span> <span class="p">=</span> <span class="o">&amp;</span><span class="nx">e</span>
<span class="kd">var</span> <span class="nx">p2</span> <span class="o">*</span><span class="kt">int</span> <span class="p">=</span> <span class="o">&amp;</span><span class="nx">f</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">p1</span> <span class="o">!=</span> <span class="nx">p2</span><span class="p">)</span> <span class="c1">// true: 포인터에 저장된 메모리 주소가 다르므로 true</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&lt;</td>
  <td>작다</td>
  <td>앞의 값이 작은지 비교합니다. 문자열은 ASCII 코드 값을 기준으로 판단합니다. 또한, 첫 글자가 같다면 그 다음 글자부터 차례대로 비교하여 최종 값을 구합니다.


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mi">1</span> <span class="p">&lt;</span> <span class="mi">2</span><span class="p">)</span>             <span class="c1">// true: 1이 2보다 작으므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mf">3.5</span> <span class="p">&lt;</span> <span class="mf">5.5</span><span class="p">)</span>         <span class="c1">// true: 3.5가 5.5보다 작으므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="s">"Hello"</span> <span class="p">&lt;</span> <span class="s">"world"</span><span class="p">)</span> <span class="c1">// true: H가 w보다 ASCII 코드 값이 </span>
                               <span class="c1">// 작으므로 true</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&lt;=</td>
  <td>작거나 같다</td>
  <td>앞의 값이 작거나 같은지 비교합니다.


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mi">2</span> <span class="o">&lt;=</span> <span class="mi">2</span><span class="p">)</span>             <span class="c1">// true: 2가 2보다 작거나 같으므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mf">3.5</span> <span class="o">&lt;=</span> <span class="mf">5.5</span><span class="p">)</span>         <span class="c1">// true: 3.5가 5.5보다 작거나 같으므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="s">"Hello"</span> <span class="o">&lt;=</span> <span class="s">"world"</span><span class="p">)</span> <span class="c1">// true: H가 w보다 ASCII 코드 값이 </span>
                                <span class="c1">// 작거나 같으므로 true</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&gt;</td>
  <td>크다</td>
  <td>앞의 값이 큰지 비교합니다.


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mi">2</span> <span class="p">&gt;</span> <span class="mi">1</span><span class="p">)</span>             <span class="c1">// true: 2가 1보다 크므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mf">5.5</span> <span class="p">&gt;</span> <span class="mf">3.5</span><span class="p">)</span>         <span class="c1">// true: 5.5가 3.5보다 크므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="s">"world"</span> <span class="p">&gt;</span> <span class="s">"Hello"</span><span class="p">)</span> <span class="c1">// true: w가 H보다 ASCII 코드 값이 크므로 true</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&gt;=</td>
  <td>크거나 같다</td>
  <td>앞의 값이 크거나 같은지 비교합니다.


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mi">2</span> <span class="o">&gt;=</span> <span class="mi">2</span><span class="p">)</span>             <span class="c1">// true: 2가 2보다 크거나 같으므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="mf">5.5</span> <span class="o">&gt;=</span> <span class="mf">3.5</span><span class="p">)</span>         <span class="c1">// true: 5.5가 3.5보다 크거나 같으므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="s">"world"</span> <span class="o">&gt;=</span> <span class="s">"Hello"</span><span class="p">)</span> <span class="c1">// true: w가 H보다 ASCII 코드 값이 </span>
                                <span class="c1">// 크거나 같으므로 true</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&amp;&amp;</td>
  <td>AND 논리 연산</td>
  <td>두 불 값이 모두 참인지 확인합니다.


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="kc">true</span> <span class="o">&amp;&amp;</span> <span class="kc">true</span><span class="p">)</span>   <span class="c1">// true: 두 값이 모두 true이므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="kc">true</span> <span class="o">&amp;&amp;</span> <span class="kc">false</span><span class="p">)</span>  <span class="c1">// false: 두 값 중 하나가 false이므로 false</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="kc">false</span> <span class="o">&amp;&amp;</span> <span class="kc">false</span><span class="p">)</span> <span class="c1">// false: 두 값이 모두 false이므로 false</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>||</td>
  <td>OR 논리 연산</td>
  <td>두 불 값 중 한 개라도 참인지 확인합니다.


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="kc">true</span> <span class="o">||</span> <span class="kc">true</span><span class="p">)</span>   <span class="c1">// true: 두 값이 모두 true이므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="kc">true</span> <span class="o">||</span> <span class="kc">false</span><span class="p">)</span>  <span class="c1">// true: 두 값 중 하나가 true이므로 true</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="kc">false</span> <span class="o">||</span> <span class="kc">false</span><span class="p">)</span> <span class="c1">// false: 두 값이 모두 false이므로 false</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>!</td>
  <td>NOT 논리 연산</td>
  <td>불값을 반대로 연산합니다.


<div class="highlight"><pre><code class="go"><span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(!</span><span class="kc">true</span><span class="p">)</span>  <span class="c1">// false: true의 반대는 false</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(!</span><span class="kc">false</span><span class="p">)</span> <span class="c1">// true: false의 반대는 true</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&amp;</td>
  <td>참조(레퍼런스) 연산</td>
  <td>현재 변수의 메모리 주소를 구합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">1</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="o">&amp;</span><span class="nx">a</span>        <span class="c1">// a의 메모리 주소를 b에 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">b</span><span class="p">)</span> <span class="c1">// 0xc0820062d0 (메모리 주소)</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>*</td>
  <td>역참조 연산</td>
  <td>현재 포인터 변수에 저장된 메모리에 접근하여 값을 가져오거나 저장합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="nb">new</span><span class="p">(</span><span class="kt">int</span><span class="p">)</span>
<span class="o">*</span><span class="nx">a</span> <span class="p">=</span> <span class="mi">1</span>          <span class="c1">// a에 저장된 메모리에 접근하여 1을 저장</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="o">*</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 1: a에 저장된 메모리에 접근하여 값을 가져옴</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>&lt;-</td>
  <td>채널 수신 연산</td>
  <td>채널에 값을 보내거나 값을 가져옵니다.


<div class="highlight"><pre><code class="go"><span class="nx">c</span> <span class="o">:=</span> <span class="nb">make</span><span class="p">(</span><span class="kd">chan</span> <span class="kt">int</span><span class="p">)</span>

<span class="k">go</span> <span class="kd">func</span><span class="p">()</span> <span class="p">{</span>
	<span class="nx">c</span> <span class="o">&lt;-</span> <span class="mi">1</span> <span class="c1">// 채널 c에 1을 보냄</span>
<span class="p">}()</span>

<span class="nx">a</span> <span class="o">:=</span> <span class="o">&lt;-</span><span class="nx">c</span>       <span class="c1">// 채널 c에서 값을 가져와서 a에 대입</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 1</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>++</td>
  <td>증가</td>
  <td>변수의 값을 1 증가시킵니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다. 


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">1</span>
<span class="nx">a</span><span class="o">++</span> <span class="c1">// 2: 정수 1을 1 증가시켜서 2</span>

<span class="nx">b</span> <span class="o">:=</span> <span class="mf">1.5</span>
<span class="nx">b</span><span class="o">++</span> <span class="c1">// 2.5: 실수 1.5를 1 증가시켜서 2.5</span>

<span class="nx">c</span> <span class="o">:=</span> <span class="mi">1</span> <span class="o">+</span> <span class="m">2i</span>
<span class="nx">c</span><span class="o">++</span> <span class="c1">// (2+2i): 복소수 1+2i를 1 증가시켜서 2+2i</span>

<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 2</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">b</span><span class="p">)</span> <span class="c1">// 2.5</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">c</span><span class="p">)</span> <span class="c1">// (2+2i)</span>
</code></pre></div>


Go 언어에서는 ++ 연산자를 사용한 뒤 값을 대입할 수 없고, 변수 뒤에서만 사용할 수 있습니다. 따라서 ++ 연산자는 단독으로 사용하거나 <b>if</b> 조건문, <b>for</b> 반복문 안에서 주로 사용합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">1</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="nx">a</span><span class="o">++</span> <span class="c1">// 컴파일 에러</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="o">++</span><span class="nx">a</span> <span class="c1">// 컴파일 에러</span>
<span class="o">++</span><span class="nx">a</span>      <span class="c1">// 컴파일 에러</span>
</code></pre></div>

  </td>
</tr>
<tr>
  <td>--</td>
  <td>감소</td>
  <td>변수의 값을 1 감소시킵니다. 사용할 수 있는 자료형은 정수, 실수, 복소수입니다. 


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">1</span>
<span class="nx">a</span><span class="o">--</span> <span class="c1">// 0: 정수 1을 1 감소시켜서 0</span>

<span class="nx">b</span> <span class="o">:=</span> <span class="mf">1.5</span>
<span class="nx">b</span><span class="o">--</span> <span class="c1">// 0.5: 실수 1.5를 1 감소시켜서 0.5</span>

<span class="nx">c</span> <span class="o">:=</span> <span class="mi">1</span> <span class="o">+</span> <span class="m">2i</span>
<span class="nx">c</span><span class="o">--</span> <span class="c1">// (0+2i): 복소수 1+2i를 1 감소시켜서 0+2i</span>

<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">a</span><span class="p">)</span> <span class="c1">// 0</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">b</span><span class="p">)</span> <span class="c1">// 0.5</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nx">Println</span><span class="p">(</span><span class="nx">c</span><span class="p">)</span> <span class="c1">// (0+2i)</span>
</code></pre></div>


Go 언어에서는 -- 연산자를 사용한 뒤 값을 대입할 수 없고, 변수 뒤에서만 사용할 수 있습니다. 따라서 -- 연산자는 단독으로 사용하거나 <b>if</b> 조건문, <b>for</b> 반복문 안에서 주로 사용합니다.


<div class="highlight"><pre><code class="go"><span class="nx">a</span> <span class="o">:=</span> <span class="mi">1</span>
<span class="nx">b</span> <span class="o">:=</span> <span class="nx">a</span><span class="o">--</span> <span class="c1">// 컴파일 에러</span>
<span class="nx">c</span> <span class="o">:=</span> <span class="o">--</span><span class="nx">a</span> <span class="c1">// 컴파일 에러</span>
<span class="o">--</span><span class="nx">a</span>      <span class="c1">// 컴파일 에러</span>
</code></pre></div>

  </td>
</tr>
</tbody></table>
kmykさんの[online-judge-tools](https://online-judge-tools.readthedocs.io/en/master/introduction.ja.html)と
Vimを組み合わせて使い始めてから半年ぐらい経ちました。
それから使い続けて以来、新たに不満は生じず「割と便利な運用なのでは？」と思えるようになってきたので、単純なものではあるものの紹介してみようと思った次第です。

## 競プロに便利なCUIツールを求めて

とりあえず、以下の2つの課題をなんとかしたいなと思っていました。

### 1. サンプルのテスト

今となってはちょっと考えられないですが、toolを導入するまでは、サンプルのチェックのために
**問題のすべてのサンプルのコピペを繰り返す** 、ということをしていました。

デバッグが不要で、一発で通せるぐらい自分にとって易しい問題であればこのコストを受け入れてもいいかもしれません。
しかしながら、コンテストで通すべき問題というのは、ときにはデバッグ出力を何回も確認しながら慎重に実装したり、
WAによってコードの修正とサンプルテストの何回もの繰り返しが余儀なくされるものです。
よって、 **自身にとって重要な問題であるほど、このコストは大きくなっていきます。**

ですので、 **「CUIからの一回のコマンドにより、一括で問題ごとのサンプルすべてが検証される」** のが理想的です。

### 2. コードの提出

tool導入前は、
**「エディタのコードをコピーして、問題の提出欄にペーストし、選択言語が正しいことを確認してからボタンを押下する」**
という作業をしていました。

サンプルのテストほどではないですが、これもいくつかの手順があり、更に問題になるのは以下のような事項だと思います。

1. 選択言語を間違えてCEしてしまう（単純に時間の無駄）。
2. 提出先の問題を間違えてREもしくはWAしてしまう（最終的にペナになる可能性があり最悪）。
3. ブザービートに失敗する（レアケースといえばレアケースだが、逃した時のショックは大きそう）。
4. コードのコピペをミスる（Vimだと普通にコピペするとクリップボードに載らない((設定で、Vim内のコピー先をクリップボードと共有することはできますが、個人的に好きじゃないのでその点はデフォルトのままとしています。そのせいで以前に一度、直前の問題のコードを間違って貼って提出し、REしてペナを貰うというのをやったことがあります))）。

よって、 **「確実に今自分が目に入れているエディタのコードを、正しい問題・正しい選択言語で提出できる」** のが理想的です。

## できるだけシンプルなツールを求めて

一応、[atcoder-tools](https://github.com/kyuridenamida/atcoder-tools)の存在は知っていたのですが、
一見したところ「C++やPythonといった競プロメジャー言語に寄っているっぽい（他の言語は使えない or 使いづらい？）」
とか「特定のディレクトリ構成が強要されるっぽい（逸脱しようとすると凝ったことをしないとダメそう or 調査は必須）」
という印象を持ち、ちょっと自分の要件には合わないかなぁと思っていました((提示されているデフォルトの使用方法がマッチしているという方には、とても便利なツールなのだと思います。))。

結局、頑張って自分用に自作していたのですが((大方の機能は実装できたのですが、古いARCの問題のクローリングで早々とコケてしまい、途方に暮れていたところojに出会いました（圧倒的感謝）。))、
ふとしたきっかけで[online-judge-tools](https://online-judge-tools.readthedocs.io/en/master/introduction.ja.html)
の存在を知り、また自分の求めているものにかなりマッチしていると気づきました。

- サンプルのダウンロード: `oj download -d {{target_directory}} {{problem_url}}`
  - オプションでダウンロード先を簡単に変更できるのが嬉しい。
- サンプルのテスト: `oj test -c "go run {{target_program}}" -d {{sample_directory}} -t 4` （Goの場合）
  - オプションで実行コマンドを変えられるので、他言語の対応も簡単そう。
- コードの提出: `oj submit -y {{problem_url}} {{target_source_file}}`
  - 提出言語を推定してくれるので、ヒューマンエラーがない。

また、AtCoderのみならず、Codeforces、yukicoder、AOJといった主要なサイトに対応しているのも、非常にありがたいですね。

## Vimから呼び出せるよう、連携しようという試み

私は普段エディタにVim（厳密にはneovimですが）を使っており、またVimの利点として **「ターミナルやCUIとの距離感が近い」** というものがあると思っています。
個人的には、他のエディタに比べて、CUIツールとの連携が簡単にできるのではないかと感じています。

よって、 **「先述のojコマンドをいい感じに呼び出すVimのコマンドを定義すること」** を目指します。

### 仕様

Vimのコマンドラインモード（コロン打ったら遷移するモード）から、以下のコマンドを打てるようにします。

- サンプルのダウンロードコマンド: `:DonwloadSamples`
  - コマンドを実行すると、 **今エディタに載っている問題のサンプルが同じディレクトリ階層にDLされる。**
    - 例えばコンテスト中のコードは `contests/2020/08/20200815_ABC175/a/a.go` みたいに整理しているのですが、この問題のサンプルは `contests/2020/08/20200815_ABC175/a/test/` ディレクトリに収まってほしい、という具合です。
- サンプルのテストコマンド: `:TestCurrentBuffer`
  - コマンドを実行すると、 **今エディタに載っているコードに対してすべてのサンプルが検証される。**
    - 先程DLしたものが素直に実行されてほしい、という具合です。
- コードの提出コマンド: `:SubmitCode`
  - コマンドを実行すると、 **今エディタに載っているコードが対応する問題に対して提出される。**
    - 検証が済んだらそのままの流れでシームレスに提出まで持っていきたい、という具合です。

### コマンドのデモ

各コマンドの動作イメージは、以下のようなものになります。

<figure class="figure-image figure-image-fotolife" title="サンプルのダウンロード">[f:id:maguroguma:20200819004910g:plain]<figcaption>サンプルのダウンロード</figcaption></figure>

<figure class="figure-image figure-image-fotolife" title="サンプルのテスト">[f:id:maguroguma:20200819004818g:plain]<figcaption>サンプルのテスト</figcaption></figure>

<figure class="figure-image figure-image-fotolife" title="コードの提出">[f:id:maguroguma:20200819005716g:plain]<figcaption>コードの提出</figcaption></figure>

### 各コマンドを定義するVim script

各コマンドについて1つずつ観ていきます。

```vim
" ファイル上部に記述される「問題のURL」を取得する関数
function! s:ReadProblemURLFromCurrentBuffer()
  let l:lines = getline(0, line("$"))
  for l:line in l:lines
    let l:record = split(l:line, ' ')
    for l:r in l:record
      let l:url = matchstr(r, '^\(http\|https\):.*$')
      if l:url != ''
        return l:url
      endif
    endfor
  endfor
  return ''
endfunction
```

はい、いきなり **「コンテスタントの運用でカバー」** 的な要素があります。
この関数は、 **現在ロードしているソースファイルの上部に「問題のURL」が記載されていることを期待** しています。

最初は、コマンドの引数に問題のURLを渡す設計で考えていたのですが、
このURLはコード提出時にも必要になることから、
**「ファイル中のコメントとしてはじめに一度だけペーストしてしまうほうが、以降の手間もミスもなくなって良いのではないか？」**
と思い、このようにしました。

なので、私の競プロのルーティンとして **「問題を開いたらURLをファイルのトップにコピーする」**
というものが組み込まれることとなりました((案外気にならない上に、何度も提出する必要がある難しい問題になるほど、恩恵は大きくなります。後は、問題の復習をするときにコードからすぐに問題ページを開けるのもよいです。))。

そして、サンプルのダウンロードコマンドが以下になります。

```vim
" サンプルダウンロードのための関数とコマンド
function! s:MakeSampleDLCommand(url)
  let l:cur_buf_dir = expand("%:h")
  let l:target_dir = l:cur_buf_dir . "/test"
  let l:dl_command = printf("oj download -d %s %s", l:target_dir, a:url)
  return l:dl_command
endfunction
function! s:DownloadSamples(url)
  let l:command = s:MakeSampleDLCommand(a:url)
  echo "[Run] " . l:command . "\n"
  call execute('vs')
  call execute('terminal ' . l:command)
endfunction

command! -nargs=0 DownloadSamples :call s:DownloadSamples(s:ReadProblemURLFromCurrentBuffer())
```

やっていることは、
**「Vim scriptで本来ターミナルで実行したいコマンドを組み立てて、Vimの `tarminal` コマンドに渡して実行させている」** 、というだけです。
以降のコマンドでもそうですが、 `system()` 関数で実行し結果を `echo` するよりも、
見栄え的にこちらのほうがいい感じです（多分）。

続いて、ダウンロードしたサンプルの実行コマンドです。

```vim
" サンプルテストのための関数とコマンド
function! s:MakeTestSamplesCommand()
  let l:cur_buf_go = expand("%")
  let l:cur_buf_dir = expand("%:h")
  let l:sample_file_dir = l:cur_buf_dir . "/test"
  let l:test_command = printf("oj test -c \"go run %s\" -d %s -t 4", l:cur_buf_go, l:sample_file_dir)
  return l:test_command
endfunction
function! s:TestSamples()
  let l:command = s:MakeTestSamplesCommand()
  echo "[Run] " . l:command . "\n"
  call execute('vs')
  call execute('terminal ' . l:command)
endfunction

" Go版テスト実行コマンド
command! -nargs=0 TestCurrentBufferGoCode :call s:TestSamples()
```

これも、コマンドの組み立て部分が微妙に変わっただけで、ダウンロードとほとんど変わらないですね。
実行コマンドを差し替えたものを用意すれば、他の好きな言語の実行コマンドも作れると思います((たまにBashの練習に競プロを使ったりするので、Bashバージョンも持っていたりします。))。

最後に、コードの提出コマンドになります。

```vim
" コード提出のための関数とコマンド定義
function! s:MakeSubmitCommand(url)
  let l:cur_buf_go = expand("%")
  let l:submit_command = printf("oj submit -y %s %s", a:url, l:cur_buf_go)
  return l:submit_command
endfunction
function! s:SubmitCode(url)
  let l:command = s:MakeSubmitCommand(a:url)
  echo "[Run] " . l:command . "\n"
  call execute('vs')
  call execute('terminal ' . l:command)
endfunction

command! -nargs=0 SubmitCode :call s:SubmitCode(s:ReadProblemURLFromCurrentBuffer())
```

サンプルのダウンロードの際に必要となったURLが、ここでも必要となります。

以上で紹介したコマンドや関数は、すべてojが実行可能であることを前提としたものですので、
スクリプトファイルとするにあたっては、以下のように実行可能時のみ定義するようにするのが良いかと思います。

```vim
if executable('oj')
  " ファイル上部に記述される「問題のURL」を取得する関数
  function! s:ReadProblemURLFromCurrentBuffer()
  ...
  command! -nargs=0 SubmitCode :call s:SubmitCode(s:ReadProblemURLFromCurrentBuffer())
endif
```

## 最後に

もはやAtCoder Problemsと同じくらい「これがなきゃ競プロやってらんねぇ」なツールになってきたので、
感謝するだけじゃなくcommitできるようコード読まないとなぁと思います。

kmykさんおよびコミッタの皆様、本当にありがとうございます。

あと、これくらい簡単なVim scriptが書けるだけでも、自分用の便利コマンドは案外簡単に作れたりするので、ぜひVimを使いましょう！

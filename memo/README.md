# 見出し
# 箇条書き/リスト
# 水平線
# リンク/画像
# コード/引用
# 表

# 見出し
先頭行に # を書く
vimで検索する時は /^# で検索し見出しを行き来できる

# 箇条書き/リスト
aaa
bbb
 ccc
この3行をビジュアルモードで選択した状態で :'<,'>s/^/* / を実行
* aaa
* bbb
* ccc
箇条書きの完成

# 水平線
https://www.google.co.jp/
URLをヴィジュアルモードで選択状態で
\mduをタイプ
[] (https://www.google.co.jp/)
[]の中にカーソルが移動する。あとはタイトルを入力

# コード/引用
基本的な記述
```javascript
  console.log("JavaScript");
```
.vimrcに下記の設定をすれば
markdownに埋め込まれたコードに対しても色つけできる
let g:markdown_fenced_languages = [
\  'css',
\  'javascript',
\  'js=javascript',
\  'json=javascript',
\  'ruby',
\  'python',
\]
引用は、引用表記したい部分をビジュアルモードで選択して
:'<,'>s/^/* /


# 表
品名|値段
コーラ|120
ハンバーガー|200
スマイル|0
表部分をビジュアルモードで選択して :'<,'>Alignta |
品名         | 値段
-------------|------:
コーラ       | 120
ハンバーガー | 200
スマイル     | 0

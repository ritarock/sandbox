# 1章のはなし
特記事項なし

# 2章のはなし
* scaffoldジェネレータを使えばMVCとかCRUDの雛形を作ってくれてる感じ?
```bash
rails generate scaffold リソース名の単数形 データモデル名:データ型 ・・・
```
* コントローラ内で宣言したインスタンス変数(@で始まる変数)はビューで使える
* データの関連付け
```ruby
has_many :microposts
とか
belongs_to :user
```
* rails consoleを使ってDBを操作

# 3章のはなし
アクションはCの中に置く  
generateスクリプトでアクションを作成
```bash
rails generate コントローラ名(キャメルケース) アクション名(小文字) ・・・
```

railsコマンドの短縮形
| 完全なコマンド | 短縮形  |
| :------------- | :-----  |
| rails server   | rails s |
| rails console  | rails c |
| rails generate | rails g |
| rails test     | rails t |
| bundle install | bundle  |

## TODO
yield

erbについて
<% %>だと実行するだけ、<%= %>だと実行結果を出力する

# 4章のはなし
後置if
```ruby
puts "x is not empty" if !x.empty?
```

2回否定で論理値に変更  
falseとnilのみがfalse

p :name と puts :name.inspect は同価

# 5章のはなし
リンクの生成
```ruby
<% link_to "リンクテキスト", "URL" %>
```

image_tag ヘルパーを使用した場合、app/assets/imagesから探す  

パーシャル
```ruby
<%= render 'layouts/shim' %>
```
この場合、app/views/layouts/_shim.html.erbの内容を評価する  
_始まりが命名規約

アセットディレクトリ  
アプリケーション固有のアセットは app/assets  
チームによって開発されたアセットは lib/assets  
サードパーティのアセットは vendor/assets

## 統合テスト
テストのテンプレートを作る
```bash
rails generate integration_test テスト名
```

テスト内容はこんな感じ
```ruby
assert_select "a[href=?]", about_path
```
上記で探すhtmlのリンクは下記
```html
<a href="/about">.....</a>
```

リンクの個数を数える
```ruby
ssert_select "a[href=?]", root_path, count: 2
```

メソッドで複雑なテストはしない方が賢明  
今回のようなレイアウト内で頻繁に変更されるHTML要素 (リンクなど) をテストするぐらいに抑えておくとよい

テストの実行
```bash
rails test:integration
```

# 6章のはなし
モデルの作成
```ruby
rails generate model モデル名 属性:型情報 ・・・
```

マイグレーションは、データベースの構造をインクリメンタルに変更する手段を提供  
マイグレーション適応
```bash
rails db:migrate
```

rails console のサンドボックスモードでの起動
```bash
rails console --sandbox
```

文字列の配列を簡単に作る
```bash
>> %w[foo bar baz]
=> ["foo", "bar", "baz"]
```

## バリデーション
* 存在性検証
* 長さ検証
* フォーマット検証
* 一意性検証

## セキュアなパスワードの実装
bcryptが必要
has_secure_passwordメソッドを呼び出すだけ
パスワードの最小文字数


# 7章のはなし
リソースの追加
```ruby
resources リソースシンボル
```
この1行でRestfullなアクションを全て使用できる

debuggerを使えば値の確認等が出来る

## form_for
使い方
```ruby
<%= form_for(@user) do |f|
・
・
・
<% end %>
```
このfオブジェクトに関して  
@userの属性を設定するために特別に設計されたHTMLを返す

```ruby
<%= f.label :name %>
<%= f.text_field :name %>
```
これならUserモデルのname属性を設定する、ラベル付きテキストフィールド要素を作成するのに必要なHTMLを作成
```html
<label for="user_name">Name</label>
<input id="user_name" name="user[name]" type="text" />
```

## pluralize
単語を複数形にする
```bash
>> helper.pluralize(1, "error")
=> "1 error"

>> helper.pluralize(5, "error")
=> "5 errors"
```

## redirect
以下は等価
```ruby
redirect_to @user

redirect_to user_url(@user)
```

## flash
こんな使い方
```ruby
flash[:success] = "Welcome to the Sample App!"
```
successのときにメッセージを代入する

## DBの初期化
```bash
rails db:migrate:reset
```

```bash
$ rails g model (単数形) カラム:型 カラム:型 ・・・
$ rails db:migrate
$ rails g controller (複数形)
```

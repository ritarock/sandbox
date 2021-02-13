# クラスを作る

キャメルケースで書くのが一般的。
```ruby
class Sample
end
```

この時に呼び出されるメソッドがinitializeメソッド
# initializeメソッド
```ruby
sample = Sample.new
```

```ruby
class Sample
  def initialize
    p "initialized"
  end
end

sample = Sample.new
```

引数付き
```ruby
class Sample
  def initialize(str)
    p "#{str}"
  end
end

sample = Sample.new("hello")
```

# メソッドの定義
```ruby
class Sample
  def hello
    "hello"
  end
end

sample = Sample.new
p sample.hello
```

## インスタンス変数
インスタンス変数は@から始める
```ruby
class Sample
  def initialize(str)
    @str = str
  end

  def hello
    "hello, #{@str}"
  end
end

sample = Sample.new('world')
p sample.hello
```

## インスタンス変数に外部からアクセス
参照用のメソッドを作成
```ruby
class Sample
  def initialize(str)
    @str = str
  end

  def str
    @str
  end
end

sample = Sample.new("hello")
p sample.str
```

## インスタンス変数を外部から変更
```ruby
class Sample
  def initialize(str)
    @str = str
  end

  def str
    @str
  end

  def str=(value)
    @str = value
  end
end

sample = Sample.new("sanple")
p sample.str
sample.str = "sample"
p sample.str
```

# アクセサメソッドの省略
```ruby
class Sample
  attr_accessor :str

  def initialize(str)
    @str = str
  end
end
sample = Sample.new("sanple")
p sample.str
sample.str = "sample"
p sample.str
```
* 読み取り専用
  * attr_reader
* 書き込み専用
  * attr_writer
* どっちも
  * attr_accessor

## クラスメソッドを定義
```ruby
class Sample
  def initialize(str)
    @str = str
  end

  # selfを付けるとクラスメソッド
  def self.create_string(strings)
    strings.map do |string|
      Sample.new(string)
    end
  end

  def hello
    "Hello, #{@str}"
  end
end

names = ["A", "B", "C"]
sample = Sample.create_string(names)
sample.each do |n|
  p n.hello
end
```

# 継承
書き換え

```ruby
class S
  attr_reader :first, :second
  def initialize(first, second)
    @first = first
    @second = second
  end
end

class F < S
  attr_reader :third
  def initialize(first, second, third)
    @first = first
    @second = second
    @third = third
  end
end
```

サブクラスででスーパークラスを呼び出す
```ruby
class S
  attr_reader :first, :second
  def initialize(first, second)
    @first = first
    @second = second
  end
end

class S < F
  attr_reader :third
  def initialize(first, second, third)
    super(first, second)
    @third = third
  end
end
```

スーパークラスとサブクラスで引数の数が同じ場合、引数なしのsuperを呼ぶだけ

```ruby
class S
  attr_reader :first, :second
  def initialize(first, second)
    @first = first
    @second = second
  end
end

class S < F
  attr_reader :third
  def initialize(first, second, third)
    super
    @third = third
  end
end
```

# メソッドの公開範囲
## public
クラスの外部から呼び出せるメソッド
```ruby
class Sample
  def hello
    'hello'
  end
end
```

## private
レシーバを指定して呼び出せないメソッド  
クラスの外部からは呼び出せず内部のみで使用できる
```ruby
class Sample
  def hello
    "hello #{name}"
  end

  private
  def name
    "XXX"
  end
end
```

## protected
定義したクラス自身と、サブクラスからのみレシーバ付きで呼び出せる

# 定数
```ruby
class S
  DEFAULT = 0
end
```

参照
```ruby
S::DEFAULT
```

# モジュール
モジュールとクラスの比較
* モジュールからインスタンスを作成することはできない
* 他のモジュールやクラスを継承できない

## ミックスイン
is-a関係成り立たない共通の機能を持たせる時に使う

## モジュールのエクステンド
クラスメソッドとしてミックスインする
```ruby
extend module
```

# yield
ブロックをメソッドの引数にして受け取る
```ruby
def method(&block)
  block.call
end
```

# Procオブジェクト
定義
```ruby
sample_proc = Proc.new do
  "sample"
end
```
```ruby
sample_proc = Proc.new { "sample" }
```

procオブジェクトを実行する場合はcallメソッドを使う
```ruby
sample_proc = Proc.new { 'sample' }
sample_proc.call
```

# ラムダ
```ruby
->(a, b) { a + b }
lambda {|a, b| a + b}
```

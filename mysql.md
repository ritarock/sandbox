# docker
```
$ docker run --name mysql -e MYSQL_ROOT_PASSWORD=mysql -d -p 3306:3306 mysql
$ docker exec -it コンテナ名 bash
# mysql -u root -p mysq
mysql >
```


# データベースの操作
## データベースの作成
```sql
CREATE DATABASE データベース名;
```

## データベースの削除
```sql
DROP DATABASE データベース名;
```

## データベースの存在確認
```sql
SHOW DATABASES;
```

## カレントデータベースの変更
```sql
USE データベース名;
```

# テーブルの作成
```sql
CREATE TABLE テーブル名 (列名 型,...) [テーブルオプション];
```
## 列定義オプション
| 種類                 | 定義オプション | 詳細                           |
| :------------------- | :------------  | :----------------------------- |
| 主キー               | PRIMARY KEY    | 重複とNULLはNG                 |
| ユニークキー         | UNIQUE         | 重複はNG, NULLはOK             |
| ノンユニークキー     | KEY            | 重複OK                         |
| 空間インデックス     | SPATIAL        | 座標などの空間情報インデックス |
| 全件検索インデックス | FULLTEXT       | 全文検索インデックス           |

# テーブルの削除
```sql
DROP TABLE テーブル名;
```

# データの挿入
useで指定している場合はデータベース名は省略可
```sql
INSERT INTO データベース名.テーブル名(列名1, 列名2,...) VALUES (値1, 値2,...);
```
全てのカラムに値を追加する場合
```sql
INSERT INTO テーブル名 VVALUES (値1, 値2,...);
```

# データの出力
```sql
SELECT * FROM テーブル名;
```

# データの検索
```sql
SELECT 列名1, 列名2,... FROM テーブル名 [条件];
```

## 条件
### 比較演算子
| 演算子  | 内容          |
| :------ | :------------ |
| =       | 等しい        |
| <       | 小さい        |
| >       | 大きい        |
| <=      | 以下          |
| >=      | 以上          |
| <>      | 等しくない    |

### BETWEEN
指定した範囲に値があるものを取得
```sql
SELECT * FROM テーブル名 WHERE 列名 BETWEEN 値1 AND 値2;
```

### IN
指定した値のリストの中にあるものを取得
```sql
SELECT * FROM テーブル名 WHERE 列名 IN ('値1', '値2');
```

### LIKE
%は任意の文字数の任意の文字、_は1文字の任意の文字
```sql
SELECT * FROM テーブル名 WHERE 列名 LIKE 検索条件('_100%'とか)
```

### サブクエリ
副問い合わせ
```sql
SELECT * FROM テーブル名 WHERE 列名1 = (SELECT * FROM テーブル名 WHERE 列名2 = 値);
```

# 並べ替え
## ORDER BY
ASCは昇順、DESCは降順、ASCは省略可能
```sql
SELECT * FROM テーブル名 ORDER BY 列名 ASC;
```

# 集約
## GROUP BY
重複を除いたりできる
```sql
SELECT 列名 FROM テーブル名 GROUP BY 列名;
```
distinct でも重複を除ける
```sql
SELECT distinct 列名 FROM テーブル名;
```

# 集計関数
```sql
SELECT 列名, 集計関数(集計をとる列名) FROM テーブル名 GROUP BY 列名;
```

## 集計関数の種類
| 関数名  | 内容          |
| :------ | :------------ |
| MAX     | 最大値        |
| MIN     | 最小値        |
| SUM     | 合計          |
| AVG     | 平均          |
| COUNT   | カウント      |

# 集計処理の条件
## HAVING句
```sql
SELECT 列名1, AVG(列名2) FROM テーブル名 GROUP BY 列名1 HAVING COUNT(列名1) >= 値;
```

# テーブルの結合
## 内部結合
テーブルの指定した列の値が一致するデータのみを取得
```sql
SELECT テーブル名.カラム名,... FROM テーブル名1
INNER JOIN テーブル名2 ON テーブル名1.カラム名1 = テーブル名2.カラム名2;
```
### USING
共通の列名でJOINする
```sql
SELECT * FROM テーブル1 INNER JOIN テーブル2 USING(列名)
```

## 外部結合
テーブルの指定した列の値が一致するデータとどちらかのテーブルにしか存在しないデータも取得
```sql
SELECT テーブル名.カラム名,... FROM テーブル名1
(LEFT | RIGHT) OUTER JOIN テーブル名2 ON テーブル名1.カラム名1 = テーブル名2.カラム名2;
```
| 書式             | 内容                                         |
| :--------------- | :------------------------------------------- |
| LEFT OUTER JOIN  | FROMの後に書かれたテーブルのデータだけを取得 |
| RIGHT OUTER JOIN | JOINの後に書かれたテーブルのデータだけを取得 |
内部結合同様USINGが使える
```sql
SELECT * FROM テーブル名1 LEFT OUTER JOIN テーブル名2 USING(列名)
```

## 交差結合 (クラス結合)
2つのテーブルの組み合わせを作る
```sql
SELECT * FROM テーブル1 CROSS JOIN テーブル2;
```

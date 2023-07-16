## Partial<Type>

`Partial<Type>` で指定した型に一時的に切り替えることができる。値を返すときはキャストする必要がある。

## Required<Type>

パラメータが欠損したらエラー。 `Partial` の逆の使い方。

## Readonly<Type>

初期化したあとにパラメータを変更できないようにする。

## Record<Keys, Type>

Key と その型を指定できる。辞書として使いたい場合に適している。

## Pick<Type, Keys>

keys に指定したプロパティの型になる。

## Omit<Type, Keys>

keys に指定した以外のプロパティの型になる。

## Exclude<Type, ExcludedUnion>

`Exclude<T, U>` T から U を除いた型になる。

## Extract<Type, Union>

`Extract<T, U>` T と U のユニオン型になる。

## NonNullable<Type>

`null` と `undefined` を除いた型になる。

## Parameters<Type>

`Parameters<T>` T は関数型の必要がある。 T が関数のとき T の引数一覧をタプル型で作成する。

## ConstructorParameters<Type>

class の コンストラクタ関数の引数の型からタプル型を作成する。

## ReturnType<Type>

戻り値からなる型になる。

## InstanceType<Type>

コンストラクタの戻り値からなる型。

## ThisParameterType<Type>

関数型の this パラメータの型を抽出した型になる。

## OmitThisParameter<Type>

関数型の this パラメータの型を除いた型になる。

## ThisType<Type>

オブジェクト内の this の型を正しい型にする。

import pandas as pd
import numpy as np

wine = pd.read_csv("winequality-red.csv", sep=";")

# 単回帰分析
# from sklearn import linear_model
# clf = linear_model.LinearRegression()

# 説明変数に"density(濃度)"を利用
# X = wine.loc[:, ['density']].as_matrix()

# 目的変数に"alcohol"を利用
# Y = wine['alcohol'].as_matrix()

# モデルを作成
# clf.fit(X, Y)

# 回帰係数
# print(clf.coef_)
# 切片
# print(clf.intercept_)
# 決定係数
# print(clf.score(X, Y))

# 重回帰分析
from sklearn import linear_model
clf = linear_model.LinearRegression()
# 説明変数に"quality"以外を利用
wine_except_quality = wine.drop("quality", axis=1)
X = wine_except_quality.as_matrix()

# 目的変数に"quality"を利用
Y = wine['quality'].as_matrix()

# モデルを作成
clf.fit(X, Y)

# 偏回帰係数
print(pd.DataFrame({"Name":wine_except_quality.columns,
                    "Coefficients":clf.coef_}).sort_values(by = 'Coefficients') )

# 切片
print(clf.intercept_)

# 3σを基準に外れ値をNanで埋める関数
def sigma(target):
   for i in range(len(target)):
       # 平均と標準偏差
       average = np.mean(target)
       sd = np.std(target)

       # 外れ値の基準点
       outlier_min = average - (sd) * 3
       outlier_max = average + (sd) * 3

       # 範囲から外れている値を除く
       target[target < outlier_min] = None
       target[target > outlier_max] = None

       return target

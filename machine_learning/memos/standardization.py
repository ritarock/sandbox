# 標準化する関数
def zscore(target):
   target_mean = target.mean()
   target_std = np.std(target)

   zscore_point = (target - target_mean) / target_std

   return zscore_point

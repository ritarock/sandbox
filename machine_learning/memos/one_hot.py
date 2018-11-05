# one-hot変換関数
def one_hot(target):
   from sklearn import preprocessing
   from sklearn.preprocessing import OneHotEncoder
   target_vector = target.values
   target_vector_enc = preprocessing.LabelEncoder().fit_transform(target_vector).reshape(-1,1)
   profession_enc = OneHotEncoder().fit_transform(target_vector_enc).toarray()
   return profession_enc

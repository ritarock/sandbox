def main():
    xList = [1,2,3,4,5,6,7,8,9,10]
    yList = [5,10,15,20,25,30,35,40,45,50]

    n = len(xList)

    # 平均値
    xAve = sum(xList) / n
    yAve = sum(yList) / n

    # xの分散
    xDispersion = dispersion(xList,xAve)

    # 共分散
    xyCovariance = covariance(xList,xAve,yList,yAve)

    a = xyCovariance / xDispersion
    b = a * (-xAve) + yAve

    print('y = {0}x + {1}'.format(a,b))

def covariance(xList,xAve,yList,yAve):
    x = []
    y = []
    xy = []

    for i in xList:
        x.append(i - xAve)

    for i in yList:
        y.append(i - yAve)

    for i in range(len(xList)):
        xy.append(x[i] * y[i])

    covarianceDate = sum(xy) / len(xList)

    return covarianceDate

def dispersion(dateList,average):
    deviation = []
    squareDeviation = []
    # 偏差
    for i in dateList:
        deviation.append(i - average)

    # 偏差の2乗
    for i in deviation:
        squareDeviation.append(pow(i,2))

    dispersionDate = sum(squareDeviation) / len(deviation)

    return dispersionDate

if __name__ == '__main__':
    main()

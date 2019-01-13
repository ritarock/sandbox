import glob
import os


def main():
    path = './point/'
    files = glob.glob(path + '*')

    for i in range(len(files)):
        os.rename(files[i], path + '{0:03d}'.format(i) + '.txt')


if __name__ == '__main__':
    main()

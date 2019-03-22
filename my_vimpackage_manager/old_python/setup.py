import os
import subprocess

START_PATH = os.path.expanduser('~') + "/.vim/pack/mypackage/start/"
OPT_PATH = os.path.expanduser('~') + "/.vim/pack/mypackage/opt/"


def start():
    check_git = ['git', '--version']

    try:
        print("installed git")
    except:
        print("install git!")

    if os.path.isdir(START_PATH):
        pass
    else:
        os.makedirs(START_PATH)

    if os.path.isdir(OPT_PATH):
        pass
    else:
        os.makedirs(OPT_PATH)

    print('OK!')


def main():
    start()


if __name__ == '__main__':
    main()

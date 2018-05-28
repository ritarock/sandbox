import os
import subprocess

def start():
    check_git = ['git', '--version']
    try:
        print("installed git")
    except:
        print("install git!")

    home_path = os.path.expanduser('~')
    start_path = home_path + "/.vim/pack/mypackage/start/"
    opt_path = home_path + "/.vim/pack/mypackage/opt/"

    if os.path.isdir(start_path):
        pass
    else:
        os.makedirs(start_path)

    if os.path.isdir(opt_path):
        pass
    else:
        os.makedirs(opt_path)

    print('OK!')

if __name__ == '__main__':
    start()

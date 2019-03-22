import os
import subprocess
import json
import shutil
import sys

START_DIR = '/.vim/pack/mypackage/start/'
OPT_DIR = '/.vim/pack/mypackage/opt/'
HOME_PATH = os.path.expanduser('~')


def main(arg):
    start_package_list = []
    opt_package_list = []
    f = open('package.json', 'r')
    tmp_list = json.load(f)

    for start in tmp_list["start"]:
        start_package_list.append(start)

    for opt in tmp_list["opt"]:
        opt_package_list.append(opt)

    if len(arg) == 0:
        check(start_package_list, HOME_PATH, START_DIR, tmp_list)
        check(opt_package_list, HOME_PATH, OPT_DIR, tmp_list)
    else:
        if arg[0] == 'update':
            for package in start_package_list:
                print('CHECK ' + package.split('/')[1])
                subprocess.run(['git', '-C', HOME_PATH+START_DIR + package.split('/')[1], 'pull'])
        else:
            print('no argument')


def check(package_list, HOME_PATH, d, tmp_list):
    path = HOME_PATH + d
    check_package = []
    now_package = os.listdir(path)
    diff = []
    add_diff = []
    add_list = []
    remove_list = []

    if len(now_package) == 0:
        for package in package_list:
            add_list.append(package.split('/'))
    else:
        for d in package_list:
            diff.append(d.split('/')[1])
        add_tmp = set(diff) - set(now_package)
        add_diff = list(add_tmp)
        for l in tmp_list["start"]:
            for ad in add_diff:
                if ad in l:
                    add_list.append(l.split('/'))

        for plug in package_list:
            check_package.append(plug.split('/')[1])
        rm_tmp = set(now_package) - set(check_package)
        remove_list = list(rm_tmp)

    addfunc(add_list, HOME_PATH, dir)

    removefunc(remove_list, HOME_PATH, dir)


def addfunc(add_list, HOME_PATH, dir):
    for package in add_list:
        git_url = 'https://github.com/'
        repo = git_url + '/'.join(package) + '.git'
        path = HOME_PATH + d + package[1]
        subprocess.run(['git', 'clone', repo, path])


def removefunc(remove_list, HOME_PATH, d):
    for plug in remove_list:
        path = HOME_PATH + d + plug
        shutil.rmtree(path)
        print('remove ' + plug)


if __name__ == '__main__':
    main(sys.argv[1:])

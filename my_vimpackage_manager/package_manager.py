import os
import subprocess
import json
import shutil
import sys

def main(arg):
    start_package_list = []
    opt_package_list = []
    start_dir = '/.vim/pack/mypackage/start/'
    opt_dir = '/.vim/pack/mypackage/opt/'
    home_path = os.path.expanduser('~')
    f = open('package.json','r')
    tmp_list = json.load(f)

    for start in tmp_list["start"]:
        start_package_list.append(start)

    for opt in tmp_list["opt"]:
        opt_package_list.append(opt)

    if len(arg) == 0:
        check(start_package_list,home_path,start_dir,tmp_list)
        check(opt_package_list,home_path,opt_dir,tmp_list)
    else:
        # アップデートの処理
        if arg[0] == 'update':
            for package in start_package_list:
                print('CHECK ' + package.split('/')[1])
                subprocess.run(['git','-C',home_path+start_dir + package.split('/')[1],'pull'])
        else:
            print('no argument')

def check(package_list,home_path,dir,tmp_list):
    path = home_path + dir
    check_package = []
    now_package = os.listdir(path)
    diff = []
    add_diff = []
    add_list = []
    remove_list = []

    # 導入済みのパッケージがない
    if len(now_package) == 0:
        for package in package_list:
            add_list.append(package.split('/'))
    else:
        # package_list.jsonにあるパッケージを導入
        for d in package_list:
            diff.append(d.split('/')[1])
        add_tmp = set(diff) - set(now_package)
        add_diff = list(add_tmp)
        for l in tmp_list["start"]:
            for ad in add_diff:
                if ad in l:
                    add_list.append(l.split('/'))

        # package_list.jsonにないパッケージを削除
        for plug in package_list:
            check_package.append(plug.split('/')[1])
        rm_tmp = set(now_package) - set(check_package)
        remove_list = list(rm_tmp)

    # パッケージ追加処理
    addfunc(add_list,home_path,dir)

    # パッケージ削除処理
    removefunc(remove_list,home_path,dir)

def addfunc(add_list,home_path,dir):
    for package in add_list:
        git_url = 'https://github.com/'
        repo = git_url + '/'.join(package) + '.git'
        path = home_path + dir + package[1]
        subprocess.run(['git','clone',repo,path])

def removefunc(remove_list,home_path,dir):
    for plug in remove_list:
        path = home_path + dir + plug
        shutil.rmtree(path)
        print('remove ' + plug)

if __name__ == '__main__':
    main(sys.argv[1:])

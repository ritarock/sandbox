import os
import subprocess
import json
import shutil

def main():
    start_plugin_list = []
    opt_plugin_list = []
    start_dir = '/.vim/pack/mypackage/start/'
    opt_dir = '/.vim/pack/mypackage/opt/myplugin/'
    home_path = os.path.expanduser('~')
    f = open('plugin.json','r')
    tmp_list = json.load(f)
    for start in tmp_list["start"]:
        start_plugin_list.append(start)
    for opt in tmp_list["opt"]:
        opt_plugin_list.append(opt)
    check(start_plugin_list,home_path,start_dir,tmp_list)
    check(opt_plugin_list,home_path,opt_dir,tmp_list)

def check(plugin_list,home_path,dir,tmp_list):
    path = home_path + dir
    check_plugin = []
    now_plugin = os.listdir(path)
    diff = []
    add_diff = []
    add_list = []
    remove_list = []

    # 導入済みのプラグインがない
    if len(now_plugin) == 0:
        for plugin in plugin_list:
            add_list.append(plugin.split('/'))
    else:
        # plugin_list.josにあるプラグインを導入
        for d in plugin_list:
            diff.append(d.split('/')[1])
        add_tmp = set(diff) - set(now_plugin)
        add_diff = list(add_tmp)
        for l in tmp_list["start"]:
            for ad in add_diff:
                if ad in l:
                    add_list.append(l.split('/'))

        # plugin_list.josにないプラグインを削除
        for plug in plugin_list:
            check_plugin.append(plug.split('/')[1])
        rm_tmp = set(now_plugin) - set(check_plugin)
        remove_list = list(rm_tmp)

    addfunc(add_list,home_path,dir)
    removefunc(remove_list,home_path,dir)

def addfunc(add_list,home_path,dir):
    for plugin in add_list:
        git_url = 'https://github.com/'
        repo = git_url + '/'.join(plugin) + '.git'
        path = home_path + dir + plugin[1]
        subprocess.run(['git','clone',repo,path])

def removefunc(remove_list,home_path,dir):
    for plug in remove_list:
        path = home_path + dir + plug
        shutil.rmtree(path)

if __name__ == '__main__':
    main()

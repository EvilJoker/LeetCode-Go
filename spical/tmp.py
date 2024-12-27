#!/usr/local/bin/env python3

import os
import shutil
import subprocess
import yaml

# 读取文件
specs={}
with open("./top150.md") as fin:
    lines = fin.readlines()
    items = []
    catagory = ""
    count = 1
    for line in lines:
        line = line.strip().replace(" ", "")
        if '\n' == line or line == None or line == "":
            continue
        if line[0].isdigit():
            items.append(line)
        elif line[0] == '中' or line[0] == '困' or line[0] == '简':
            items.append(items.pop()+ "__" + line)
        else:
            if catagory == "":
                catagory = str(count).zfill(2) +line
        
                count +=1
            else:
                specs[catagory] = items
                items = []
                catagory = str(count).zfill(2) +line
        
                count +=1

yaml.dump(specs, open("./top150.yaml", "w", encoding="utf-8"), allow_unicode=True, default_flow_style=False)

# 转化要创建的信息
label_format = "catagory={};status=03plan;recommend=01basic;difficulty={};lastupdate=2024-12-19;alias=none;"
name_format = "{}_{}"
dirs_lc = []
for _, dirname, _ in os.walk("../leetcode"):
    dirs_lc.extend(dirname)

shutil.rmtree("../tmp")
os.mkdir("../tmp")


for catagory, items in specs.items():
    count = 0 
    for item in items:
        if "__" not in item:
            continue
        count +=1
        # 解析信息
        item_name = item.split("__")[0]
        item_diff = item.split("__")[1]
        order = item_name.split(".")[0]
        pure_name = item_name.split(".")[1]
        
        difficulty = "01easy"
        if item_diff.find("中等") != -1:
            difficulty = "02mid"
        elif item_diff.find("困难") != -1:
            difficulty = "03hard"
        
        targent_dir_name = name_format.format(catagory, str(count).zfill(2) + pure_name)
        label = label_format.format(catagory, difficulty, item_name)
        
        # 创建目录生成文件
        order4 = str(order).zfill(4)
        dirname_lc = ""
        for dir_lc in dirs_lc:
            if dir_lc.startswith(order4):
                dirname_lc = dir_lc
                break
        if dirname_lc == "":
            continue
        
        ## 创建目录
        dir_name_path = os.path.join("../tmp", targent_dir_name)
        os.mkdir(dir_name_path)

        ## copy 
        subprocess.run("cp -rf ../leetcode/{}/* {}/".format(dirname_lc, dir_name_path.replace(")", "\)").replace("(", "\(")), shell=True)
        
        ## 创建 label
        label_path = os.path.join(dir_name_path, "label.txt")
        with open(label_path, "w") as fout:
            fout.write(label)
            
        ## 创建 run.sh
        sh_path = os.path.join(dir_name_path, "run.sh")
        with open(sh_path, "w") as fout:
            fout.write("""current_path=$(pwd)

cd ../

ln -s $current_path tmp

cd tmp

go test -v ./...

cd ../

rm -rf tmp
            """
            )      
            
        
        



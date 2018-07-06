# wncdu

wncdu是对ncdu工具的一个封装，主要目的是找出指定目录下最大的10个文件

## 编译
```bash
cd /path/to/source
go build -o wncdu
```

##  用法
```bash
Usage: wncdu dir [options]
  -v, --version               Print help message and quit.
  -h, --help                  Print version and quit.
  -t, --timeout int           Timeout for ncdu. (default 30)
  -x, --cross                 Do not cross filesystem boundaries, i.e. only count files and directories on the same filesystem as the directory being scanned.
  -e, --exclude string        Exclude files that match PATTERN.
  -X, --exclude-from string   Exclude files that match any pattern in FILE. Patterns should be separated by a newline.
```

## 结果

```bash
Progress: 100%
path                                                                                                                    size      inode
/home/gaoguangpeng/go/src/github.com/coreos/etcd/.git/objects/pack/pack-1a297a052343807581606db23613539762d87c5e.pack   37 MB     395362
/home/gaoguangpeng/go/src/github.com/gogo/protobuf/.git/objects/pack/pack-c15b7615f35861086e07117b854c4bd8acc0799b.pack 29 MB     527071
/home/gaoguangpeng/go/src/wonder/.git/objects/pack/pack-c9d5da4cb18f0cb3bbe4a5cd60a0cffafda31cce.pack                   26 MB     528455
/home/gaoguangpeng/go/src/wonder/.git/objects/pack/pack-303aa9d5dd0763b089e1cd6076d5ced5590969b8.pack                   16 MB     530581
/home/gaoguangpeng/go/src/ncdu/a                                                                                        10 MB     396999
/home/gaoguangpeng/go/src/wonder/vendor/github.com/pierrec/lz4/testdata/207326ba-36f8-11e7-954a-aca46ba8ca73.png        10 MB     655891
/home/gaoguangpeng/go/src/wonder/qbus/libQBus_go.so                                                                     7.9 MB    528977
/home/gaoguangpeng/go/src/wonder/.git/objects/3a/31e314eebb830f2912b2f6792529f554f48c06                                 7.2 MB    789335
/home/gaoguangpeng/go/src/latte/vendor/golang.org/x/text/collate/tables.go                                              5.0 MB    396482
/home/ops/ops-agent/data_agent/log/acquisition.0                                                                        4.7 MB    58251
```
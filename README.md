# 字符串模糊匹配
* 在语音转文字后，提供几个简单的服务来比较文字的差别

现在提供的服务包括
* 直接进行比较， 完全一样输出0，不同输出1
* 编辑距离（Edit distance）, 0为完全一样，越大区别越大
* Jaccard Similarity（中文切词后）, 1为完全一样，0为完全不同
* SimHash: 哈希后比对相似度， 1为完全一样，0为完全不同

# 使用方式
## Setup
安装golang 1.13, 并且安装依赖
```bash
go mod tidy
```
编译并运行这个服务
```bash
# compile
./build.sh 
# run
./output/bootstrap.sh 
```

## 使用方式
router

`POST http://127.0.0.1:8080/api/v1/:method`

`:method`可以被替换成
* `base`
* `edit_distance`
* `jaccard_distance`
* `hamming_distance` (使用的simhash的方式做的计算)

Request Body
```
{
    "source":"哈哈哈，今天真好",
    "target":"哈哈，天气不错"
}
```
Response
```
// if success
{
    "message": "success",
    "score": 0.14285714285714285
}
// if fail
{
    "error": "[error message]"
}
```

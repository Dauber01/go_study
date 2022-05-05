Go Packager管理发展史
1. GOPATH模式
    GOPATH目录是所有工程的公共依赖包目录,所有需要编译的go工程的依赖包都放在GOPATH目录下
2. Vendor特性
    为了解决GOPATH模式下,多个工程需要共享GOPATH目录,无法适应各个工程对于不同版本的依赖包的使用,不便于更新某个依赖包.go 1.6之后开启了vendor目录
3. Go Moduler包管理
    从go1.11以后开始支持module依赖管理工具,从而实现了依赖包的进行升级更新,在Go 1.13版本之后默认打开

Go Moduler
    go.mod文件中会使用到的语法关键词及含义
1. module: 定义当前项目的模块路径,不再赘述
2. go: 标定当前模块的Go语言版本,目前看来还只是个标示的作用
3. require: 说明Module需要什么版本的依赖
4. exclude: 用于从使用中排除一个特定的模块版本
5. replace: 替换require中声明的依赖,使用另外的依赖及版本号

Go Modules Checksum
    Go开发团队在引入go.mod的同时引入了go.sum文件,用于记录每个依赖包的哈希值,在构建时,如果本地的依赖包的hash值与go.sum文件中记录的不一致,则会拒绝构建
1. go.sum文件中的每行记录由module名,版本和哈希组成,并由空格分开
2. 在引入新的依赖时,通常会使用go get命令获取该依赖,并下载至本地缓存目录 $GOPATH/pkg/mod/cache/download ,该依赖包为一个后缀为.zip的压缩包,并且把hash运算同步到go.sum文件中
3. 在构建应用时,会从本地缓存中查找所有go.mod中记录的依赖包,并计算本地依赖包的hash值,然后与go.sum中的记录进行对比,当校验失败,go命令将拒绝构建

Go Modules Proxy
    Go1.13将GOPROXY默认成大陆无法访问的 https://proxy.golang.org ,所以在国内需配置代理进行使用
GOPROXY可以解决一些公司内部的使用问题
1. 访问公司内网的git server
2. 防止公网仓库变更或消失,导致线上编译失败或者紧急回退失败
3. 公司审计和安全需要
4. 防止公司内部开发人员配置不当造成 import path 泄漏
5. cache 热点依赖,降低公司公网出口带宽

# 配置 GOPROXY 环境变量,其中 direct 表示可以回源
export GOPROXY=https://goproxy.io,direct
# 还可以设置不走 proxy 的私有仓库或组,用多个逗号相隔(可选)
export GOPRIVATE=git.mycompany.com,github.com/my/private

配置 GOPRIVATE 之后,涉及到的包会跳过 GOPROXY 和包的检查
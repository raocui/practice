# practice
go api接口端

第一次clone带子模块的项目
git clone --recurse-submodules url

如果第一次没有clone子模块，下载完成后，但是子模块并没有下载下来，使用下面的命令
git submodule update --init --recursive

过了一段时间，子模块更新了，需要重新更新子模块，一个命令解决
git submodule foreach git pull origin master

子模块更新
git submodule update --remote

更新所有子模块
git submodule foreach git submodule update

如果已经克隆了项目并忘记了 ，则可以通过运行 来组合 和 步骤。为了还可以初始化、提取和签出任何嵌套子模块，可以使用万无一失的子模块。--recurse-submodules git submodule init git submodule updategit submodule update --initgit submodule update --init --recursive


为没有子模块的项目，添加子模块或者把子目录换成子模块
git submodule add <submodule url>

如果项目报错，提示 'common' already exists in the index，说明common属于该项目的子目录，需要移除
git rm -r --cached common
移除后，再手动将common改个别名，然后再重新添加子模块的 引用


git submodule init
git submodule update --remote
git submodule update --init --recursive
git clone --recurse-submodules <submodule url>



===========================================
go get github.com/ppkg/xxl-job-executor-go@85e9922ab9520abe8a27990ecafc2662206fe510


Windows客户端启动举例：

1.设置host
172.16.2.29			exadmin.exchain.com
172.16.2.30			service.exchain.com
#127.0.0.1			f.exadmin.exchain.com
172.16.2.29			f.exadmin.exchain.com

2.解压goadmin.tar.gz选择不要替换文件，运行goadmin.exe

3.浏览器访问http://f.exadmin.exchain.com








——————————————————————————————————————————————————————————————————————

升级脚本举例：

pkill goadmin
cd /data/server/FinanceAdmin
git pull

tar -zvxf goadmin.tar.gz
rm -rf views/login.tpl
git checkout -- views/login.tpl
setsid ./goadmin >> logs &


















Linux服务端部署举例：
1.cd /data/server; git clone xxx
1.设置nginx，参考nginx.conf配置文件
2.配置Host指向：172.16.2.37 exadmin.exchain.com
2.启动setsid ./goadmin >> logs/log &

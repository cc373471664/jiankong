基于GO+Redis的高性能运行状态监控
====
定时发起HTTP请求去判断web站点是否存活,如果挂掉发送企业微信通知
## 架构
    前端:Vue+Ant Design+Axios
    后端:Go1.13.3
    存储:Redis5
## v1版本功能:
    1.支持自定义请求
    2.可查看进程日志情况
    3.可配置进程访问间隔
    4.配置企业微信推送
## 流程
   **1.首页可查看所有接口运行状态**
   ![avatar](https://github.com/cc373471664/jiankong/blob/master/md/shouye.png)
   **2.点击[添加计划]进行任务添加**
   ![avatar](https://github.com/cc373471664/jiankong/blob/master/md/tianjia.png)
   **3.运行时可通过[日志]查询历史监控记录**
   ![avatar](https://github.com/cc373471664/jiankong/blob/master/md/rizhi.png)
   **4.如果配置了[企业微信]会进行推送**
   ![avatar](https://github.com/cc373471664/jiankong/blob/master/md/qiyeweixin.png)
## v2版本计划
    1.短信的推送提醒
    2.定期清理日志
    3.添加数据库状态监控
    4.性能健壮监控

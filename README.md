# 获取真实的url跳转地址
网址的跳转分两种  
+ 后端 30x 跳转
    如: http://www.weobo.com/  
    301 跳转-> https://www.weibo.com/  
    302 跳转-> https://weibo.com/newlogin?url=https%3A%2F%2Fwww.weibo.com%2F  

+ 前端js跳转
  如: http://60.190.82.94:8083/  
  window.location.href = '/MAIN/index_sx.html'

30x 的跳转很好搞定,但是js的跳转很难搞.此项目是为了应对js的跳转

# 使用:  
在input.txt中输入需要确定的链接,会输出到output.txt中
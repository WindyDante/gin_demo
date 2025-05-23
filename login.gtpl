<html>
<head>
<title></title>
</head>
<body>
// 这里的username=astaxie属于是多携带了一个参数,如果此时在表单中输入username
// 则参数会一起带过来,给服务器多个username,做切片返回
<form action="/login?username=astaxie" method="post">
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    <input type="submit" value="登录">
</form>
</body>
</html>
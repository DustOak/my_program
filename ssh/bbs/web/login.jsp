<%--
  Created by IntelliJ IDEA.
  User: cuiwenbin
  Date: 19-4-16
  Time: 下午7:12
  To change this template use File | Settings | File Templates.
--%>
<!doctype html>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="s" uri="/struts-tags" %>
<html lang="en">
<head>
    <title>用户登录</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/4.0.0/css/bootstrap.min.css"
          integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <link rel="stylesheet" href="css/main.css">
    <script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.4.0/jquery.min.js"></script>
    <script type="text/javascript" src="js/main.js"></script>

</head>
<body id="login">
<div class="container">
    <div class="row">
        <div class="col-sm">
        </div>
        <div class="col-sm">
            <div class="loginBox text-center">
                <h4>校园BBS系统</h4>
                <hr>
                <form action="loginAction" method="post" target="_self" >
                    <s:fielderror class="p-3 mb-2 bg-danger text-white "
                           name="error"      style="border-radius: 3px;margin: 0 auto;list-style-type:none"/>
                    <div class="form-group">
                        <input class="form-control" id="account" type="text" placeholder="User Name" name="account">
                    </div>
                    <div class="form-group">
                        <input class="form-control" id="password" type="password" placeholder="Password" name="password">
                    </div>
                    <div class="form-check form-check-inline ">
                        <input class="form-check-input" checked type="radio" name="adminOrStudent" id="inlineRadio1" value="0">
                        <label class="form-check-label" for="inlineRadio1">学生</label>
                    </div>
                    <div class="form-check form-check-inline">
                        <input class="form-check-input" type="radio" name="adminOrStudent" id="inlineRadio2" value="1">
                        <label class="form-check-label" for="inlineRadio2">管理员</label>
                    </div>
                    <div class="form-group" style="margin-top: 5%">
                        <button class="btn btn-block btn-info" type="submit">Login</button>
                    </div>
                </form>
            </div>
        </div>
        <div class="col-sm">
        </div>
    </div>
</div>
</body>
</html>

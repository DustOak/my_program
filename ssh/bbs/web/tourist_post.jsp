<%--
Created by IntelliJ IDEA.
User: cuiwenbin
Date: 19-5-10
Time: 下午9:14
To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="s" uri="/struts-tags" %>
<html>
<head>
    <title>梦想科大</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://cdn.staticfile.org/twitter-bootstrap/4.3.1/css/bootstrap.min.css">
    <link rel="stylesheet" href="css/main.css">
    <script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.4.0/jquery.min.js"></script>
</head>
<body>
<!--head bar-->
<nav class="navbar navbar-expand-lg navbar-light bg-light " style="position:fixed;width:100%;z-index:99;">
    <a class="navbar-brand" href="#" style="margin-right: 10%">梦想科大</a>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <form class="form-inline my-2 my-lg-0">
            <input class="form-control mr-sm-2" type="search" placeholder="帖子名" aria-label="Search">
            <button class="btn btn-outline-success my-2 my-sm-0" type="submit">查找</button>
        </form>
    </div>
</nav>
<!--left bar-->
<div class="float-left" style="top: 6%;background: #293134;width: 12%;color: white; position:fixed;height: 100%;">
    <div class="card-img text-center" style="margin-bottom: 5%">
        <img src="images/head/none.png" style="margin-top:16%;border-radius: 100px; width: 100px;">
    </div>
    <div class="text-center">
        <p>当前为游客状态<br/>无法进行发/回帖操作</p>
        <p>请<a href="/login">登录</a></p>
    </div>
</div>


<!-- hot post bar-->
<div class="float-right" style="margin-top: 4%;
width:100%;color: black; overflow:hidden;padding-left: 17%;">
    <div style="width: 90%;">
        <div>
            <div class="alert alert-info" role="alert"
                 style="width: 100%;border-radius: 5px 5px 0 0;color: black;padding:.4% .4% .4% .9%;">
                <a href="index" style="color: black">主页</a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;
                <a href="posts?board=<s:property
        value="postData.getBid().getId()"/>" style="color: black"><s:property
                        value="postData.getBid().getName()"/></a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;文章:<s:property
                    value="postData.getName()"/>
            </div>
            <table class="table table-bordered">
                <tbody>
                <tr>
                    <td width="23%">
                        <s:if test="postData.getSid()!=null">
                            <div style="background: white;margin: 4% auto 0 auto; width:50%;border: #CCC solid 1px"
                                 class="text-center">
                                <s:if test="postData.getSid().getPhotoPath()!=null">
                                    <img src="images/head/<s:property value="postData.getSid().getPhotoPath()"/>"
                                         width="150" height="150">
                                </s:if>
                                <s:else>
                                    <img src="images/head/noneHead.jpg"
                                         width="150" height="150">
                                </s:else>
                            </div>
                            <div style=" margin-top: 0;" class="text-left">

                                <p style="margin-left: 25%;margin-bottom: .5rem"><s:property
                                        value="postData.getSid().getNickName()"/></p>
                                <p style="margin-left: 25%;margin-bottom: 0;color:#888 ">QQ:<s:property
                                        value="postData.getSid().getQq()"/></p>
                                <p style="margin-left: 25%;margin-bottom: 0;color:#888">Email:<s:property
                                        value="postData.getSid().getEmail()"/></p>
                            </div>
                        </s:if>
                        <s:else>
                            <div style="background: white;margin: 4% auto 0 auto; width:50%;border: #CCC solid 1px"
                                 class="text-center">
                                <s:if test="postData.getAid().getPhotoPath()!=null">
                                    <img src="images/head/<s:property value="postData.getAid().getPhotoPath()"/>"
                                         width="150" height="150">
                                </s:if>
                                <s:else>
                                    <img src="images/head/noneHead.jpg"
                                         width="150" height="150">
                                </s:else>
                            </div>
                            <!--此处html代码有问题 需要改成执行html代码-->
                            <div style=" margin-top: 0;" class="text-left">
                                <p style="margin-left: 25%;margin-bottom: .5rem;color: red"><s:property
                                        value="postData.getAid().getNickname()"/></p>
                            </div>
                        </s:else>
                    </td>
                    <td width="76%">
                        <s:property
                                value="%{postData.getContent()}"/>
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>
</div>
</body>
</html>

<%--
Created by IntelliJ IDEA.
User: cuiwenbin
Date: 19-5-10
Time: 下午9:14
To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="s" uri="/struts-tags" %>
<%@ taglib prefix="spring" uri="http://www.springframework.org/tags" %>
<html>
<head>
    <title><s:property
            value="postData.getName()"/></title>
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
                 style="width: 100%;border-radius: 5px 5px 0 0;color: black;padding:.4% .4% 0 .9%;height: 6%;">
                <a href="index" style="color: black">主页</a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;
                <a href="posts?board=<s:property
        value=" postData.getBid().getId()"/>" style="color: black">
                    <s:property
                            value="postData.getBid().getName()"/>
                </a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;文章:
                <s:property
                        value="postData.getName()"/>
                <span class="d-inline-block " style="float: right;margin-right: 8%" tabindex="0" data-toggle="tooltip"
                      title="游客模式,无法回帖">
  <button class="btn btn-secondary" style="pointer-events: none;" type="button" disabled>回帖</button>
                </span>
            </div>
            <div class="card-group" style="margin-bottom: 1%;border-radius: 3px">
                <div class="card" style="width: 20%; flex:none; ">
                    <div class="alert alert-info" role="alert"
                         style="margin-bottom: 0;border-radius: 3px 0 0 0 ;border-right: none;height: 30px;padding: 1% 0 0 5%;">
                        楼主
                    </div>
                    <s:if test="postData.getSid()!=null">
                        <div class="card-body" style="padding-left: 23%;">
                            <s:if test="postData.getSid().getPhotoPath()!=null">
                                <img src="images/head/<s:property value="postData.getSid().getPhotoPath()"/>"
                                     width="150" height="150">
                            </s:if>
                            <s:else>
                                <img src="images/head/noneHead.jpg" width="150" height="150">
                            </s:else>
                            <p><s:property
                                    value="postData.getSid().getNickName()"/></p>
                            <p style="margin-bottom: 0; color: #999">QQ:<s:property
                                    value="postData.getSid().getQq()"/></p>
                            <p style="margin-bottom: 0;color: #999">Email:<s:property
                                    value="postData.getSid().getEmail()"/></p>
                        </div>
                    </s:if>
                    <s:else>
                        <div class="card-body" style="padding-left: 23%;">
                            <s:if test="postData.getAid().getPhotoPath()!=null">
                                <img src="images/head/<s:property value="postData.getAid().getPhotoPath()"/>"
                                     width="150" height="150">
                            </s:if>
                            <s:else>
                                <img src="images/head/noneHead.jpg" width="150" height="150">
                            </s:else>
                            <p style="color: red;margin-bottom: 0"><s:property
                                    value="postData.getAid().getNickname()"/></p>
                            <p style="color: red">管理员发帖</p>
                        </div>
                    </s:else>
                </div>
                <div class="card">
                    <div class="alert alert-info" role="alert"
                         style="margin-bottom: 0;border-radius: 3px 0 0 0 ;border-right: none;height: 30px;padding: .2% 0 1% 1%;">
                        <s:property
                            value="postData.getPublishTime()"/>
                    </div>
                    <div class="card-body"><s:property
                            value="postData.getContent()" escapeHtml="false"/>
                    </div>
                </div>
            </div>
            <s:iterator value="postData.getReplies()" var="reply" status="index">
                <div class="card-group" style="margin-bottom: 1%;border-radius: 3px">
                    <div class="card" style="width: 20%; flex:none; ">
                        <div class="alert alert-info" role="alert"
                             style="margin-bottom: 0;border-radius: 3px 0 0 0 ;border-right: none;height: 30px;padding: 1% 0 0 5%;">
                            第<s:property value="#index.getIndex()+1"></s:property>楼
                        </div>
                        <s:if test="#reply.getSid()!=null">
                            <div class="card-body" style="padding-left: 23%;">
                                <s:if test="postData.getSid().getPhotoPath()!=null">
                                    <img src="images/head/<s:property value="#reply.getSid().getPhotoPath()"/>"
                                         width="150" height="150">
                                </s:if>
                                <s:else>
                                    <img src="images/head/noneHead.jpg" width="150" height="150">
                                </s:else>
                                <p><s:property
                                        value="#reply.getSid().getNickName()"/></p>
                                <p style="margin-bottom: 0; color: #999">QQ:<s:property
                                        value="#reply.getSid().getQq()"/></p>
                                <p style="margin-bottom: 0;color: #999">Email:<s:property
                                        value="#reply.getSid().getEmail()"/></p>
                            </div>
                        </s:if>
                        <s:else>
                            <div class="card-body" style="padding-left: 23%;">
                                <s:if test="#reply.getAid().getPhotoPath()!=null">
                                    <img src="images/head/<s:property value="#reply.getAid().getPhotoPath()"/>"
                                         width="150" height="150">
                                </s:if>
                                <s:else>
                                    <img src="images/head/noneHead.jpg" width="150" height="150">
                                </s:else>
                                <p style="color: red;margin-bottom: 0"><s:property
                                        value="#reply.getAid().getNickname()"/></p>
                                <p style="color: red">管理员发帖</p>
                            </div>
                        </s:else>
                    </div>
                    <div class="card">
                        <div class="alert alert-info" role="alert"
                             style="margin-bottom: 0;border-radius: 0 3px 0 0;height: 30px;border-left: none;padding: .3% 0 0 1%;">
                            <s:property
                                value="#reply.getPublishTime()"/>
                            <a href="" id="back-to-top" onclick="backTop()"> 返回顶部</a>
                        </div>
                        <div class="card-body"><s:property
                                value="#reply.getContent()" escapeHtml="false"/>
                        </div>
                    </div>
                </div>
            </s:iterator>
        </div>
    </div>

</div>
</body>
<script>
    var timer = null;

    function backTop() {
        cancelAnimationFrame(timer);
        //获取当前毫秒数
        var startTime = +new Date();
        //获取当前页面的滚动高度
        var b = document.body.scrollTop || document.documentElement.scrollTop;
        var d = 500;
        var c = b;
        timer = requestAnimationFrame(function func() {
            var t = d - Math.max(0, startTime - (+new Date()) + d);
            document.documentElement.scrollTop = document.body.scrollTop = t * (-c) / d + b;
            timer = requestAnimationFrame(func);
            if (t == d) {
                cancelAnimationFrame(timer);
            }
        });
    }
</script>
</html>

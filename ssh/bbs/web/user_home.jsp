<%--
Created by IntelliJ IDEA.
User: cuiwenbin
Date: 19-4-16
Time: 下午9:32
To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="s" uri="/struts-tags" %>
<!DOCTYPE html>
<html lang="zh-CN" style="height: 100%;padding: 0; margin: 0;">
<head>
    <title>梦想科大</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/4.0.0/css/bootstrap.min.css"
          integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <link rel="stylesheet" href="css/main.css">
    <script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.4.0/jquery.min.js"></script>
</head>
<body style="background: rgb(236, 244, 253);overflow-y: scroll;height: auto;padding: 0; margin: 0;">
<!--head bar-->
<nav class="navbar navbar-expand-lg navbar-light bg-light " style="position:fixed;width:100%;z-index:99;">
    <a class="navbar-brand" href="#" style="margin-right: 10%">梦想科大</a>

</nav>
<!--left bar-->
<div class="float-left" style="top: 6%;background: #293134;width: 12%;color: white; position:fixed;height: 100%;">
    <s:if test="student!=null">
        <div class="card-img text-center" style="margin-bottom: 5%">
            <s:if test="student.getPhotoPath()!=null">
                <img src="images/head/<s:property value="student.getPhotoPath()"/>"
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:if>
            <s:else>
                <img src="images/head/noneHead.jpg"
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:else>

        </div>
        <div class="text-center">
            <p>昵称:<s:property value="student.getNickName()"/></p>
            <p>姓名:<s:property value="student.getRealName()"/></p>
            <p><a class="btn btn-info" href="/myInfo?sessionId=<s:property value="sessionId"/>">个人信息</a></p>
            <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value="sessionId"/>">查看我的帖子</a></p>
            <p><a class="btn btn-info" href="/myReplies?sessionId=<s:property value="sessionId"/>">查看我的回复</a></p>
            <p><a class="btn btn-warning" href="/logout?sessionId=<s:property value="sessionId"/>">注销</a></p>
        </div>
    </s:if>
    <s:else>
        <div class="card-img text-center" style="margin-bottom: 5%">
            <s:if test="admin.getPhotoPath()!=null">
                <img src="images/head/<s:property value="admin.getPhotoPath()"/> "
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:if>
            <s:else>
                <img src="images/head/noneHead.jpg"
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:else>
        </div>
        <div class="text-center">
            <p>昵称:<s:property value="admin.getNickname()"/></p>
            <p>姓名:<s:property value="admin.getName()"/></p>
            <p><a class="btn btn-info" href="/myInfo?sessionId=<s:property value="sessionId"/>">个人信息</a></p>
            <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value="sessionId"/>">查看我的帖子</a></p>
            <p><a class="btn btn-info" href="/myReplies?sessionId=<s:property value="sessionId"/>">查看我的回复</a></p>
            <p><a class="btn btn-info" href="/boardManager?sessionId=<s:property value="sessionId"/>">板块操作</a></p>
            <p><a class="btn btn-warning" href="/logout?sessionId=<s:property value="sessionId"/>">注销</a></p>
        </div>
    </s:else>
</div>

<!--写到这了!!!-->
<!-- hot post bar-->
<div class="float-right" style="margin-top: 4%;margin-left: .5%;
width:87.5%;height:auto;color: white; overflow-y: auto;">
    <div style="background: white;width: 80%;height:auto;margin: 2% auto;">
        <div id="carouselExampleControls" class="carousel slide "
             data-ride="carousel" style="width: 40%;">
            <div class="carousel-inner">
                <div class="carousel-item">
                    <img class="d-block w-100" src="images/imgs/p1.jpg" alt="Second slide">
                </div>
                <div class="carousel-item">
                    <img class="d-block w-100" src="images/imgs/p3.jpg" alt="Third slide">
                </div>
                <div class="carousel-item">
                    <img class="d-block w-100" src="images/imgs/p4.jpg" alt="Third slide">
                </div>
                <div class="carousel-item active">
                    <img class="d-block w-100" src="images/imgs/p5.jpg" alt="First slide">
                </div>
            </div>
            <a class="carousel-control-prev" href="#carouselExampleControls" role="button" data-slide="prev">
                <span class="carousel-control-prev-icon" aria-hidden="true"></span>
                <span class="sr-only">Previous</span>
            </a>
            <a class="carousel-control-next" href="#carouselExampleControls" role="button" data-slide="next">
                <span class="carousel-control-next-icon" aria-hidden="true"></span>
                <span class="sr-only">Next</span>
            </a>
        </div>
        <div style="width: 35%;height: 28%;color: black;position: absolute;
        margin-left: 31%;top: 14%;border: #9fcdff solid 1px;padding: .5%;border-radius: 5px;">
            <small style="margin: 2%;padding: 0;">热门帖子:</small>
            <hr style="margin: 1%;padding: 0;">
            <s:iterator value="hotPosts" var="post">
                <p style="margin-left: 5%;font-size: 13px">
                    <a href="/posts?sessionId=<s:property value="getSessionId()"/>&board=<s:property value=" #post.getBid().getId()"/>"
                       style="color: #1d2124;"
                       class="board">【
                        <s:property
                                value="#post.getBid().getName()"/>
                        】</a><u><a
                        href="/post?sessionId=<s:property value="getSessionId()"/>&board=<s:property value=" #post.getBid().getId()"/>&post=<s:property value="#post.getId()"/>"
                        style="color: #1d2124">
                    <s:property value="#post.getName()"/>
                </a></u>【点击量:
                    <s:property value="#post.getCount()"/>
                    】
                </p>
            </s:iterator>
        </div>
    </div>
</div>

<!--board bar-->
<div class="float-right" style="margin-top: 1%;margin-left: .5%; margin-bottom: 1%;
width:87.5%;height:auto;color: black; overflow-y: hidden;">
    <s:iterator value="boards" var="board">
        <div class="alert alert-info" role="alert" style="margin-bottom: 0;margin-top: 1%;border-radius: 5px 5px 0 0 ">
            <s:property value="#board.getName()"/>
        </div>
        <div class="row" style="max-width: 96%;padding-left: 2%">
            <s:iterator value="#board.getChildBoards()" var="child">
                <div style="width: 11%;padding: 2%;margin-left: 10%;float: left;margin-right: 0;">
                    <s:if test="#child.getBoardImg()==null">
                        <img src="images/board/none2.png" width="90" height="90"
                             style="margin: 1% 0 1% 1%">
                    </s:if>
                    <s:else>
                        <img src="images/board/<s:property value=" #child.getBoardImg()"/>" width="90" height="90"
                             style="margin: 1% 0 1% 1%">
                    </s:else>
                    <a href="/posts?sessionId=<s:property value="getSessionId()"/>&board=<s:property value=" #child.getId()"/>"
                       style="position: absolute">
                        <s:property
                                value="#child.getName()"/>
                    </a>
                    <small style="position: absolute;margin-top: 1%;width:15%;max-width:15%;overflow: hidden;text-overflow: ellipsis;white-space: nowrap;">
                        <s:property value="#child.getDescription()"/>
                    </small>
                </div>
            </s:iterator>
        </div>
    </s:iterator>
</div>
<script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.slim.min.js"
        integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN"
        crossorigin="anonymous"></script>
<script src="https://cdn.bootcss.com/popper.js/1.12.9/umd/popper.min.js"
        integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q"
        crossorigin="anonymous"></script>
<script src="https://cdn.bootcss.com/bootstrap/4.0.0/js/bootstrap.min.js"
        integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl"
        crossorigin="anonymous"></script>

</body>
</html>


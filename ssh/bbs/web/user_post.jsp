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
    <title>
        <s:property
                value="postData.getName()"/>
    </title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://cdn.staticfile.org/twitter-bootstrap/4.3.1/css/bootstrap.min.css">
    <link rel="stylesheet" href="css/main.css">
    <script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.4.0/jquery.min.js"></script>
</head>
<body>
<!--head bar-->
<nav class="navbar navbar-expand-lg navbar-light bg-light " style="position:fixed;width:100%;z-index:99;top: 0;">
    <a class="navbar-brand" href="#" style="margin-right: 10%">梦想科大</a>
</nav>
<!--left bar-->
<div class="float-left" style="top: 6%;background: #293134;width: 12%;color: white; position:fixed;height: 100%;">
    <s:if test="student!=null">
        <div class="card-img text-center" style="margin-bottom: 5%">
            <s:if test="student.getPhotoPath()!=null">
                <img src="images/head/<s:property value=" student.getPhotoPath()"/> "
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:if>
            <s:else>
                <img src="images/head/noneHead.jpg"
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:else>

        </div>
        <div class="text-center">
            <p>昵称:<s:property value="student.getNickName()"/>
            </p>
            <p>姓名:<s:property value="student.getRealName()"/>
            </p>
            <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value=" sessionId"/>">查看我的帖子</a></p>
            <p><a class="btn btn-info" href="/myReplies?sessionId=<s:property value=" sessionId"/>">查看我的回复</a></p>
            <p><a class="btn btn-warning" href="/logout?sessionId=<s:property value="sessionId"/>">注销</a></p>
        </div>
    </s:if>
    <s:else>
        <div class="card-img text-center" style="margin-bottom: 5%">
            <s:if test="admin.getPhotoPath()!=null">
                <img src="images/head/<s:property value=" admin.getPhotoPath()"/> "
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:if>
            <s:else>
                <img src="images/head/noneHead.jpg"
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:else>
        </div>
        <div class="text-center">
            <p>昵称:<s:property value="admin.getNickname()"/>
            </p>
            <p>姓名:<s:property value="admin.getName()"/>
            </p>
            <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value=" sessionId"/>">查看我的帖子</a></p>
            <p><a class="btn btn-info" href="/myReplies?sessionId=<s:property value=" sessionId"/>">查看我的回复</a></p>
            <p><a class="btn btn-info" href="/boardManager?sessionId=<s:property value="sessionId"/>">板块操作</a></p>
            <p><a class="btn btn-warning" href="/logout?sessionId=<s:property value="sessionId"/>">注销</a></p>
        </div>
    </s:else>
</div>


<!-- hot post bar-->
<div class="float-right" style="margin-top: 4%;
width:100%;color: black; overflow:hidden;padding-left: 17%;">
    <div style="width: 90%;">
        <div>
            <div class="alert alert-info" role="alert"
                 style="width: 100%;border-radius: 5px 5px 0 0;color: black;padding:.4% .4% 0 .9%;height: 6%;">
                <a href="index?sessionId=<s:property value=" getSessionId()"/>" style="color: black">主页</a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;
                <a href="posts?sessionId=<s:property value=" getSessionId()"/>&board=
                <s:property
                        value=" postData.getBid().getId()"/>
                " style="color: black">
                    <s:property
                            value="postData.getBid().getName()"/>
                </a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;文章:
                <s:property
                        value="postData.getName()"/>
                <span class="d-inline-block " style="float: right;margin-right: 8%" tabindex="0" data-toggle="tooltip"
                      title="回复帖子">
      <button class="btn btn-success" type="button" onclick="show()">回帖</button>
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
                                <img src="images/head/<s:property value=" postData.getSid().getPhotoPath()"/>"
                                     width="150" height="150">
                            </s:if>
                            <s:else>
                                <img src="images/head/noneHead.jpg" width="150" height="150">
                            </s:else>
                            <p>
                                <s:property
                                        value="postData.getSid().getNickName()"/>
                            </p>
                            <p style="margin-bottom: 0; color: #999">QQ:
                                <s:property
                                        value="postData.getSid().getQq()"/>
                            </p>
                            <p style="margin-bottom: 0;color: #999">Email:
                                <s:property
                                        value="postData.getSid().getEmail()"/>
                            </p>
                        </div>
                    </s:if>
                    <s:else>
                        <div class="card-body" style="padding-left: 23%;">
                            <s:if test="postData.getAid().getPhotoPath()!=null">
                                <img src="images/head/<s:property value=" postData.getAid().getPhotoPath()"/>"
                                     width="150" height="150">
                            </s:if>
                            <s:else>
                                <img src="images/head/noneHead.jpg" width="150" height="150">
                            </s:else>
                            <p style="color: red;margin-bottom: 0">
                                <s:property
                                        value="postData.getAid().getNickname()"/>
                            </p>
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
                    <div class="card-body">
                        <s:property
                                value="postData.getContent()" escapeHtml="false"/>
                    </div>
                </div>
            </div>
            <s:iterator value="postData.getReplies()" var="reply" status="index">
                <div class="card-group" style="margin-bottom: 1%;border-radius: 3px">
                    <div class="card" style="width: 20%; flex:none; ">
                        <div class="alert alert-info" role="alert"
                             style="margin-bottom: 0;border-radius: 3px 0 0 0 ;border-right: none;height: 30px;padding: 1% 0 0 5%;">
                            第
                            <s:property value="#index.getIndex()+1"></s:property>
                            楼
                        </div>
                        <s:if test="#reply.getSid()!=null">
                            <div class="card-body" style="padding-left: 23%;">
                                <s:if test="#reply.getSid().getPhotoPath()!=null">
                                    <img src="images/head/<s:property value="#reply.getSid().getPhotoPath()"/>"
                                         width="150" height="150">
                                </s:if>
                                <s:else>
                                    <img src="images/head/noneHead.jpg" width="150" height="150">
                                </s:else>
                                <p>
                                    <s:property
                                            value="#reply.getSid().getNickName()"/>
                                </p>
                                <p style="margin-bottom: 0; color: #999">QQ:
                                    <s:property
                                            value="#reply.getSid().getQq()"/>
                                </p>
                                <p style="margin-bottom: 0;color: #999">Email:
                                    <s:property
                                            value="#reply.getSid().getEmail()"/>
                                </p>
                            </div>
                        </s:if>
                        <s:else>
                            <div class="card-body" style="padding-left: 23%;">
                                <s:if test="#reply.getAid().getPhotoPath()!=null">
                                    <img src="images/head/<s:property value=" #reply.getAid().getPhotoPath()"/>"
                                         width="150" height="150">
                                </s:if>
                                <s:else>
                                    <img src="images/head/noneHead.jpg" width="150" height="150">
                                </s:else>
                                <p style="color: red;margin-bottom: 0">
                                    <s:property
                                            value="#reply.getAid().getNickname()"/>
                                </p>
                                <p style="color: red">管理员发帖</p>
                            </div>
                        </s:else>
                    </div>
                    <div class="card">
                        <div class="alert alert-info" role="alert"
                             style="margin-bottom: 0;border-radius: 0 3px 0 0;height: 30px;border-left: none;padding: .3% 0 0 1%;">
                            <s:property
                                    value="#reply.getPublishTime()"/>
                            <s:if test="#reply.getSid().getId()==getStudent().getId() ">
                                <a href="" id="back-to-top" onclick="backTop()"> 删除</a>
                                ●
                            </s:if>
                            <s:if test="getAdmin()!=null">
                                <a href="" id="back-to-top" onclick="backTop()"> 删除</a>
                                ●
                            </s:if>
                            <a href="" id="back-to-top" onclick="backTop()"> 返回顶部</a>
                        </div>
                        <div class="card-body">
                            <s:property
                                    value="#reply.getContent()" escapeHtml="false"/>
                        </div>
                    </div>
                </div>
            </s:iterator>
        </div>
    </div>
</div>
<!--回帖界面-->
<div style="width:50%;margin-left: 28%;margin-top: 10%;background:#f8f9fa;padding: 2%;z-index: 99;position: absolute;display: none"
     id="reply">
    <div id="textEditor" style="background: #f8f9fa"></div>
    <button class="btn btn-success" style="margin-top: 1%;width: 30%" onclick="getData()" id="submit" disabled>回复
    </button>
    <button class="btn btn-danger" style="margin-top: 1%;" onclick="clearData()">清空</button>
    <button class="btn btn-secondary" style="margin-top: 1%;" onclick="noShow()">关闭</button>
</div>
</body>
<script type="text/javascript" src="https://cdn.staticfile.org/wangEditor/10.0.13/wangEditor.min.js"></script>
<script type="text/javascript" src="https://cdn.staticfile.org/js-xss/0.3.3/xss.min.js"></script>
<script>
    var E = window.wangEditor;
    var editor2 = new E('#textEditor');
    editor2.customConfig.uploadImgShowBase64 = true;
    editor2.customConfig.onchange = function (html) {
        txt = editor2.txt.text().replace(/\s+/g, "");
        if (txt === "") {
            document.getElementById("submit").disabled = true;
        } else {
            document.getElementById("submit").removeAttribute("disabled");
        }
    };
    editor2.create();

    function show() {
        $("#reply").attr("style", "display:block;width:50%;margin-left: 28%;margin-top: 10%;background:#f8f9fa;padding: 2%;z-index: 99;position: absolute;");
    }

    function getData() {
        content = filterXSS(editor2.txt.html());
        if (content.replace(/\s+/g, "") === "") {
            alert("内容禁止为空");
        } else {
            $.ajax({
                data: {
                    "board":<s:property value="postData.getBid().getId()"/>,
                    "post":<s:property value="postData.getId()"/>,
                    "content": content,
                    "sessionId": '<s:property value="getSessionId()"/>'
                },
                url: "/saveReply",
                type: "post",
                success: function (returnData) {
                    if (returnData.error === 0) {
                        alert("回帖成功");
                        window.location.reload();
                    } else {
                        alert(returnData.msg);
                        window.location.reload();
                    }
                }
            })
        }
    }

    function noShow() {
        $('#reply').attr("style", "display:none;");
    }

    function clearData() {
        editor2.txt.html("");
    }

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

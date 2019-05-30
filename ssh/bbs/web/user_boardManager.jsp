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
        <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value="sessionId"/>">查看我的帖子</a></p>
        <p><a class="btn btn-info" href="/myReplies?sessionId=<s:property value="sessionId"/>">查看我的回复</a></p>
        <p><a class="btn btn-info" href="/index  ?sessionId=<s:property value="sessionId"/>">取消操作</a></p>
        <p><a class="btn btn-warning" href="/logout?sessionId=<s:property value="sessionId"/>">注销</a></p>
    </div>
</div>


<!-- hot post bar-->
<div class="float-right" style="margin-top: 4%;margin-left: .5%;
width:87.5%;height:auto;overflow-y: auto;">
    <div style="background: white;width: 60%;height:auto;margin: 2% auto;padding: 1%">
        <div style="border: black solid 2px ;padding: 1% ">
            <h5>添加板块</h5>
            <hr>
            <form action="addBoard" method="post" target="_self" onsubmit="return checkContent()"
                  enctype="multipart/form-data">
                <input type="hidden" value="<s:property value="sessionId"/>" name="sessionId">
                <div class="form-group">
                    <label for="exampleInputEmail1">请选择父板块</label>
                    <select class="form-control" id="exampleInputEmail1" name="parentId" style="width: 60%">
                        <option value="-1">---请选择父板块---</option>
                        <s:iterator value="getBoards()" var="parent">
                            <option value="<s:property value="#parent.getId()"/>"><s:property
                                    value="#parent.getName()"/></option>
                        </s:iterator>
                        <option value="-1">---其他---</option>
                    </select>
                    <small class="form-text text-muted">请选择父板块,如想添加父板块则选择其他
                    </small>
                </div>
                <div class="form-group">
                    <label for="boardName">新板块名称</label>
                    <input type="text" class="form-control" style="width: 60%" name="boardName"
                           id="boardName"
                           placeholder="Name">
                    <small class="form-text text-muted">新板块名称
                    </small>
                </div>
                <div class="form-group">
                    <label for="boardDescription">板块描述</label>
                    <input type="text" class="form-control" style="width: 80%" name="boardDescription"
                           id="boardDescription"
                           placeholder="Description">
                    <small class="form-text text-muted">新板块的描述信息</small>
                </div>
                <div class="form-group">
                    <label for="exampleFormControlFile1">上传图标</label>
                    <input type="file" class="form-control-file" id="exampleFormControlFile1" name="icon"
                           accept="image/*">
                </div>
                <button type="submit" class="btn btn-primary " style="width: 20%">提交</button>
                <button type="reset" class="btn btn-secondary " style="width: 20%">重置</button>
            </form>
        </div>
    </div>
</div>

<!--board bar-->
<div class="float-right" style="margin-top: 1%;margin-left: .5%; margin-bottom: 1%;
width:87.5%;height:auto;color: black; overflow-y: hidden;">
    <form action="deleteBoard?sessionId=<s:property value="sessionId"/>" onsubmit="return determineDelete()"
          method="post" target="_self">
        <s:iterator value="boards" var="board">
            <div class="alert alert-info" role="alert"
                 style="margin-bottom: 0;margin-top: 1%;border-radius: 5px 5px 0 0 ">
                <s:property value="#board.getName()"/>
                <button class="btn btn-danger btn-sm" name="board" value="<s:property value="#board.getId()"/>">删除
                </button>
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
                        <small style="position: absolute;margin-top: 3%;width:15%;max-width:15%;overflow: hidden;text-overflow: ellipsis;white-space: nowrap;">
                            <button class="btn btn-danger btn-sm" name="board" value="<s:property
                                value="#child.getId()"/>">删除
                            </button>
                        </small>
                    </div>
                </s:iterator>
            </div>
        </s:iterator>
    </form>
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
<script type="text/javascript" src="https://cdn.staticfile.org/js-xss/0.3.3/xss.min.js"></script>
<script>
    function determineDelete() {
        if (confirm("确定删除?")) {
            if (confirm("再次确定删除?删除板块会导致本板块下所有帖子和回复全部消失!")) {
                return true;
            }
        }
        return false;
    }

    function checkContent() {
        boardName = ($("#boardName").val()).replace(/\s+/g, "");
        boardDescription = ($("#boardDescription").val()).replace(/\s+/g, "");
        if (boardName !== "" && boardDescription !== "") {
            $("#boardName").val(filterXSS(boardName));
            $("#boardDescription").val(filterXSS(boardDescription));
            return true;
        }
        return false;
    }
</script>
</body>
</html>


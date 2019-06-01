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
    <link rel="stylesheet" href="https://cdn.staticfile.org/datatables/1.10.19/css/dataTables.bootstrap4.min.css">
    <script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.4.0/jquery.min.js"></script>
</head>
<body>
<!--head bar-->
<nav class="navbar navbar-expand-lg navbar-light bg-light " style="position:fixed;width:100%;z-index:99;">
    <a class="navbar-brand" href="#" style="margin-right: 10%">梦想科大</a>
</nav>
<!--left bar-->
<div class="float-left" style="top: 6%;background: #293134;width: 12%;color: white; position:fixed;height: 100%;">
    <s:if test="student!=null">
        <div class="card-img text-center" style="margin-bottom: 5%">
            <s:if test="student.getPhotoPath()!=null">
                <img src="images/head/<s:property value="student.getPhotoPath()"/> "
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:if>
            <s:else>
                <img src="images/head/noneHead.jpg"
                     style="margin-top:16%;border-radius: 100px; width: 100px;">
            </s:else>

        </div>
        <div class="text-center">
            <p>昵称:
                <s:property value="student.getNickName()"/>
            </p>
            <p>姓名:
                <s:property value="student.getRealName()"/>
            </p>
            <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value=" sessionId"/>">查看我的帖子</a></p>
            <p><a class="btn btn-info" href="/index?sessionId=<s:property value=" sessionId"/>">返回主页</a></p>
            <p><a class="btn btn-warning" href="/logout?sessionId=<s:property value=" sessionId"/>">注销</a></p>
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
            <p>昵称:
                <s:property value="admin.getNickname()"/>
            </p>
            <p>姓名:
                <s:property value="admin.getName()"/>
            </p>
            <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value=" sessionId"/>">查看我的帖子</a></p>
            <a class="btn btn-info" href="/index?sessionId=<s:property value=" sessionId"/>">返回主页</a></p>
            <p><a class="btn btn-info" href="/boardManager?sessionId=<s:property value=" sessionId"/>">板块操作</a></p>
            <p><a class="btn btn-warning" href="/logout?sessionId=<s:property value=" sessionId"/>">注销</a></p>
        </div>
    </s:else>
</div>

<!-- hot post bar-->
<div class="float-right" style="margin-top: 4%;
    width:100%;color: black; overflow:hidden;padding-left: 17%">
    <div style="margin-top: 1%;margin-left: 1%; margin-bottom: 1%;
    width:90%;color: black;border: #bee5eb solid 1px; border-radius: 5px; border-top: none ">
        <div class="alert alert-info" role="alert"
             style="width: 100%;margin-bottom: 0;margin-top: 1%;border-radius: 5px 5px 0 0;color: black">
            <a href="index?sessionId=<s:property value="getSessionId()"/>" style="color: black">主页</a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;个人发帖管理
        </div>
        <table id="example" class="table table-striped table-bordered text-center"
               style="width:100%;font-size: 13px;color: black;   table-layout:fixed;">
            <thead>
            <tr style="height: 5% " class="text-center">
                <td width="50%">帖子名</td>
                <td width="10%">发布时间</td>
                <td width="40%">回复内容</td>
            </tr>
            </thead>
            <tbody>
            <s:iterator value="getReplies()" var="data">
                <tr>
                    <td>
                        <a href="/post?sessionId=<s:property value="getSessionId()" />&&post=<s:property value="#data.getPid().getId()"/>"><s:property
                                value="#data.getPid().getName()"/></a></td>
                    <td><s:property
                            value="#data.getPublishTime()"/></td>
                    <td style="word-break:keep-all;/* 不换行 */
    white-space:nowrap;/* 不换行 */
    overflow:hidden;/* 内容超出宽度时隐藏超出部分的内容 */
    text-overflow:ellipsis;/* 当对象内文本溢出时显示省略标记(...) ；需与overflow:hidden;一起使用。*/"><s:property
                            value="#data.getContent()"/></td>
                </tr>
            </s:iterator>
            </tbody>
        </table>
    </div>
</div>
<script src="https://cdn.staticfile.org/datatables/1.10.19/js/jquery.dataTables.min.js"></script>
<script src="https://cdn.staticfile.org/datatables/1.10.19/js/dataTables.bootstrap4.min.js"></script>
<script>
    $(document).ready(function () {
        $('#example').DataTable({
            retrieve: true,
            paging: true,
            autoWidth: false,
            lengthChange: false,
            order: false,
            bSort: false,
            searching: false,
            async: false,
            info: false,
            stateSave: true,
            serverSide: false,
            processing: true,
            pageLength: 20,
            language:
                {
                    "sProcessing":
                        "处理中...",
                    "sEmptyTable":
                        "没有帖子",
                    "sLoadingRecords":
                        "载入中...",
                    "oPaginate":
                        {
                            "sFirst":
                                "首页",
                            "sPrevious":
                                "上页",
                            "sNext":
                                "下页",
                            "sLast":
                                "末页"
                        }
                    ,
                }
            ,

        })
        ;
    })
    ;
</script>
</body>
</html>

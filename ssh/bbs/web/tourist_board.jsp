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
        <img src="images/头像/none.png" style="margin-top:16%;border-radius: 100px; width: 100px;">
    </div>
    <div class="text-center">
        <p>当前为游客状态<br/>无法进行发/回帖操作</p>
        <p>请<a href="/login">登录</a></p>
    </div>
</div>


<!-- hot post bar-->
<div class="float-right" style="margin-top: 4%;
width:100%;color: black; overflow:hidden;padding-left: 17%">
    <div style="margin-top: 1%;margin-left: 1%; margin-bottom: 1%;
width:90%;color: black;border:  #bee5eb solid 1px; border-radius: 5px; border-top: none ">
        <div class="alert alert-info" role="alert"
             style="width: 100%;margin-bottom: 0;margin-top: 1%;border-radius: 5px 5px 0 0;color: black">
            <a href="index" style="color: black">主页</a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;<s:property
                value="boar.getName()"/>
        </div>
        <table id="example" class="table table-striped table-bordered" style="width:100%;font-size: 13px;color: black">
            <thead>
            <tr style="height: 5%;">
                <td width="60%">帖子名</td>
                <td width="12%">发布者</td>
                <td width="7%">发布时间</td>
                <td width="7%">阅读量</td>
                <td width="7%">回复量</td>
            </tr>
            </thead>
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
            searching: false,
            async: false,
            info: false,
            stateSave: true,
            serverSide: true,
            processing: true,
            pageLength: 40,
            ajax: {
                url: "getPosts",
                type: "POST",
                cache: false,
                data: {
                    "board": <s:property value="getBoard()"/>
                }
            },
            columns: [
                {"data": "postName"},
                {"data": "author"},
                {"data": "publishTime"},
                {"data": "readCount"},
                {"data": "replyCount"}
            ],
            language:
                {
                    "sProcessing":
                        "处理中...",
                    "sEmptyTable":
                        "未搜索到数据",
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

        });
    })
    ;
</script>
</body>
</html>

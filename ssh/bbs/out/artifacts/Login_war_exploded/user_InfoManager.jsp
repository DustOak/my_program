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
    <script type="text/javascript"
            src="https://cdn.staticfile.org/jquery-validate/1.19.0/jquery.validate.min.js"></script>
    <script type="text/javascript" src="https://cdn.staticfile.org/js-xss/0.3.3/xss.min.js"></script>
    <script>
        $(function () {
            var validate = $("#info").validate({
                debug: false, //调试模式取消submit的默认提交功能
                focusInvalid: false, //当为false时，验证无效时，没有焦点响应
                onkeyup: false,
                submitHandler: function (form) {   //表单提交句柄,为一回调函数，带一个参数：form
                    nickName = ($("#nickName").val()).replace(/\s+/g, "");
                    qq = ($("#qq").val()).replace(/\s+/g, "");
                    email = ($("#email").val()).replace(/\s+/g, "");
                    if (nickName !== "" && qq !== "" && email !== "") {
                        $("#nickName").val(filterXSS(nickName));
                        $("#qq").val(filterXSS(qq));
                    }
                    form.submit();   //提交表单
                },
                errorElement: "span",
                rules:
                    {
                        nickName: {
                            required: true
                        }
                        ,
                        email: {
                            required: true,
                            email:
                                true
                        }
                        ,
                        qq: {
                            required: true,
                            rangelength:
                                [5, 12]
                        }
                        ,

                    }
                ,
                messages: {
                    nickName: {
                        required: "(不许为空)"
                    }
                    ,
                    email: {
                        required: "(不许为空)",
                        email:
                            "(E-Mail格式不正确)"
                    }
                    ,
                    qq: {
                        required: "(不许为空)",
                        rangelength:
                            "(长度范围为5-12位)"
                    }
                }

            })
        });
    </script>
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
                <img src="images/head/<s:property value=" student.getPhotoPath()"/> "
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
            <p><a class="btn btn-info" href="/index?sessionId=<s:property value=" sessionId"/>">返回主页</a></p>
            <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value=" sessionId"/>">查看我的帖子</a></p>
            <p><a class="btn btn-info" href="/myReplies?sessionId=<s:property value=" sessionId"/>">查看我的回复</a></p>
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
            <p><a class="btn btn-info" href="/index?sessionId=<s:property value=" sessionId"/>">返回主页</a></p>
            <p><a class="btn btn-info" href="/myPost?sessionId=<s:property value=" sessionId"/>">查看我的帖子</a></p>
            <p><a class="btn btn-info" href="/myReplies?sessionId=<s:property value=" sessionId"/>">查看我的回复</a></p>
            <p><a class="btn btn-info" href="/boardManager?sessionId=<s:property value=" sessionId"/>">板块操作</a></p>
            <p><a class="btn btn-warning" href="/logout?sessionId=<s:property value=" sessionId"/>">注销</a></p>
        </div>
    </s:else>
</div>

<!-- hot post bar-->
<div class="float-right" style="margin-top: 4%;
    width:100%;color: black; overflow:hidden;padding-left: 17%">
    <div style="margin-top: 1%;margin-left: 1%; margin-bottom: 1%;
    width:90%;color: black;border: #bee5eb solid 1px; border-radius: 5px; border-top: none;height: 140%;">
        <div class="alert alert-info" role="alert"
             style="width: 100%;margin-bottom: 0;margin-top: 1%;border-radius: 5px 5px 0 0;color: black">
            <a href="index?sessionId=<s:property value=" getSessionId()"/>" style="color: black">主页</a>&nbsp;&nbsp;&nbsp;>>&nbsp;&nbsp;&nbsp;个人信息管理
        </div>
        <h3 style="margin-left: 2%;margin-top: 1%">个人信息管理</h3>
        <hr style="width: 95%">
        <form style="margin: 1% 2% 2% 15%" action="modifyInfo?sessionId=<s:property value=" getSessionId()"/>" id="info"
              method="post"
              enctype="multipart/form-data">
            <figure class="figure">
                <s:if test="getAdmin()!=null">
                    <s:if test="getAdmin().getPhotoPath()!=null">
                        <img src="images/head/<s:property value="getAdmin().getPhotoPath()"/>"
                             class="figure-img  rounded" id="headIcon"

                             width="200"
                             height="200">
                    </s:if>
                    <s:else>
                        <img src="images/head/noneHead.jpg" class="figure-img  rounded" id="headIcon"
                             width="200"
                             height="200">
                    </s:else>
                </s:if>
                <s:else>
                    <s:if test="getStudent().getPhotoPath()!=null">
                        <img src="images/head/<s:property value="getStudent().getPhotoPath()"/>" id="headIcon"
                             class="figure-img  rounded"
                             width="200"
                             height="200">
                    </s:if>
                    <s:else>
                        <img src="images/head/noneHead.jpg" class="figure-img  rounded" id="headIcon"
                             width="200"
                             height="200">
                    </s:else>
                </s:else>

                <figcaption class="figure-caption">
                    <input type="file" class="form-control-file" accept="image/*" id="icon" name="icon">
                    <label>更换头像</label>
                </figcaption>
            </figure>
            <div class="form-group" style="width: 50%;float: right;margin-right: 25%">
                <fieldset>
                    <p>
                        <label for="nickName">昵称:</label>
                        <input type="text" class="form-control" id="nickName" name="nickName"
                               style="margin-bottom: 2%"
                               value="<s:property  value="getStudent().getNickName()"/><s:property  value="getAdmin().getNickname()"/>">
                    </p>
                    <p>
                        <label for="realName">真实姓名</label>
                        <input type="text" class="form-control" id="realName" name="realName" disabled
                               p style="margin-bottom: 2%"
                               value="<s:property  value="getStudent().getRealName()"/><s:property  value="getAdmin().getName()"/>">
                    </p>
                    <s:if test="getStudent()!=null">
                        <p>
                            <label for="qq">QQ</label>
                            <input type="text" class="form-control" id="qq" name="qq"
                                   style="margin-bottom: 2%" value="<s:property  value="getStudent().getQq()"/>">
                        </p>
                        <p>
                            <label for="email">Email</label>
                            <input type="email" class="form-control" id="email" name="email"
                                   style="margin-bottom: 2%" value="<s:property  value="getStudent().getEmail()"/>">
                        </p>
                        <p>
                            <label for="major">专业</label>
                            <input type="text" class="form-control" id="major" disabled
                                   style="margin-bottom: 2%" value="<s:property  value="getStudent().getMajor()"/>">
                        </p>
                        <p>
                            <label for="className">班级</label>
                            <input type="text" class="form-control" id="className" disabled
                                   style="margin-bottom: 2%" value="<s:property  value="getStudent().getClassName()"/>">
                        </p>
                        <p>
                            <label for="stuNum">学号</label>
                            <input type="text" class="form-control" id="stuNum" disabled
                                   style="margin-bottom: 2%" value="<s:property  value="getStudent().getStuNum()"/>">
                        </p>
                    </s:if>
                </fieldset>
            </div>
            <button class="btn btn-success btn-block " style="width: 20%;" type="submit">保存</button>
            <button class="btn btn-secondary btn-block  " style="width: 20%;" type="button" onclick="backHome()">
                取消并返回首页
            </button>
        </form>
        <h3 style="margin-left: 2%;margin-top: 23%">密码管理</h3>
        <hr style="width: 95%;">
        <form style="margin: 3% 2% 2% 15%" id="pwd" action="modifyPassword" method="post"
              onsubmit="return checkEmpty()">
            <div class="form-group" style="width: 50%;float: right;margin-right: 25%">
                <div class="alert alert-danger" role="alert" id="wrong" style="display: none">
                    旧密码错误
                </div>
                <div class="alert alert-success" role="alert" id="success" style="display: none">
                    旧密码验证成功
                </div>
                <div class="alert alert-danger" role="alert" id="empty" style="display: none">
                    密码禁止为空
                </div>
                <div class="alert alert-danger" role="alert" id="repassword" style="display: none">
                    新密码两次必须相同
                </div>
                <p><label for="nickName">旧密码</label>
                    <input type="password" class="form-control" id="oldPassword" name="oldPassword"
                           style="margin-bottom: 2%"></p>

                <p><label for="nickName">新密码(禁止使用空格符号)</label>
                    <input type="password" class="form-control" id="newPassword" name="newPassword"
                           style="margin-bottom: 2%"></p>
                <p><label for="nickName">再次输入新密码(禁止使用空格符号)</label>
                    <input type="password" class="form-control" id="newPasswordRe" name="newPasswordRe"
                           style="margin-bottom: 2%"></p>
                <input type="hidden" value="<s:property  value="getSessionId()"/>" name="sessionId">
            </div>
            <button class="btn btn-success btn-block " style="width: 20%;" type="submit" id="submit" disabled>保存
            </button>
            <button class="btn btn-secondary btn-block  " style="width: 20%;" type="button" onclick="backHome()">
                取消并返回首页
            </button>
        </form>
    </div>
</div>
<script type="text/javascript">
    $("#icon").change(function () {
        var oFReader = new FileReader();
        var file = document.getElementById('icon').files[0];
        oFReader.readAsDataURL(file);
        oFReader.onloadend = function (oFRevent) {
            var src = oFRevent.target.result;
            $('#headIcon').attr('src', src);
        }
    });

    function checkEmpty() {
        oldPwd = $("#oldPassword").val().replace(/\s+/g, "");
        newPassword = $("#newPassword").val().replace(/\s+/g, "");
        newPasswordRe = $("#newPasswordRe").val().replace(/\s+/g, "");
        if (oldPwd === "" || newPassword === "" || newPasswordRe === "") {
            $("#empty").attr("style", "display:block;");
            return false;
        } else {
            if (newPasswordRe !== newPassword) {
                $("#repassword").attr("style", "display:block;");
                return false;
            } else {
                return true;
            }
        }
    }

    $("#oldPassword").focus(function () {
        $("#success").attr("style", "display:none;");
        $("#wrong").attr("style", "display:none;");
        $("#empty").attr("style", "display:none;");
    });
    $("#newPassword").focus(function () {
        $("#success").attr("style", "display:none;");
        $("#wrong").attr("style", "display:none;");
        $("#empty").attr("style", "display:none;");
        $("#repassword").attr("style", "display:none;");
    });
    $("#newPasswordRe").focus(function () {
        $("#success").attr("style", "display:none;");
        $("#wrong").attr("style", "display:none;");
        $("#empty").attr("style", "display:none;");
        $("#repassword").attr("style", "display:none;");
    });
    $("#oldPassword").blur(function () {
        oldPwd = $("#oldPassword").val();
        $.ajax({
            data: {
                "oldPwd": oldPwd,
                "sessionId": '<s:property  value="getSessionId()"/>',
            },
            url: "/checkPassword",
            type: "post",
            success: function (returnData) {
                if (returnData.error === 0) {
                    $("#success").attr("style", "display:block;");
                    document.getElementById("submit").removeAttribute("disabled");
                } else {
                    $("#wrong").attr("style", "display:block;");
                    document.getElementById("submit").disabled = true;
                }
            }
        });
    })

    function backHome() {
        location.href = "/index?sessionId="
    }
</script>
</body>
</html>

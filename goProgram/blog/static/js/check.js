function getCode(e){
    if(checkEmail()){
        $("#sub").removeAttr("disabled");
        $("#error").hide();
        var email = $('#email').val();
        var token = $('#token').val();
        $.ajax({
            url:"/admin/sendEmail",
            data:{
                email:email,
                token:token
            },
            success:function(data){
                alert(data)
            },
            method:"POST"
        })
        settime(e); //倒计时
    }else{
        $('#email').focus();
    }

}
function checkEmail(){
    var email = $('#email').val().replace(/\s+/g,"");
    var pattern =  /^\w+((.\w+)|(-\w+))@[A-Za-z0-9]+((.|-)[A-Za-z0-9]+).[A-Za-z0-9]+$/;
    if(email == '') {
        $("#errorInfo").text("邮箱不能为空");
        $("#error").show()
        return false;
    }
    if(!pattern.test(email)){
        $("#errorInfo").text("邮箱格式不正确");
        $("#error").show();
        return false;
    }
    return true;
}
//倒计时
var countdown=60;
function settime(val) {
    if (countdown == 0) {
        val.removeAttribute("disabled");
        val.value="获取验证码";
        countdown = 60;
    } else {
        val.setAttribute("disabled", true);
        val.value="重新发送(" + countdown + ")";
        countdown--;
        setTimeout(function() {
            settime(val)
        },1000)
    }

}

function checkSubmit() {
    var email = $('#email').val().replace(/\s+/g,"");
    var password=$('#password').val().replace(/\s+/g,"");
    var code=$('#validate_code').val().replace(/\s+/g,"");
    var reg=/^\d{6}$/;
    if(email == ''||password==''||code=="") {
        $("#errorInfo").text("账号密码或验证码不能为空");
        $("#error").show()
        return false;
    }
    if (reg.test(code)){
        return checkEmail();
    }else {
        $("#errorInfo").text("验证码格式不正确");
        $("#error").show()
        return false;
    }

}


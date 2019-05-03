function  checkEmpty() {
    accout = $("#account").val();
    password = $("#password").val();
    accout = accout.replace(/\s+/g, "");
    if (accout == "" || password == ""){
        alert("The Account Or Password cant empty");
        return false;
    }
    return true;
}
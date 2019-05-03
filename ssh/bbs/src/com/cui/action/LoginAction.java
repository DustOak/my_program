package com.cui.action;


import com.cui.dao.DaoOperating;
import com.cui.util.Encryption;
import com.cui.po.Admin;
import com.cui.util.Session;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionContext;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;

import javax.servlet.http.Cookie;

import java.util.List;


public class LoginAction extends ActionSupport {
private String account;
private String password;
private String sessionID;

public void validate() {
	if (account.equals("") || password.equals("")) {
		addFieldError("error", "The Account Or Password Cant Empty!");
	}
}

public String getPassword() {
	return password;
}

public void setPassword(String password) {
	this.password = password;
}

public String getAccount() {
	return account;
}

public void setAccount(String account) {
	this.account = account;
}

public String execute() {
	Admin admin = new Admin();
	admin.setAccount(this.account);
	admin.setPassword(Encryption.getMd5(this.password));
	List<Object> list = DaoOperating.Gets(admin);
	if (! list.isEmpty()) {
		admin = (Admin) list.get(0);
		String token = Encryption.getRandomToken();
		SessionManager.Put(token, new Session(admin, ServletActionContext.getRequest().getRemoteAddr()));
		Cookie cookie = new Cookie("sessionId", token);
		cookie.setMaxAge(3600);
		cookie.setPath("/");
		ServletActionContext.getResponse().addCookie(cookie);
		this.setSessionID(token);
		return SUCCESS;
	}
	addFieldError("error", "The Account Or Password Wrong!");
	return INPUT;
}


public String getSessionID() {
	return sessionID;
}

public void setSessionID(String sessionID) {
	this.sessionID = sessionID;
}
}

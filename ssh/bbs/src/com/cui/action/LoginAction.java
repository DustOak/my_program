package com.cui.action;


import com.cui.dao.DaoOperating;
import com.cui.po.Student;
import com.cui.service.AdminLoadService;
import com.cui.service.StudentLoadService;
import com.cui.util.Encryption;
import com.cui.po.Admin;
import com.cui.util.Session;
import com.cui.util.SessionManager;

import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.servlet.http.Cookie;

import java.util.List;

@Component
public class LoginAction extends ActionSupport {
private String account;
private String password;
private String sessionID;
private Integer adminOrStudent;

public AdminLoadService getAdminLoadService() {
	return adminLoadService;
}

public void setAdminLoadService(AdminLoadService adminLoadService) {
	this.adminLoadService = adminLoadService;
}

public StudentLoadService getStudentLoadService() {
	return studentLoadService;
}

public void setStudentLoadService(StudentLoadService studentLoadService) {
	this.studentLoadService = studentLoadService;
}

@Autowired
private AdminLoadService adminLoadService;
@Autowired
private StudentLoadService studentLoadService;

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
	Object oj = null;
	switch (adminOrStudent) {
		case 0:
			oj = studentLoadService.CheckUsernameAndPassword(this.getAccount(), Encryption.getMd5(this.getPassword()));
			break;
		case 1:
			oj = adminLoadService.CheckUsernameAndPassword(this.getAccount(), Encryption.getMd5(this.getPassword()));
			break;
		default:
			addFieldError("error", "Illegal identity!");
			return INPUT;
	}
	if (oj != null) {
		String token = Encryption.getRandomToken();
		SessionManager.Put(token, new Session(oj, ServletActionContext.getRequest().getRemoteAddr()));
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

public Integer getAdminOrStudent() {
	return adminOrStudent;
}

public void setAdminOrStudent(Integer adminOrStudent) {
	this.adminOrStudent = adminOrStudent;
}
}

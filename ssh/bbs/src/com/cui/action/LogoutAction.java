package com.cui.action;


import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;

import javax.servlet.http.Cookie;

public class LogoutAction extends ActionSupport {
private String sessionId;

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public String execute() {
	SessionManager.Remove(sessionId);
	Cookie cookie = new Cookie("sessionId", "-1");
	cookie.setMaxAge(0);
	cookie.setPath("/");
	ServletActionContext.getResponse().addCookie(cookie);
	return SUCCESS;
}
}

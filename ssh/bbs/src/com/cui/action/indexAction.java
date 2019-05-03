package com.cui.action;

import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;


public class indexAction extends ActionSupport {
private String sessionId = null;

public String execute() {
	this.setSessionId(ServletActionContext.getRequest().getAttribute("sessionId").toString());
	if (sessionId == null) {
		return INPUT;
	}
	ServletActionContext.getRequest().removeAttribute("sessionId");
	return SUCCESS;
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}
}

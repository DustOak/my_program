package com.cui.action;

import com.cui.po.Admin;
import com.cui.po.Student;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;

public class MyInfoAction extends ActionSupport {
private String sessionId;
private Student student;
private Admin admin;

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public Student getStudent() {
	return student;
}

public void setStudent(Student student) {
	this.student = student;
}

public Admin getAdmin() {
	return admin;
}

public void setAdmin(Admin admin) {
	this.admin = admin;
}

public String execute() {
	if (sessionId == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj instanceof Admin) {
			admin = (Admin) oj;
		} else {
			student = (Student) oj;
		}
		return SUCCESS;
	}
}
}

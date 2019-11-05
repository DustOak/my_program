package com.cui.action;


import com.cui.po.Admin;
import com.cui.po.Student;
import com.cui.service.AdminLoadService;
import com.cui.service.StudentLoadService;
import com.cui.util.Encryption;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class ModifyPasswordAction extends ActionSupport {
private String sessionId;
private String oldPassword;
private String newPassword;
private String newPasswordRe;
@Autowired
private StudentLoadService studentLoadService;
@Autowired
private AdminLoadService adminLoadService;

public String execute() {
	if (sessionId == null || oldPassword == null ||
				! newPassword.equals(newPasswordRe) || newPassword == null || newPasswordRe == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj instanceof Admin) {
			if (((Admin) oj).getPassword().equals(Encryption.getMd5(oldPassword))) {
				((Admin) oj).setPassword(Encryption.getMd5(newPassword));
				adminLoadService.Update((Admin) oj);
				SessionManager.Remove(sessionId);
				return SUCCESS;
			} else {
				return ERROR;
			}
		} else {
			if (((Student) oj).getPassword().equals(Encryption.getMd5(oldPassword))) {
				((Student) oj).setPassword(Encryption.getMd5(newPassword));
				studentLoadService.Update((Student) oj);
				SessionManager.Remove(sessionId);
				return SUCCESS;
			} else {
				return ERROR;
			}
		}
	}
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public String getOldPassword() {
	return oldPassword;
}

public void setOldPassword(String oldPassword) {
	this.oldPassword = oldPassword;
}

public String getNewPassword() {
	return newPassword;
}

public void setNewPassword(String newPassword) {
	this.newPassword = newPassword;
}

public String getNewPasswordRe() {
	return newPasswordRe;
}

public void setNewPasswordRe(String newPasswordRe) {
	this.newPasswordRe = newPasswordRe;
}

public StudentLoadService getStudentLoadService() {
	return studentLoadService;
}

public void setStudentLoadService(StudentLoadService studentLoadService) {
	this.studentLoadService = studentLoadService;
}

public AdminLoadService getAdminLoadService() {
	return adminLoadService;
}

public void setAdminLoadService(AdminLoadService adminLoadService) {
	this.adminLoadService = adminLoadService;
}
}

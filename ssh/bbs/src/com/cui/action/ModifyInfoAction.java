package com.cui.action;


import com.cui.po.Admin;
import com.cui.po.Student;
import com.cui.service.AdminLoadService;
import com.cui.service.StudentLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.commons.io.FileUtils;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.io.File;

@Component
public class ModifyInfoAction extends ActionSupport {
private String sessionId;
private String nickName;
private String qq;
private String email;

public String getDestPath() {
	return destPath;
}

public void setDestPath(String destPath) {
	this.destPath = destPath;
}

public File getIcon() {
	return icon;
}

public void setIcon(File icon) {
	this.icon = icon;
}

public String getIconContentType() {
	return iconContentType;
}

public void setIconContentType(String iconContentType) {
	this.iconContentType = iconContentType;
}

public String getIconFileName() {
	return iconFileName;
}

public void setIconFileName(String iconFileName) {
	this.iconFileName = iconFileName;
}

private File icon;
private String iconContentType;
private String iconFileName;
private String destPath;
@Autowired
private StudentLoadService studentLoadService;
@Autowired
private AdminLoadService adminLoadService;

public String execute() {
	if (sessionId == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj instanceof Admin) {
			((Admin) oj).setNickname(nickName);
			try {
				if (icon != null) {
					((Admin) oj).setPhotoPath(iconFileName);
					File icons = new File(destPath, iconFileName);
					FileUtils.copyFile(icon, icons);
				}
			} catch (Exception ex) {
				ex.printStackTrace();
				return ERROR;
			}
			adminLoadService.Update((Admin) oj);
			SessionManager.Get(sessionId).setObject(oj);
		} else {
			try {
				if (icon != null) {
					((Student) oj).setPhotoPath(iconFileName);
					File icons = new File(destPath, iconFileName);
					FileUtils.copyFile(icon, icons);
				}
			} catch (Exception ex) {
				ex.printStackTrace();
				return ERROR;
			}
			((Student) oj).setEmail(email);
			((Student) oj).setQq(qq);
			((Student) oj).setNickName(nickName);
			studentLoadService.Update((Student) oj);
			SessionManager.Get(sessionId).setObject(oj);
		}
		return SUCCESS;
	}
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public String getNickName() {
	return nickName;
}

public void setNickName(String nickName) {
	this.nickName = nickName;
}

public String getQq() {
	return qq;
}

public void setQq(String qq) {
	this.qq = qq;
}

public String getEmail() {
	return email;
}

public void setEmail(String email) {
	this.email = email;
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

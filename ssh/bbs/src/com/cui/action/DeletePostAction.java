package com.cui.action;

import com.cui.po.Admin;
import com.cui.po.Student;
import com.cui.service.PostLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class DeletePostAction extends ActionSupport {
private String sessionId;
private Integer post;

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public Integer getPost() {
	return post;
}

public void setPost(Integer post) {
	this.post = post;
}

public PostLoadService getPostLoadService() {
	return postLoadService;
}

public void setPostLoadService(PostLoadService postLoadService) {
	this.postLoadService = postLoadService;
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

@Autowired
private PostLoadService postLoadService;
private Student student;
private Admin admin;

public String execute() {
	if (sessionId == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj instanceof Admin) {
			admin = (Admin) oj;
			postLoadService.deletePost(post);
		} else {
			student = (Student) oj;
			postLoadService.deletePost(post);
		}
		return SUCCESS;
	}
}
}

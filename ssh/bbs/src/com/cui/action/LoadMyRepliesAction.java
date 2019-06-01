package com.cui.action;

import com.cui.po.Admin;
import com.cui.po.Reply;
import com.cui.po.Student;
import com.cui.service.ReplyLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class LoadMyRepliesAction extends ActionSupport {
private String sessionId;
@Autowired
private ReplyLoadService replyLoadService;
private List<Reply> replies;

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
			replies = replyLoadService.getUserReplies(oj);
		} else {
			student = (Student) oj;
			replies = replyLoadService.getUserReplies(oj);
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

public ReplyLoadService getReplyLoadService() {
	return replyLoadService;
}

public void setReplyLoadService(ReplyLoadService replyLoadService) {
	this.replyLoadService = replyLoadService;
}

public List<Reply> getReplies() {
	return replies;
}

public void setReplies(List<Reply> replies) {
	this.replies = replies;
}
}

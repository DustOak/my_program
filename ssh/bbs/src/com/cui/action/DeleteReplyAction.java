package com.cui.action;


import com.cui.po.Reply;
import com.cui.service.ReplyLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class DeleteReplyAction extends ActionSupport {
private String sessionId;
private Integer reply;
private Integer post;
@Autowired
private ReplyLoadService replyLoadService;
public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public Integer getReply() {
	return reply;
}

public void setReply(Integer reply) {
	this.reply = reply;
}

public Integer getPost() {
	return post;
}

public void setPost(Integer post) {
	this.post = post;
}

public ReplyLoadService getReplyLoadService() {
	return replyLoadService;
}

public void setReplyLoadService(ReplyLoadService replyLoadService) {
	this.replyLoadService = replyLoadService;
}

public String execute() {
	if (sessionId == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
	
		replyLoadService.deleteReply(new Reply(reply));
		return SUCCESS;
	}
}


}

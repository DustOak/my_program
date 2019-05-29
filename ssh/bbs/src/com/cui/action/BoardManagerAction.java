package com.cui.action;

import com.cui.po.Admin;
import com.cui.po.Board;
import com.cui.service.BoardLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;


@Component
public class BoardManagerAction extends ActionSupport {
private String sessionId;
@Autowired
private BoardLoadService boardLoadService;

public List<Board> getBoards() {
	return boards;
}

private Admin admin;

public void setBoards(List<Board> boards) {
	this.boards = boards;
}

private List<Board> boards;

public String execute() {
	if (sessionId == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
		boards = boardLoadService.loadRootBoards();
		admin = (Admin) SessionManager.Get(sessionId).getObject();
		return SUCCESS;
	}
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public BoardLoadService getBoardLoadService() {
	return boardLoadService;
}

public void setBoardLoadService(BoardLoadService boardLoadService) {
	this.boardLoadService = boardLoadService;
}

public Admin getAdmin() {
	return admin;
}

public void setAdmin(Admin admin) {
	this.admin = admin;
}
}

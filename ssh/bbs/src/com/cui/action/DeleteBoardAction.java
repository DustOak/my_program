package com.cui.action;

import com.cui.service.BoardLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class DeleteBoardAction extends ActionSupport {
private String sessionId;
private Integer board;
@Autowired
private BoardLoadService boardLoadService;

public String execute() {
	if (board == null || sessionId == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
		boardLoadService.Delete(board);
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

public Integer getBoard() {
	return board;
}

public void setBoard(Integer board) {
	this.board = board;
}
}

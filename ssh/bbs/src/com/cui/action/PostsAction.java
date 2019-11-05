package com.cui.action;

import com.cui.po.Admin;
import com.cui.po.Board;
import com.cui.po.Student;
import com.cui.service.BoardLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


@Component
public class PostsAction extends ActionSupport {
private int board;
private Board boar;
@Autowired
private BoardLoadService boardLoadService;
private String sessionId;
private Admin admin;
private Student student;

public Admin getAdmin() {
	return admin;
}

public void setAdmin(Admin admin) {
	this.admin = admin;
}

public Student getStudent() {
	return student;
}

public void setStudent(Student student) {
	this.student = student;
}

public String execute() {
	boar = boardLoadService.loadBoard(board);
	if (sessionId != null && ! SessionManager.IsExist(sessionId) && ! SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj != null) {
			if (oj instanceof Admin) {
				admin = (Admin) oj;
			} else {
				student = (Student) oj;
			}
		}
		if (oj != null) {
			return "USER_LOGIN";
		} else {
			return "TOURIST";
		}
		
	} else {
		return "TOURIST";
	}
	
}

public int getBoard() {
	return board;
}

public void setBoard(int board) {
	this.board = board;
}


public BoardLoadService getBoardLoadService() {
	return boardLoadService;
}

public void setBoardLoadService(BoardLoadService boardLoadService) {
	this.boardLoadService = boardLoadService;
}

public Board getBoar() {
	return boar;
}

public void setBoar(Board boar) {
	this.boar = boar;
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}
}

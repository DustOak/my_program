package com.cui.action;

import com.cui.po.Admin;
import com.cui.po.Board;
import com.cui.po.Post;
import com.cui.po.Student;
import com.cui.service.BoardLoadService;

import com.cui.service.PostLoadService;
import com.opensymphony.xwork2.ActionSupport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Component
public class MainAction extends ActionSupport {
private List<Board> boards;
private Map<Integer, List<Board>> childBoard = new HashMap<>();
private List<Post> hotPosts;
@Autowired
BoardLoadService boardLoad;
private String sessionId;
private Admin admin;

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

private Student student;

public PostLoadService getPostLoad() {
	return postLoad;
}

public void setPostLoad(PostLoadService postLoad) {
	this.postLoad = postLoad;
}

@Autowired
PostLoadService postLoad;

public Map<Integer, List<Board>> getChildBoard() {
	return childBoard;
}

public void setChildBoard(Map<Integer, List<Board>> childBoard) {
	this.childBoard = childBoard;
}


public String execute() {
	boards = boardLoad.loadRootBoards();
	for (int j = 0; j < boards.size(); j++) {
		childBoard.put(boards.get(j).getId(), boardLoad.loadChildBoards(boards.get(j).getId()));
	}
	hotPosts = postLoad.rankPosts(6);
	return SUCCESS;
}

public List<Board> getBoards() {
	return boards;
}

public void setBoards(List<Board> boards) {
	this.boards = boards;
}

public BoardLoadService getBoardLoad() {
	return boardLoad;
}

public void setBoardLoad(BoardLoadService boardLoad) {
	this.boardLoad = boardLoad;
}


public List<Post> getHotPosts() {
	return hotPosts;
}

public void setHotPosts(List<Post> hotPosts) {
	this.hotPosts = hotPosts;
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}
}

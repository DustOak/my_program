package com.cui.action;

import com.cui.po.Board;
import com.cui.service.BoardLoadService;

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
@Autowired
BoardLoadService boardLoad;

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


}

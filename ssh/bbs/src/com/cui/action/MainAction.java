package com.cui.action;

import com.cui.po.Board;
import com.cui.service.BoardLoad;
import com.opensymphony.xwork2.ActionSupport;

import java.util.List;

public class MainAction extends ActionSupport {
List<Board> boards;
BoardLoad boardLoad;

public List<Board> getBoards() {
	return boards;
}

public void setBoards(List<Board> boards) {
	this.boards = boards;
}

public BoardLoad getBoardLoad() {
	return boardLoad;
}

public void setBoardLoad(BoardLoad boardLoad) {
	this.boardLoad = boardLoad;
}


}

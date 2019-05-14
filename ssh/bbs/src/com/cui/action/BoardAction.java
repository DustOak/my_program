package com.cui.action;

import com.cui.po.Board;
import com.cui.service.BoardLoadService;
import com.opensymphony.xwork2.ActionSupport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


@Component
public class BoardAction extends ActionSupport {
private int board;
private Board boar;
@Autowired
private BoardLoadService boardLoadService;

public String execute() {
	boar = boardLoadService.loadBoard(board);
	System.out.println();
	System.out.println();
	return SUCCESS;
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
}

package com.cui.action;

import com.opensymphony.xwork2.ActionSupport;

public class BoardAction extends ActionSupport {
private int board;

public String execute() {
	return SUCCESS;
}

public int getBoard() {
	return board;
}

public void setBoard(int board) {
	this.board = board;
}
}

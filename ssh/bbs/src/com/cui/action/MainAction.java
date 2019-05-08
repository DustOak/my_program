package com.cui.action;

import com.cui.po.Board;
import com.cui.service.BoardLoadService;
import com.cui.util.ApplicationContextOperator;
import com.opensymphony.xwork2.ActionSupport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


import java.util.List;

@Component
public class MainAction extends ActionSupport {
List<Board> boards;

@Autowired
BoardLoadService boardLoad;

public String execute() {
	for (String s : ApplicationContextOperator.getApplicationContext().getBeanDefinitionNames()) {
		System.out.println(s);
	}
	boards = boardLoad.loadAllBoards();
//	for (int i = 0; i < boards.size(); i++) {
//		System.out.println(boards.get(i).getAid());
//	}
	return INPUT;
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

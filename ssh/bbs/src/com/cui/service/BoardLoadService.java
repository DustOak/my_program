package com.cui.service;


import com.cui.dao.DaoOperating;
import com.cui.po.Board;


import java.util.List;

public class BoardLoadService implements BoardLoad {
@Override
public Board loadBoard(int id) {
	return (Board) DaoOperating.Get(new Board(), id);
}

@Override
public List<Board> loadChildBoards(int parentId) {
	String str = "from board where parentId=:0 order by id asc";
	return DaoOperating.Finds(str, parentId);
}

@Override
public List<Board> loadAllBoards() {
	return DaoOperating.Gets(new Board());
}

@Override
public List<Board> loadRootBoards() {
	String str = "from board where parentId=null order by id asc ";
	return DaoOperating.Finds(str);
}

@Override
public boolean SaveOrUpdate(Board board) {
	try {
		if (board.getId() == null) {
			return DaoOperating.Save(board);
		} else {
			DaoOperating.Update(board);
			return true;
		}
	} catch (Exception ex) {
		ex.printStackTrace();
		return false;
	}
}
}

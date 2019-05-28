package com.cui.service;


import com.cui.dao.DaoOperating;
import com.cui.po.Board;


import java.util.List;

public class BoardLoadService implements BoardLoad {
@Override
public Board loadBoard(Integer id) {
	return (Board) DaoOperating.Get(new Board(), id);
}

@Override
public List<Board> loadChildBoards(Integer parentId) {
	String str = "from Board  where parentId=" + parentId + " order by id asc";
	return DaoOperating.Finds(str);
}

@Override
public List<Board> loadAllBoards() {
	return DaoOperating.Gets(new Board());
}

@Override
public List<Board> loadRootBoards() {
	String str = "from Board where parentId is null order by id asc ";
	return DaoOperating.Finds(str);
}

@Override
public boolean SaveOrUpdate(Board board) {
	try {
		if (board.getId() == null) {
			DaoOperating.Save(board);
			return true;
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

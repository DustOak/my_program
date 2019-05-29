package com.cui.service;


import com.cui.dao.DaoOperating;
import com.cui.po.Board;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


import java.util.List;

@Component
public class BoardLoadService implements BoardLoad {
public DaoOperating getDaoOperating() {
	return daoOperating;
}

public void setDaoOperating(DaoOperating daoOperating) {
	this.daoOperating = daoOperating;
}

@Autowired
DaoOperating daoOperating;

@Override
public Board loadBoard(Integer id) {
	return (Board) daoOperating.Get(new Board(), id);
}

@Override
public List<Board> loadChildBoards(Integer parentId) {
	String str = "from Board  where parentId=" + parentId + " order by id asc";
	return daoOperating.Finds(str);
}

@Override
public List<Board> loadAllBoards() {
	return daoOperating.Gets(new Board());
}

@Override
public List<Board> loadRootBoards() {
	String str = "from Board where parentId is null order by id asc ";
	return daoOperating.Finds(str);
}

@Override
public boolean SaveOrUpdate(Board board) {
	try {
		if (board.getId() == null) {
			daoOperating.Save(board);
			return true;
		} else {
			daoOperating.Update(board);
			return true;
		}
	} catch (Exception ex) {
		ex.printStackTrace();
		return false;
	}
}

@Override
public boolean Delete(Integer id) {
	
	try {
		return daoOperating.Delete(new Board(id));
	} catch (Exception ex) {
		ex.printStackTrace();
		return false;
	}
}
}

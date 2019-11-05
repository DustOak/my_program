package com.cui.service;

import com.cui.po.Board;

import java.util.List;

public interface BoardLoad {
Board loadBoard(Integer id);

List<Board> loadChildBoards(Integer parentId);

List<Board> loadAllBoards();

List<Board> loadRootBoards();

boolean SaveOrUpdate(Board board);

boolean Delete(Integer id);
}

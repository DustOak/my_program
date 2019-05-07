package com.cui.service;

import com.cui.po.Board;

import java.util.List;

public interface BoardLoad {
Board loadBoard(int id);

List<Board> loadChildBoards(int parentId);

List<Board> loadAllBoards();

List<Board> loadRootBoards();

boolean SaveOrUpdate(Board board);

}

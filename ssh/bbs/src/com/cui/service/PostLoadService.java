package com.cui.service;

import com.cui.dao.DaoOperating;
import com.cui.po.Post;

import java.util.List;

public class PostLoadService implements PostLoad {
@Override
public boolean saveOrUpdate(Post post) {
	try {
		if (post.getId() == null) {
			return DaoOperating.Save(post);
		} else {
			DaoOperating.Update(post);
			return true;
		}
	} catch (Exception ex) {
		ex.printStackTrace();
		return false;
	}
}

@Override
public List<Post> allPost() {
	return DaoOperating.Gets(new Post());
}

@Override
public List<Post> pageAllPost(int bid, int pageNo, int pageSize) {
	String str = "from Post where bid=?0 order by id desc limit ?1*?2,?3 ";
	return DaoOperating.Finds(str, bid, pageNo, pageSize, pageSize);
}

@Override
public int getPostsCount() {
	return 0;
}

@Override
public Post loadPost(int id) {
	return null;
}

@Override
public List<Post> allPostsByUser(Object user) {
	return null;
}

@Override
public boolean deletePost(int id) {
	return false;
}

@Override
public List<Post> searchPosts(String searchKey) {
	return null;
}

@Override
public List<Post> rankPosts(int size) {
	return null;
}

@Override
public int countTotalPost() {
	return 0;
}

@Override
public int countTodayPost() {
	return 0;
}

@Override
public int countYesteradyPost() {
	return 0;
}

@Override
public int countDayLargestPost() {
	return 0;
}
}

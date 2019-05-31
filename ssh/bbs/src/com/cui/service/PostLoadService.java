package com.cui.service;


import com.cui.dao.DaoOperating;
import com.cui.po.Admin;

import com.cui.po.Post;
import com.cui.po.Student;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


import java.util.ArrayList;
import java.util.List;

@Component
public class PostLoadService implements PostLoad {
public DaoOperating getDaoOperating() {
	return daoOperating;
}

public void setDaoOperating(DaoOperating daoOperating) {
	this.daoOperating = daoOperating;
}

@Autowired
DaoOperating daoOperating;

@Override
public boolean saveOrUpdate(Post post) {
	try {
		if (post.getId() == null) {
			daoOperating.Save(post);
			return true;
		} else {
			daoOperating.Update(post);
			return true;
		}
	} catch (Exception ex) {
		ex.printStackTrace();
		return false;
	}
}

@Override
public List<Post> allPost() {
	return daoOperating.Gets(new Post());
}

@Override
public List<Post> pageAllPost(int bid, int start, int length) {
	String hql = "from Post as p where p.bid=" + bid + "order by publishTime desc";
	return daoOperating.LimitQuery(hql, start, length);
}

@Override
public int getPostsCount() {
	return daoOperating.Count(new Post());
}

public int getBoardPostsCount(int bid) {
	String hql = "from Post as p where p.bid = " + bid;
	return daoOperating.Finds(hql).size();
}

@Override
public List getUserPosts(Object oj) {
	
	return daoOperating.Gets(oj);
}

@Override
public Post loadPost(int id) {
	return (Post) daoOperating.Get(new Post(), id);
}

@Override
public List<Post> allPostsByUser(Object user) {
	
	if (user instanceof Admin) {
		String hql = "from Post as p where p.aid = " + ((Admin) user).getId();
		return daoOperating.Finds(hql);
	}
	if (user instanceof Student) {
		String hql = "from Post as p where p.sid = " + ((Student) user).getId();
		return daoOperating.Finds(hql);
	}
	return null;
}

@Override
public boolean deletePost(int id) {
	return daoOperating.Delete(new Post(id));
}

@Override
public List<Post> searchPosts(String searchKey) {
	String str = "from Post where name like %?0%";
	return daoOperating.Finds(str, searchKey);
}

@Override
public List<Post> rankPosts(int size) {
	String str = "from Post as p order by p.count desc";
	List<Post> posts = daoOperating.Finds(str);
	List<Post> hotPosts = new ArrayList<>();
	for (int i = 0; i < size && i < posts.size(); i++) {
		hotPosts.add(posts.get(i));
	}
	
	return hotPosts;
}

@Override
public int countTotalPost() {
	String str = "select count(*) from post";
	return daoOperating.Finds(str).size();
}

@Override
public int countTodayPost() {
	String str = "SELECT COUNT(*) FROM post WHERE DATEDIFF(publishtime,NOW()) = 0";
	return daoOperating.Finds(str).size();
}

@Override
public int countYesteradyPost() {
	String str = "SELECT COUNT(*) FROM post WHERE DATEDIFF(publishtime,NOW()) = -1";
	return daoOperating.Finds(str).size();
}

@Override
public int countDayLargestPost() {
	String str = "SELECT MAX(COUNT) FROM post WHERE DATEDIFF(publishtime,NOW())=0";
	return daoOperating.Finds(str).size();
}
}

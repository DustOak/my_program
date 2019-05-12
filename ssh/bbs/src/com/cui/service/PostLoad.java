package com.cui.service;

import com.cui.po.Post;

import java.util.List;

public interface PostLoad {
boolean saveOrUpdate(Post post);

List<Post> allPost();

List<Post> pageAllPost(int bid, int start, int length);

int getPostsCount();

Post loadPost(int id);

List<Post> allPostsByUser(Object user);

boolean deletePost(int id);

List<Post> searchPosts(String searchKey);

List<Post> rankPosts(int size);

int countTotalPost();

int countTodayPost();

int countYesteradyPost();

int countDayLargestPost();
}

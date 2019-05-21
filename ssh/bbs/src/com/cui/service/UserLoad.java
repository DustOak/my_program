package com.cui.service;

public interface UserLoad<T> {
	T GetUser(int id);
	T CheckUsernameAndPassword(String username,String password);
}

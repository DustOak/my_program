package com.cui.service;

import com.cui.dao.DaoOperating;


import com.cui.po.Student;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class StudentLoadService implements UserLoad<Student> {
public DaoOperating getDaoOperating() {
	return daoOperating;
}

public void setDaoOperating(DaoOperating daoOperating) {
	this.daoOperating = daoOperating;
}

@Autowired
DaoOperating daoOperating;

@Override
public Student GetUser(int id) {
	return (Student) daoOperating.Get(new Student(), id);
}

@Override
public Student CheckUsernameAndPassword(String username, String password) {
	String hql = "from Student  s where s.stuNum='" + username + "' and  s.password='" + password + "'";
	return daoOperating.Finds(hql).size() == 0 ? null : (Student) daoOperating.Finds(hql).get(0);
}

@Override
public void Update(Student student) {
	daoOperating.Update(student);
}
}

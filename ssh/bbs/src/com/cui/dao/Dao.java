package com.cui.dao;

import java.util.List;

public interface Dao {
Object Get(Object oj, int id);

List Gets(Object oj);

List Finds(String fields, Object... value);

void Save(Object oj);

void Update(Object oj);

boolean Delete(Object oj);

Object Query(String str);

List LimitQuery(String hql, int start, int length);

int Count(Object oj);
}

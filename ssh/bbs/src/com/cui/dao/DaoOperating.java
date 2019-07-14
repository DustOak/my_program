package com.cui.dao;

import org.hibernate.HibernateException;
import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.Transaction;
import org.hibernate.query.Query;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.orm.hibernate5.HibernateCallback;
import org.springframework.orm.hibernate5.HibernateTemplate;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;


import java.util.List;
import java.util.Locale;

@Component
@Service
public class DaoOperating implements Dao {

@Autowired
private HibernateTemplate hibernateTemplate;

@Transactional
public Object Get(Object oj, int id) {
	return getHibernateTemplate().get(oj.getClass(), id);
}

@Transactional
public List Gets(Object oj) {
	return getHibernateTemplate().findByExample(oj);
}

@Transactional
public List Finds(String fields, Object... value) {
	return getHibernateTemplate().execute(new HibernateCallback<List>() {
		@Override
		public List doInHibernate(Session session) throws HibernateException {
			Query query = session.createQuery(fields);
			for (int i = 0; i < value.length; i++) {
				query.setParameter(i, value[i]);
			}
			return query.list();
		}
	});
}

@Transactional
public void Save(Object oj) {
//	Session session = getHibernateTemplate().getSessionFactory().getCurrentSession();
//	Transaction transaction = session.getTransaction();
//	try {
//		session.save(oj);
//		transaction.commit();
//	} catch (Exception ex) {
//		transaction.rollback();
//		ex.printStackTrace();
//	} finally {
//		session.close();
//	}
	getHibernateTemplate().save(oj);
}

@Transactional
public void Update(Object oj) {
	getHibernateTemplate().update(oj);
}

@Transactional
public boolean Delete(Object oj) {
	try {
		getHibernateTemplate().delete(oj);
		return true;
	} catch (Exception ex) {
		ex.printStackTrace();
		return false;
	}
}

@Override
public Object Query(String str) {
	return null;
}

@Transactional
public List LimitQuery(String hql, int start, int length) {
	return getHibernateTemplate().execute(new HibernateCallback<List>() {
		@Override
		public List doInHibernate(Session session) throws HibernateException {
			Query query = session.createQuery(hql);
			query.setFirstResult(start);
			query.setMaxResults(length);
			return query.list();
		}
	});
}

@Transactional
public int Count(Object oj) {
	return getHibernateTemplate().findByExample(oj).size();
}

public HibernateTemplate getHibernateTemplate() {
	return this.hibernateTemplate;
}

public void setHibernateTemplate(HibernateTemplate hibernateTemplate) {
	this.hibernateTemplate = hibernateTemplate;
}

}

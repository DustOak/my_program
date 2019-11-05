package com.cui.util;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class ApplicationContextOperator {
private static ApplicationContext applicationContext;

public static void Init() {
	applicationContext = new ClassPathXmlApplicationContext("applicationContext.xml");
}

public static ApplicationContext getApplicationContext() {
	return applicationContext;
}

public static void setApplicationContext(ApplicationContext ac) {
	applicationContext = ac;
}
}

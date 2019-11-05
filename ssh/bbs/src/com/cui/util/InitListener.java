package com.cui.util;

import javax.servlet.ServletContextEvent;
import javax.servlet.ServletContextListener;

public class InitListener implements ServletContextListener {

public void contextDestroyed(ServletContextEvent sce) {
}

public void contextInitialized(ServletContextEvent sce) {
	ApplicationContextOperator.Init();
	SessionManager.Init(60 * 60 * 1000);
}
}
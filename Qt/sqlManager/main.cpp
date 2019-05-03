#include "mainwindow.h"
#include <QApplication>
#include <QDesktopWidget>
int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    MainWindow w;
    w.setWindowFlags(w.windowFlags()&~Qt::WindowMaximizeButtonHint&~Qt::WindowMinimizeButtonHint);
    QDesktopWidget *desk=new QDesktopWidget();
    w.move((desk->width()-w.width())/2,(desk->height()-w.height()-w.height())/2);
    w.show();


    return a.exec();
}

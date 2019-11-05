#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "index.h"
#include <QDesktopWidget>
#include <QMessageBox>
MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    this->setFixedSize(320,200);
    ui->setupUi(this);
    index=new Index();


}

MainWindow::~MainWindow()
{
    delete ui;
}


void MainWindow::on_connection_clicked()
{   QDesktopWidget *desk=new QDesktopWidget();
    index->SetConnectionValue(ui->serverName->text(),ui->port->text(),ui->userName->text(),ui->password->text());
    if(!index->ConnectionDatabase()){
       QMessageBox::critical(this,"Connection Failed",
                     "The Connection Of Database Is Failed! Please Restart The Program"
                             ,QMessageBox::Yes);
       return ;
    }
    this->close();
    index->move((desk->width()-index->width())/2,(desk->height()-index->height())/2);
    index->show();
    index->GetDatabaseList();
}


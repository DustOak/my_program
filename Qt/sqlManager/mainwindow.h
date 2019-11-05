#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include "index.h"
namespace Ui {
class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = nullptr);
    ~MainWindow();

private slots:
    void on_connection_clicked();


private:
    Ui::MainWindow *ui;
    Index *index;
};

#endif // MAINWINDOW_H

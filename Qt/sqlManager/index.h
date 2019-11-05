#ifndef INDEX_H
#define INDEX_H

#include <QMainWindow>
#include <QtSql/QSqlDatabase>
#include <QtSql/QSqlDriver>
#include <QtSql/QSqlError>
namespace Ui {
class Index;
}

class Index : public QMainWindow
{
    Q_OBJECT

public:
    explicit Index(QWidget *parent = nullptr);
    void SetConnectionValue(QString serverAddress,QString port,QString username,QString password);
    bool ConnectionDatabase();
    void GetDatabaseList();
    ~Index();

private slots:
   void on_db_clicked(const QModelIndex &index);


   void on_select_2_clicked();

   void on_TableListComboBox_5_activated(const QString &arg1);

   void on_fieldListComboBox_4_activated(const QString &arg1);

   void on_add_clicked();

   void on_submit_clicked();

private:
    Ui::Index *ui;
    QString serverAddress;
    QString port;
    QString username;
    QString password;
    QSqlDatabase conn;
    QString dbName;
    QString tableName;
    QList <QString>tableList;
    QStringList selectTable;
    QStringList fieldList;
    QStringList selectedField;
};

#endif // INDEX_H

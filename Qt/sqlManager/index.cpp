#include "index.h"
#include "ui_index.h"
#include <QtSql/QSqlDatabase>
#include <QtSql/QSqlQuery>
#include <QStandardItemModel>
#include <QtDebug>
#include <qmessagebox.h>
#include <QtSql/QSqlRecord>
Index::Index(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::Index)
{
    ui->setupUi(this);
}

void Index::SetConnectionValue(QString serverAddress, QString port, QString username, QString password)
{
    this->serverAddress=serverAddress;
    this->port=port;
    this->username=username;
    this->password=password;
}

bool Index::ConnectionDatabase()
{   this->conn=QSqlDatabase::addDatabase("QMYSQL");
    this->conn.setHostName(this->serverAddress);
    this->conn.setPort(this->port.toInt());
    this->conn.setUserName(this->username);
    this->conn.setPassword(this->password);
    return this->conn.open();
}

void Index::GetDatabaseList()
{
    ui->select_2->setEnabled(false);
    ui->insert_2->setEnabled(false);
    ui->update_2->setEnabled(false);
    ui->delete_3->setEnabled(false);
    ui->control->setHidden(true);
    QStandardItemModel* model = new QStandardItemModel(ui->db);
    model->setHorizontalHeaderLabels(QStringList()<<QStringLiteral("Databases"));
    QList<QStandardItem*> root;
    QSqlQuery getDbs(this->conn);
    getDbs.exec("SELECT schema_name FROM information_schema.schemata");
//    " WHERE schema_name!='information_schema' AND  schema_name!='mysql' AND "
//    "  schema_name!='performance_schema' "
    while(getDbs.next()){
        QString db=getDbs.value(0).toString();
        this->tableList.append("-------------"+db+"-------------");
        QSqlQuery getTables(this->conn);
        getTables.exec("show tables from "+db);
        QStandardItem *item = new QStandardItem(db);
        while (getTables.next()) {
             QString table=getTables.value(0).toString();
             this->tableList.append(db+"."+table);
             QStandardItem *item2=new QStandardItem(table);
             item->appendRow(+item2);
        }
        model->appendRow(item);
    }
    ui->db->setModel(model);

}


Index::~Index()
{
    this->conn.close();
    delete ui;
}

void Index::on_db_clicked(const QModelIndex &index)
{
    this->dbName=index.parent().data().toString();
    QString table=dbName.append(".");
    table.append(index.data().toString());
    QSqlQuery getData(this->conn);
    getData.exec("SELECT * FROM "+table);
    QList <QString>fields;
    for(int i=0;i<getData.record().count();i++){
        fields.append( getData.record().fieldName(i));
    }
    QStandardItemModel* model = new QStandardItemModel(ui->table);
    model->setHorizontalHeaderLabels(fields);
    while (getData.next()) {
        QList<QStandardItem*> temp;
       for(int i=0;i<getData.record().count();i++){
            temp.append(new QStandardItem(getData.value(i).toString()));
        }
         model->appendRow(temp);
    }
    ui->table->setModel(model);
    ui->table->horizontalHeader()->setSectionResizeMode(QHeaderView::ResizeToContents);
    ui->select_2->setEnabled(true);
    ui->insert_2->setEnabled(true);
    ui->update_2->setEnabled(true);
    ui->delete_3->setEnabled(true);
}


void Index::on_select_2_clicked()
{
    ui->control->setHidden(false);
    ui->table->setHidden(true);
    ui->TableListComboBox_5->addItems(this->tableList);
    ui->TableListComboBox_5->setCurrentIndex(-1);
}

void Index::on_TableListComboBox_5_activated(const QString &arg1)
{
    QString str=ui->selectedTable->toPlainText();
    QSqlQuery exe(this->conn);
    this->selectTable=str.split(" , ");
    for(int i=0;i<selectTable.length();i++){
        if(selectTable[i].contains(arg1)){
         return ;
        }
    }
    if(str==""){
         ui->selectedTable->setText(str+arg1);
    }else {
        ui->selectedTable->setText(str+" , "+arg1);
    }
    exe.exec("show columns from "+arg1);
    ui->fieldListComboBox_4->addItem("-------------"+arg1+"-------------");
    this->fieldList.append("-------------"+arg1+"-------------");
    while (exe.next()) {
        ui->fieldListComboBox_4->addItem(arg1+"."+exe.value(0).toString()+":"+exe.value(1).toString());
        this->fieldList.append(arg1+"."+exe.value(0).toString()+":"+exe.value(1).toString());
    }
    ui->fieldListComboBox->clear();
    ui->fieldListComboBox->addItems(this->fieldList);
    ui->fieldListComboBox->setCurrentIndex(-1);
    ui->orderComboBox->clear();
    ui->orderComboBox->addItems(this->fieldList);
    ui->orderComboBox->setCurrentIndex(-1);
}

void Index::on_fieldListComboBox_4_activated(const QString &arg1)
{
    QString str=ui->selectedField->toPlainText();
    this->selectedField=str.split(" , ");
    QStringList a=arg1.split(":");
    for(int i=0;i<this->selectedField.length();i++){
        if(this->selectedField[i].contains(a[0])){
             return ;
        }
    }
    if(str==""){
         ui->selectedField->setText(str+a[0]);
    }else {
        ui->selectedField->setText(str+" , "+a[0]);
    }
}

void Index::on_add_clicked()
{
    QString str=ui->ConditionText->toPlainText();
    QString field=ui->fieldListComboBox->currentText();
    QString domain=ui->domainComboBox->currentText();
    QString value=ui->valueLineEdit->text();
    QString andor=ui->comboBox_2->currentText();
   QStringList a=ui->fieldListComboBox->currentText().split(":");
    if(field==""||domain==""||value==""){
        QMessageBox::warning(this,"ERROR","Value Can't Is NULL!",QMessageBox::Yes);
    }else {
        if(str==""){
             ui->ConditionText->setText(str+a[0]+domain+value+" "+andor);
        }else {
            ui->ConditionText->setText(str+" "+a[0]+domain+value+" "+andor);
        }
    }
    ui->fieldListComboBox->setCurrentIndex(-1);
    ui->domainComboBox->setCurrentIndex(-1);
    ui->valueLineEdit->setText("");
    ui->comboBox_2->setCurrentIndex(-1);
}

void Index::on_submit_clicked()
{
    QStringList a=ui->orderComboBox->currentText().split(":");
    QString sql="SELECT "+ui->selectedField->toPlainText()+" FROM "+ui->selectedTable->toPlainText();
    if(ui->ConditionText->toPlainText()!=""){
        sql+=" WHERE "+ui->ConditionText->toPlainText();
    }
    if(ui->orderComboBox->currentText()!="" && ui->sortcomboBox->currentText()!="" ){
        sql+=" ORDER BY "+a[0]+" "+ui->sortcomboBox->currentText();
    }
    ui->control->setHidden(true);
    ui->table->setHidden(false);
    QSqlQuery exe(this->conn);
    QList <QString> fields;
    exe.exec(sql);
    for(int i=0;i<exe.record().count();i++){
        fields.append(exe.record().fieldName(i));
    }
    QStandardItemModel* model = new QStandardItemModel(ui->table);
    model->setHorizontalHeaderLabels(fields);
    QList <QStandardItem*>items;
    while (exe.next()) {
        QList<QStandardItem*> temp;
       for(int i=0;i<exe.record().count();i++){
            temp.append(new QStandardItem(exe.value(i).toString()));
        }
         model->appendRow(temp);
    }
    ui->table->setModel(model);
}

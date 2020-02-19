package models

import (
	"db"
)

type Model struct {
	table   string
	dbnanme string
	dbObj   *db.DBPool
}

func InitModel(dbname string, table string) *Model {
	model := new(Model)
	model.table = table
	model.dbnanme = dbname
	model.dbObj = db.DB(dbname).Table(table)
	return model
}

/**
获取单条数据通过where,like
*/
func (m *Model) GetOneByWhereLike(selectField string, whereMap map[string]interface{}, likeMap map[string]interface{}) map[string]interface{} {
	row := m.dbObj.MyFilterKeyEqValue(whereMap).MyFilterKeyLikeValue(likeMap).SelectField(selectField).MyGetOne()
	return row
}

/**
获取所有数据通过where,like
*/
func (m *Model) GetAllByWhereLike(selectField string, whereMap map[string]interface{}, likeMap map[string]interface{}) []map[string]interface{} {
	row := m.dbObj.MyFilterKeyEqValue(whereMap).MyFilterKeyLikeValue(likeMap).SelectField(selectField).MyGetAll()
	return row
}

/**
获取一条数据通过sql
*/
func (m *Model) GetOneBySql(sql string) map[string]interface{} {
	row := m.dbObj.FetchOne(sql)
	return row
}

/**
获取所有数据通过sql语句
*/
func (m *Model) GetAllBySql(sql string) []map[string]interface{} {
	row := m.dbObj.FetchAll(sql)
	return row
}

func (m *Model) ExecuteSql(sql string) {
	m.dbObj.Execute(sql)
}

/**
更新数据
*/
func (m *Model) UpdateByWhereLike(dataMap map[string]interface{}, whereMap map[string]interface{}, likeMap map[string]interface{}) int {
	id, err := m.dbObj.MyFilterKeyEqValue(whereMap).MyFilterKeyLikeValue(likeMap).Update(dataMap)
	if err == nil {
		return id
	}
	return 0
}

/**
插入数据
*/
func (m *Model) Insert(createMap map[string]interface{}) int {
	id, err := m.dbObj.Create(createMap)
	if err == nil {
		return id
	}
	return 0
}

func (m *Model) CheckTableIsExists() bool {
	return m.dbObj.CheckTableIsExists(m.dbnanme)
}

func (m *Model) Delete(whereMap map[string]interface{}, likeMap map[string]interface{}) int {

	id, err := m.dbObj.MyFilterKeyEqValue(whereMap).MyFilterKeyLikeValue(likeMap).Delete()
	if err == nil {
		return id
	}
	return 0
}

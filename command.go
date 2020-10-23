package torm

import "database/sql"

// Insert one record.
func (this *DbBaseDao) Create(bean interface{}) (int64, error) {
	if this.Session == nil {
		return this.Engine.Insert(bean)
	} else {
		return this.Session.Insert(bean)
	}
}

// Update records, bean's non-empty fields are updated contents, Update with ID primary key.
func (this *DbBaseDao) Update(bean interface{}) (int64, error) {
	if this.Session == nil {
		return this.Engine.ID(bean).AllCols().Update(bean)
	} else {
		return this.Engine.ID(bean).AllCols().Update(bean)
	}
}

//Update special specified fields, Update with ID primary key.
func (this *DbBaseDao) UpdateCols(bean interface{}, cols ...string) (int64, error) {
	if this.Session == nil {
		return this.Engine.ID(bean).Cols(cols...).Update(bean)
	} else {
		return this.Engine.ID(bean).Cols(cols...).Update(bean)
	}
}

// Delete records, bean's non-empty fields are conditions, Delete with ID primary key.
func (this *DbBaseDao) Delete(bean interface{}) (int64, error) {
	if this.Session == nil {
		return this.Engine.ID(bean).Delete(bean)
	} else {
		return this.Engine.ID(bean).Delete(bean)
	}
}

// Exec raw sql
func (this *DbBaseDao) Exec(sqlOrArgs ...interface{}) (sql.Result, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.Exec(sqlOrArgs...)
}

// Query a raw sql and return records as []map[string][]byte
func (this *DbBaseDao) Query(sqlorArgs ...interface{}) (resultsSlice []map[string][]byte, err error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.Query(sqlorArgs...)
}

// QueryString runs a raw sql and return records as []map[string]string
func (this *DbBaseDao) QueryString(sqlorArgs ...interface{}) ([]map[string]string, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.QueryString(sqlorArgs...)
}

// QueryInterface runs a raw sql and return records as []map[string]interface{}
func (this *DbBaseDao) QueryInterface(sqlorArgs ...interface{}) ([]map[string]interface{}, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.QueryInterface(sqlorArgs...)
}

// Get retrieve one record from table, bean's non-empty fields
// are conditions
func (this *DbBaseDao) Get(bean interface{}) (bool, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.Get(bean)
}

// Exist returns true if the record exist otherwise return false
func (this *DbBaseDao) Exist(bean ...interface{}) (bool, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.Exist(bean...)
}

// Find retrieve records from table, condiBeans's non-empty fields
// are conditions. beans could be []Struct, []*Struct, map[int64]Struct
// map[int64]*Struct
func (this *DbBaseDao) Find(beans interface{}, condiBeans ...interface{}) error {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.Find(beans, condiBeans...)
}

// FindAndCount find the results and also return the counts
func (this *DbBaseDao) FindAndCount(rowsSlicePtr interface{}, condiBean ...interface{}) (int64, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.FindAndCount(rowsSlicePtr, condiBean...)
}

// Count counts the records. bean's non-empty fields are conditions.
func (this *DbBaseDao) Count(bean ...interface{}) (int64, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.Count(bean...)
}

// Sum sum the records by some column. bean's non-empty fields are conditions.
func (this *DbBaseDao) Sum(bean interface{}, colName string) (float64, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.Sum(bean, colName)
}

// SumInt sum the records by some column. bean's non-empty fields are conditions.
func (this *DbBaseDao) SumInt(bean interface{}, colName string) (int64, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.SumInt(bean, colName)
}

// Sums sum the records by some columns. bean's non-empty fields are conditions.
func (this *DbBaseDao) Sums(bean interface{}, colNames ...string) ([]float64, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.Sums(bean, colNames...)
}

// SumsInt like Sums but return slice of int64 instead of float64.
func (this *DbBaseDao) SumsInt(bean interface{}, colNames ...string) ([]int64, error) {
	session := this.Engine.NewSession()
	defer session.Close()
	return session.SumsInt(bean, colNames...)
}

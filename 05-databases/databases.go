package _0205_databases_

import (
	accessingDatabases "go-advanced/05-databases/01-accessingDatabases"
	dataRetrieval "go-advanced/05-databases/02-dataRetrieval"
	preparedStatements "go-advanced/05-databases/03-preparedStatements"
	dataManipulation "go-advanced/05-databases/04-dataManipulation"
	transactions "go-advanced/05-databases/05-transactions"
)

var Demos = struct {
	AccessingDatabases        func()
	SingleRowDataRetrieval    func()
	MultipleRowsDataRetrieval func()
	PreparedStatements        func()
	DataManipulation          func()
	Transactions              func()
}{
	AccessingDatabases:        accessingDatabases.Demo,
	SingleRowDataRetrieval:    dataRetrieval.SingleRowDemo,
	MultipleRowsDataRetrieval: dataRetrieval.MultipleRowsDemo,
	PreparedStatements:        preparedStatements.Demo,
	DataManipulation:          dataManipulation.Demo,
	Transactions:              transactions.Demo,
}

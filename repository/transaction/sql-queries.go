package transaction

const (
	createTransaction = "INSERT INTO transaction (valued, status, info, `from`, `to`, account_id, amount_id) VALUES (now(), ?, ?, ?, ?, ?, ?)"

	getTransactions = `Select * from transaction inner join amount on amount_id=amount.id where account_id =?`
	getTransaction  = `Select * from transaction where id =?`

	updateAmount = `Update amount set value=? where id=?`

	createAmount = `INSERT INTO amount (currency, value) 
	VALUES (?, ?)`
)

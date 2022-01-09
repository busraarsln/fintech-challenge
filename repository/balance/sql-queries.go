package balance

const (
	createBalance = `INSERT INTO balance (credit_limit, amount_id, created_at) 
	VALUES (?, ?, now())`
	getBalance = `Select balance.id,balance.credit_limit,balance.created_at, amount.* from account inner join balance on balance.id = account.balance_id inner join amount on amount_id=balance.amount_id where account.id=?`

	updateBalance = `Update amount set value = ?, currency=?  where id = ?`
)

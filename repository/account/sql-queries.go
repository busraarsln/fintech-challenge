package account

const (
	createAccount = `INSERT INTO account (type, account_no, iban, currency, description, status, nickname, created_at, updated_at, is_active, customer_id, balance_id) 
	VALUES (?, ?, ?, ?, ?, ?, ?, now(), now(), 1, ?,? )`

	getAccount = `Select account.*, balance.id, balance.credit_limit, balance.amount_id, amount.id, amount.currency, amount.value from account inner join balance on balance.id = account.balance_id
	inner join amount on amount.id = balance.amount_id where account.id = ?`

	getAccountByIban = `Select account.*, balance.id, balance.credit_limit, balance.amount_id, amount.id, amount.currency, amount.value from account inner join balance on balance.id = account.balance_id
	inner join amount on amount.id = balance.amount_id where account.iban = ? and is_active = 1`

	getAccounts = `Select account.*, balance.id, balance.credit_limit, balance.amount_id, amount.id, amount.currency, amount.value from account inner join balance on balance.id = account.balance_id
	inner join amount on amount.id = balance.amount_id where customer_id=? and is_active = 1`

	deleteAccount = `Update account set is_active = 0 where id = ?`
)

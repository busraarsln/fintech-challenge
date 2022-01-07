package payment

const (
	createPayment = `INSERT INTO payment (from, to, booked, valued, account_id, amount_id) 
	VALUES (?, ?, ?, ?, ?, ? )`

	getPayments = `Select * from payment where account_id=?`
	getPayment = `Select * from payment where id=?`
)

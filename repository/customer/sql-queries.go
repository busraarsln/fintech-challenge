package customer

const (
	createCustomer = `INSERT INTO customer (name, surname, phone_number, email, role, created_at, updated_at, is_active, password) 
	VALUES (?, ?, ?, ?, ?, now(), now(), 1, ?)`

	getCustomer = `Select * from customer where id=?`

	getCustomers = `Select * from customer where is_active = 0`

	deleteCustomer = `Update customer set is_active = 0 where id = ?`
)

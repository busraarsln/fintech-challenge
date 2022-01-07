create table amount
(
    id       int auto_increment
        primary key,
    currency varchar(10)    not null,
    value    decimal(18, 9) not null
);

create table balance
(
    id           int auto_increment
        primary key,
    credit_limit decimal(18, 9) not null,
    amount_id    int            null,
    updated_at   datetime       null,
    created_at   datetime       null,
    constraint balance_amount_id_fk
        foreign key (amount_id) references amount (id)
            on delete cascade
);

create index balance_amount_id_index
    on balance (amount_id);

create table customer
(
    id           int auto_increment
        primary key,
    name         varchar(30) not null,
    surname      varchar(30) not null,
    phone_number varchar(30) null,
    email        varchar(30) null,
    role         varchar(30) not null,
    created_at   datetime    not null,
    updated_at   datetime    null,
    is_active    tinyint(1)  not null,
    password     varchar(60) not null
);

create table account
(
    id          int auto_increment
        primary key,
    type        varchar(10)          null,
    account_no  varchar(16)          not null,
    iban        varchar(36)          not null,
    currency    varchar(10)          null,
    description text                 null,
    nickname    varchar(30)          null,
    created_at  datetime             not null,
    updated_at  datetime             null,
    is_active   tinyint(1) default 1 not null,
    customer_id int                  null,
    balance_id  int                  null,
    status      varchar(10)          null,
    constraint account_balance_id_fk
        foreign key (balance_id) references balance (id)
            on delete cascade,
    constraint account_customer_id_fk
        foreign key (customer_id) references customer (id)
);

create index account_balance_id_customer_id_index
    on account (balance_id, customer_id);

create table address
(
    id               int auto_increment
        primary key,
    type             varchar(15)  null,
    city             varchar(30)  null,
    district         varchar(30)  null,
    neighborhood     varchar(30)  null,
    description      varchar(150) null,
    phone_number     varchar(30)  null,
    floor            varchar(10)  null,
    door_number      varchar(10)  null,
    apartment_number varchar(10)  null,
    postal_code      varchar(10)  null,
    customer_id      int          null,
    constraint address_customer_id_fk
        foreign key (customer_id) references customer (id)
);

create index address_customer_id_index
    on address (customer_id);

create table payment
(
    id         int auto_increment
        primary key,
    `from`     varchar(26) not null,
    `to`       varchar(26) not null,
    booked     datetime    not null,
    valued     datetime    null,
    account_id int         null,
    amount_id  int         null,
    constraint payment_account_id_fk
        foreign key (account_id) references account (id),
    constraint payment_amount_id_fk
        foreign key (amount_id) references amount (id)
);

create index payment_account_id_index
    on payment (account_id);

create index payment_amount_id_index
    on payment (amount_id);

create table transaction
(
    id         int auto_increment
        primary key,
    valued     datetime     not null,
    status     varchar(30)  null,
    info       varchar(150) null,
    `from`     varchar(36)  null,
    `to`       varchar(36)  null,
    account_id int          null,
    amount_id  int          null,
    constraint transaction_account_id_fk
        foreign key (account_id) references account (id),
    constraint transaction_amount_id_fk
        foreign key (amount_id) references amount (id)
);

create index transaction_account_id_index
    on transaction (account_id);

create index transaction_amount_id_index
    on transaction (amount_id);


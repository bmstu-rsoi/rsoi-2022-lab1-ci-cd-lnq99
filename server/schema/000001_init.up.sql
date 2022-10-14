create table Persons (
    id      serial      primary key,
    name    varchar(64) not null,
    age     int         default 0,
    address text        default '',
    work    text        default ''
);
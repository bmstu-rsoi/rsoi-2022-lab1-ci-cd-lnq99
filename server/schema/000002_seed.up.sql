insert into Persons(name, age, work, address)
values
    ('User 1', 21, 'Work 1', 'Address 1'),
    ('User 2', 22, 'Work 2', 'Address 2'),
    ('User 3', 23, 'Work 3', 'Address 3'),
    ('User 4', 24, 'Work 4', 'Address 4')
on conflict
do nothing;
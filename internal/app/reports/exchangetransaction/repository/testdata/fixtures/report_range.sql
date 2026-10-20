INSERT INTO exchanges (id, name, minimum_age, maximum_transfer_amount)
VALUES ('aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa', 'report fixture', 0, 1000000.00);

INSERT INTO exchange_transactions (id, exchange_id, type, amount, executed_at)
VALUES
    ('00000000-0000-4000-8000-000000000001', 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa', 'DEPOSIT', 999.00, '2024-01-01 12:00:00'),
    ('00000000-0000-4000-8000-000000000002', 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa', 'TRANSFER', 11.00, '2024-01-02 09:00:00'),
    ('00000000-0000-4000-8000-000000000003', 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa', 'DEPOSIT', 10.00, '2024-01-03 10:00:00'),
    ('00000000-0000-4000-8000-000000000004', 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa', 'DEPOSIT', 15.00, '2024-01-03 15:00:00'),
    ('00000000-0000-4000-8000-000000000005', 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa', 'WITHDRAWAL', 7.00, '2024-01-03 18:00:00'),
    ('00000000-0000-4000-8000-000000000006', 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa', 'DEPOSIT', 20.00, '2024-01-04 08:00:00'),
    ('00000000-0000-4000-8000-000000000007', 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa', 'DEPOSIT', 888.00, '2024-01-05 12:00:00');

INSERT INTO users (id, name, birthday, document_number)
VALUES
    ('10000000-0000-4000-8000-000000000001', 'integration account user', '1990-01-01', 'integration-account-doc-001'),
    ('10000000-0000-4000-8000-000000000002', 'integration second account user', '1991-02-02', 'integration-account-doc-002');

INSERT INTO exchanges (id, name, minimum_age, maximum_transfer_amount)
VALUES
    ('20000000-0000-4000-8000-000000000001', 'integration account exchange one', 0, 100000.00),
    ('20000000-0000-4000-8000-000000000002', 'integration account exchange two', 0, 100000.00);

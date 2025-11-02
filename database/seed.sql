INSERT INTO payment_providers (id, name) VALUES
(1, 'VNPAY'),
(2, 'MOMO'),
(3, 'ZALOPAY')
ON CONFLICT (id) DO NOTHING;

ALTER SEQUENCE payment_providers_id_seq RESTART WITH 4;

INSERT INTO payment_statuses (id, name) VALUES
(1, 'Pending'),
(2, 'Completed'),
(3, 'Failed')
ON CONFLICT (id) DO NOTHING;

ALTER SEQUENCE payment_statuses_id_seq RESTART WITH 4;

INSERT INTO order_statuses (id, name) VALUES
(1, 'Pending'),
(2, 'Processing'),
(3, 'Shipped'),
(4, 'Delivered'),
(5, 'Cancelled')
ON CONFLICT (id) DO NOTHING;

ALTER SEQUENCE order_statuses_id_seq RESTART WITH 6;

INSERT INTO return_request_statuses (id, name) VALUES
(1, 'Pending'),
(2, 'Approved'),
(3, 'Rejected'),
(4, 'Completed')
ON CONFLICT (id) DO NOTHING;

ALTER SEQUENCE return_request_statuses_id_seq RESTART WITH 5;

INSERT INTO refund_statuses (id, name) VALUES
(1, 'Pending'),
(2, 'Processed'),
(3, 'Failed')
ON CONFLICT (id) DO NOTHING;

ALTER SEQUENCE refund_statuses_id_seq RESTART WITH 4;

CREATE UNIQUE INDEX idx_orders_unique_pending
ON orders (user_id, course_id)
WHERE status = 'pending';
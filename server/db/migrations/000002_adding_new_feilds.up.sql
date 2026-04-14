CREATE TABLE reviews_products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL REFERENCES products(id),
    user_id    UUID NOT NULL REFERENCES users(id),
    order_id   UUID NOT NULL REFERENCES orders(id),
    rating     SMALLINT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment    TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, order_id)
);


-- ALTER TABLE reviews_products
--     DROP CONSTRAINT reviews_user_order_unique,
--     ADD CONSTRAINT reviews_user_order_product_unique
--         UNIQUE (user_id, order_id, product_id);

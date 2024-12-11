INSERT INTO categories (name, created_at, updated_at)
VALUES 
    ('Electronics', NOW(), NOW()),
    ('Books', NOW(), NOW()),
    ('Clothing', NOW(), NOW()),
    ('Home & Kitchen', NOW(), NOW()),
    ('Sports', NOW(), NOW());

INSERT INTO products (name, description, price, category_id, created_at, updated_at)
VALUES
    ('Smartphone X', 'Latest smartphone with advanced features.', 799.99, 1, NOW(), NOW()),
    ('Laptop Pro', 'High-performance laptop for professionals.', 1299.50, 1, NOW(), NOW()),
    ('Wireless Headphones', 'Noise-cancelling over-ear headphones.', 199.99, 1, NOW(), NOW()),
    ('Fiction Novel', 'Bestselling fiction novel by a renowned author.', 14.99, 2, NOW(), NOW()),
    ('Cookbook', 'Delicious recipes for home cooking.', 29.99, 2, NOW(), NOW()),
    ('Men''s T-Shirt', '100% cotton, various sizes available.', 19.99, 3, NOW(), NOW()),
    ('Women''s Jeans', 'Comfortable and stylish denim jeans.', 49.99, 3, NOW(), NOW()),
    ('Blender', 'High-speed blender for smoothies and more.', 89.99, 4, NOW(), NOW()),
    ('Coffee Maker', 'Automatic drip coffee maker with timer.', 59.99, 4, NOW(), NOW()),
    ('Yoga Mat', 'Non-slip yoga mat for all types of exercises.', 25.99, 5, NOW(), NOW()),
    ('Running Shoes', 'Lightweight running shoes for all terrains.', 89.99, 5, NOW(), NOW()),
    ('Smartwatch', 'Track your fitness and notifications.', 149.99, 1, NOW(), NOW()),
    ('E-book Reader', 'Portable e-book reader with high-resolution display.', 129.99, 1, NOW(), NOW()),
    ('Mystery Thriller', 'Engaging mystery thriller novel.', 12.99, 2, NOW(), NOW()),
    ('Dress', 'Elegant evening dress in various colors.', 79.99, 3, NOW(), NOW()),
    ('Air Fryer', 'Healthier way to fry with little oil.', 99.99, 4, NOW(), NOW()),
    ('Dumbbell Set', 'Adjustable dumbbell set for home workouts.', 59.99, 5, NOW(), NOW()),
    ('Tablet S', 'Lightweight tablet with high-resolution screen.', 399.99, 1, NOW(), NOW()),
    ('Graphic T-Shirt', 'T-shirts with unique graphic designs.', 24.99, 3, NOW(), NOW()),
    ('Gardening Kit', 'Complete kit for home gardening.', 34.99, 4, NOW(), NOW()),
    ('Basketball', 'Official size and weight basketball.', 29.99, 5, NOW(), NOW());

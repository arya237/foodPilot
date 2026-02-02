-------------- Tables -------------------
CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    auto_save BOOLEAN DEFAULT FALSE,
    role user_role DEFAULT 'user',
    token VARCHAR(10000)
);

CREATE TABLE restaurant_credentials (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    username VARCHAR(100),
    password VARCHAR(100),
    token text
);

CREATE TABLE foods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(500) NOT NULL
);

CREATE TABLE rates (
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    food_id INTEGER NOT NULL REFERENCES foods(id) ON DELETE CASCADE,
    score INTEGER NOT NULL,
    PRIMARY KEY (user_id, food_id)
);

------------ Create indexes for ----------
CREATE INDEX idx_users_username ON users(username);

------------ Test Data ----------------------
INSERT INTO foods (name) VALUES
('چلو کباب کوبیده زعفرانی'),
('خوراک گوشت چرخ‌کرده با سیب زمینی'),
('چلو خورشت آلو با اسفناج'),
('چلو جوجه کباب'),
('فرنی'),
('لوبیا پلو'),
('خوراک فلافل'),
('خوراک فیله سوخاری'),
('چلو خورشت قیمه'),
('کلم پلو'),
('خوراک عدسی'),
('زرشک پلو با مرغ'),
('شکلات صبحانه ،شیر،پنیر ،چای'),
('تخم مرغ(24عدد)'),
('پنیر ،مربا،خامه،چای'),
('تخم مرغ(14 عدد)'),
('شیرموز(2عدد)، شیر کاکائو(2عدد)،کیک(4عدد)،چایی(5عدد)'),
('غلات صبحانه، شیر'),
('خوراک پاستا'),
('چلو ماهی قزل آلا'),
('خوراک ناگت مرغ'),
('استانبولی پلو'),
('چلو تن ماهی'),
('عدس پلو'),
('چلو خورشت قورمه سبزی'),
('خوراک دلمه'),
('چلو کباب کوبیده مرغ'),
('خوراک شنیسل مرغ'),
('خوراک لوبیا + پوره'),
('سوپ جو'),
('چلو خورشت بادمجان');

INSERT INTO users (username, password, auto_save, role, token) values ('40112358043', 'arya1383', TRUE, 'admin', '')

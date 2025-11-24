-------------- Tables -------------------
CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    auto_save BOOLEAN DEFAULT FALSE,
    role user_role DEFAULT 'user',
    token VARCHAR(500)
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
INSERT INTO foods (id, name) VALUES
(1, 'چلو کباب کوبیده زعفرانی'),
(3, 'خوراک گوشت چرخ‌کرده با سیب زمینی'),
(4, 'چلو خورشت آلو با اسفناج'),
(5, 'چلو جوجه کباب'),
(6, 'فرنی'),
(7, 'لوبیا پلو'),
(8, 'خوراک فلافل'),
(9, 'خوراک فیله سوخاری'),
(10, 'چلو خورشت قیمه'),
(11, 'کلم پلو'),
(12, 'خوراک عدسی'),
(13, 'زرشک پلو با مرغ'),
(14, 'شکلات صبحانه ،شیر،پنیر ،چای'),
(15, 'تخم مرغ(24عدد)'),
(16, 'پنیر ،مربا،خامه،چای'),
(17, 'تخم مرغ(14 عدد)'),
(18, 'شیرموز(2عدد)، شیر کاکائو(2عدد)،کیک(4عدد)،چایی(5عدد)'),
(19, 'غلات صبحانه، شیر'),
(20, 'خوراک پاستا'),
(21, 'چلو ماهی قزل آلا'),
(22, 'خوراک ناگت مرغ'),
(23, 'استانبولی پلو'),
(24, 'چلو تن ماهی'),
(25, 'عدس پلو'),
(26, 'چلو خورشت قورمه سبزی'),
(27, 'خوراک دلمه'),
(28, 'چلو کباب کوبیده مرغ'),
(29, 'خوراک شنیسل مرغ'),
(30, 'خوراک لوبیا + پوره'),
(31, 'سوپ جو'),
(32, 'چلو خورشت بادمجان');

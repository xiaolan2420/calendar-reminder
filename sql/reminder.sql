CREATE TABLE reminders
(
    id              INT AUTO_INCREMENT PRIMARY KEY,
    creator_id      INT,
    content         TEXT        NOT NULL,
    reminder_at     DATETIME    NOT NULL,
    reminder_method VARCHAR(10) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                             -- 创建时间
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间
    FOREIGN KEY (creator_id) REFERENCES users (id) ON DELETE CASCADE
);

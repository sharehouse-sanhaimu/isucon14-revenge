-- chairsテーブルからデータを挿入
INSERT INTO chair_available (chair_id)
SELECT id
FROM chairs;

-- 外部キー制約を追加（オプション）
-- ALTER TABLE chair_available
-- ADD CONSTRAINT fk_chair_id
-- FOREIGN KEY (chair_id)
-- REFERENCES chairs (id)
-- ON DELETE CASCADE;

-- インデックスの追加（オプション）
-- CREATE INDEX idx_chair_id ON chair_available (chair_id);

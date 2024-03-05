ALTER TABLE tasks DROP CONSTRAINT tasks_group_id_foreign;
ALTER TABLE tasks DROP COLUMN group_id;
DROP TABLE task_groups;
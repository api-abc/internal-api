CREATE table data (
    name varchar(100) NOT NULL,
    age int NOT NULL,
    status boolean DEFAULT false,
    job_details jsonb,
    worker_update timestamp
)

INSERT INTO data(name, age, status, job_details, worker_update)
VALUES (
    'Lala',
    22,
    true,
    '{"position":"Employee", "years_work_experience":2, "work_status":"retired"}',
    current_timestamp
)

SELECT name, age, status, job_details, worker_update FROM data WHERE name = 'Lala' AND status = true
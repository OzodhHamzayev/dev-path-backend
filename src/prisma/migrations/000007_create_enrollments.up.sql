CREATE TABLE enrollments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    course_id INTEGER NOT NULL
);


ALTER TABLE enrollments ADD CONSTRAINT fk_enrollments_user FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE enrollments ADD CONSTRAINT fk_enrollments_course FOREIGN KEY (course_id) REFERENCES courses(id);